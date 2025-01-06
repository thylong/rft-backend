package server

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thylong/rft-backend/pkg/db" // Import sqlc-generated code
	pb "github.com/thylong/rft-backend/pkg/proto"
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
		return nil, fmt.Errorf("failed to fetch okrs: %w", err)
	}

	totalCount, err := s.queries.GetOkrsCount(ctx, pgtype.Text{String: req.Search, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("failed to fetch okr count: %w", err)
	}

	var pbOkrs []*pb.Okr
	for _, e := range okrs {
		pbOkrs = append(pbOkrs, &pb.Okr{
			Id:            e.OkrID.String(),
			Name:          e.OkrName,
			Number:        e.OkrNumber,
			Description:   e.OkrDescription,
			EmbeddedChild: e.KeyResults.([]*pb.Kr),
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
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	// Create a pgtype.UUID instance
	okrID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	okr, err := s.queries.GetOkrByID(ctx, okrID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("okr not found")
		}
		return nil, fmt.Errorf("failed to fetch okr: %w", err)
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
		Id:            req.Id.String(),
		Name:          req.OkrName,
		Number:        req.OkrNumber,
		Description:   req.OkrDescription,
		EmbeddedChild: req.KeyResults.([]*pb.Kr),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to insert okr: %w", err)
	}

	startAtProto, _ := convertToProtoDateTime(okr.StartAt)

	return &pb.PutOkrResponse{
		Okr: &pb.Okr{
			OkrId:       okr.OkrID.String(),
			OkrPrivacy:  pb.OkrPrivacy(okr.OkrPrivacy),
			Name:        okr.Name,
			Description: okr.Description,
			Type:        okr.Type,
			Department:  okr.Department,
			Regions:     okr.Regions,
			Tags:        okr.Tags,
			StartAt:     startAtProto,
		},
	}, nil
}

// DeleteOkr handles deleting an okr by ID
func (s *OkrServiceServer) DeleteOkr(ctx context.Context, req *pb.DeleteOkrRequest) (*pb.DeleteOkrResponse, error) {
	// Parse and validate the UUID
	parsedUUID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, fmt.Errorf("invalid UUID: %w", err)
	}

	// Create a pgtype.UUID instance
	okrID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	err = s.queries.DeleteOkr(ctx, okrID)
	if err != nil {
		return nil, fmt.Errorf("failed to delete okr: %w", err)
	}
	return &pb.DeleteOkrResponse{}, nil
}
