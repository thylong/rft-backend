package handler

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thylong/rft-backend/pkg/db" // Import sqlc-generated code
	pb "github.com/thylong/rft-backend/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type OkrServiceServer struct {
	pb.UnimplementedOkrServiceServer
	queries *db.Queries
}

// NewOkrServiceServer initializes the server with sqlc queries
func NewOkrServiceServer(queries *db.Queries) *OkrServiceServer {
	return &OkrServiceServer{queries: queries}
}

// GetOkrs handles fetching paginated and filtered okrs
func (s *OkrServiceServer) GetOkrs(ctx context.Context, req *pb.GetOkrsRequest) (*pb.GetOkrsResponse, error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	okrs, err := s.queries.GetOkrs(ctx, db.GetOkrsParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch okrs: %s", err)
	}

	totalCount, err := s.queries.GetOkrsCount(ctx, pgtype.Text{String: req.Search, Valid: true})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch okr count: %s", err)
	}

	var pbOkrs []*pb.Okr
	for _, e := range okrs {
		// Convert KeyResults to []*pb.Kr
		var pbKeyResults []*pb.Kr
		if keyResults, ok := e.KeyResults.([]interface{}); ok {
			for _, kr := range keyResults {
				if krMap, ok := kr.(map[string]interface{}); ok {
					pbKeyResults = append(pbKeyResults, &pb.Kr{
						Id:          krMap["id"].(string),
						Name:        krMap["name"].(string),
						Number:      int32(krMap["number"].(float64)),
						Description: krMap["description"].(string),
						Sponsor:     krMap["sponsor"].(string),
						Kpis:        krMap["kpis"].(string),
						Scope:       krMap["scope"].(string),
						Initiatives: krMap["initiatives"].(string),
					})
				}
			}
		}

		// Append OKR to the response list
		pbOkrs = append(pbOkrs, &pb.Okr{
			Id:            e.OkrID.String(),
			Name:          e.OkrName,
			Number:        e.OkrNumber,
			Description:   e.OkrDescription,
			EmbeddedChild: pbKeyResults,
		})
	}

	return &pb.GetOkrsResponse{
		Okrs:       pbOkrs,
		TotalCount: int32(totalCount),
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// GetOkr handles fetching a single okr by ID
func (s *OkrServiceServer) GetOkr(ctx context.Context, req *pb.GetOkrRequest) (*pb.GetOkrResponse, error) {
	// Parse and validate the UUID
	parsedUUID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UUID: %s", err)
	}

	// Create a pgtype.UUID instance
	okrID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	okr, err := s.queries.GetOkrByID(ctx, okrID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "okr not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to fetch okr: %s", err)
	}

	return &pb.GetOkrResponse{
		Okr: &pb.Okr{
			Id:            okr.OkrID.String(),
			Name:          okr.OkrName,
			Number:        okr.OkrNumber,
			Description:   okr.OkrDescription,
			EmbeddedChild: okr.KeyResults.([]*pb.Kr),
		},
	}, nil
}

// PutOkr handles inserting a new okr
func (s *OkrServiceServer) PutOkr(ctx context.Context, req *pb.PutOkrRequest) (*pb.PutOkrResponse, error) {
	okr, err := s.queries.InsertOkr(ctx, db.InsertOkrParams{
		Name:        req.Name,
		Number:      req.Number,
		Year:        req.Year,
		Description: req.Description,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert okr: %s", err)
	}

	return &pb.PutOkrResponse{
		Okr: &pb.Okr{
			Id:          okr.OkrID.String(),
			Name:        okr.OkrName,
			Number:      okr.OkrNumber,
			Year:        okr.OkrYear,
			Description: okr.OkrDescription,
		},
	}, nil
}

// DeleteOkr handles deleting an okr by ID
func (s *OkrServiceServer) DeleteOkr(ctx context.Context, req *pb.DeleteOkrRequest) (*pb.DeleteOkrResponse, error) {
	// Parse and validate the UUID
	parsedUUID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UUID: %s", err)
	}

	// Create a pgtype.UUID instance
	okrID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	err = s.queries.DeleteOkr(ctx, okrID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete okr: %s", err)
	}
	return &pb.DeleteOkrResponse{}, nil
}

// PutKr handles inserting a new kr
func (s *OkrServiceServer) PutKr(ctx context.Context, req *pb.PutKrRequest) (*pb.PutKrResponse, error) {
	// TODO: Make sure the okr exists

	var pgUUID pgtype.UUID
	err := pgUUID.Scan(req.OkrId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse okr_id UUID: %s", err)
	}
	kr, err := s.queries.InsertKeyResult(ctx, db.InsertKeyResultParams{
		OkrID:       pgUUID,
		Name:        req.Name,
		Number:      req.Number,
		Description: req.Description,
		Sponsor:     req.Sponsor,
		Kpis:        req.Kpis,
		Scope:       req.Scope,
		Initiatives: req.Initiatives,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert kr: %s", err)
	}

	return &pb.PutKrResponse{
		Kr: &pb.Kr{
			Id:          kr.ID.String(),
			OkrId:       kr.OkrID.String(),
			Name:        kr.Name,
			Number:      kr.Number,
			Description: kr.Description,
			Sponsor:     kr.Sponsor,
			Kpis:        kr.Kpis,
			Scope:       kr.Scope,
			Initiatives: kr.Initiatives,
		},
	}, nil
}
