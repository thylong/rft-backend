package handler

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/thylong/rft-backend/pkg/db" // Import sqlc-generated code
	kpipb "github.com/thylong/rft-backend/pkg/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KpiServiceServer struct {
	kpipb.UnimplementedKpiServiceServer
	queries *db.Queries
}

// NewKpiServiceServer initializes the server with sqlc queries
func NewKpiServiceServer(queries *db.Queries) *KpiServiceServer {
	return &KpiServiceServer{queries: queries}
}

// GetKpis handles fetching paginated and filtered kpis
func (s *KpiServiceServer) GetKpis(ctx context.Context, req *kpipb.GetKpisRequest) (*kpipb.GetKpisResponse, error) {
	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 10
	}
	offset := (page - 1) * pageSize

	kpis, err := s.queries.GetKpis(ctx, db.GetKpisParams{
		Column1: pgtype.Text{String: req.Search, Valid: true},
		Limit:   int32(pageSize),
		Offset:  int32(offset),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch kpis: %w", err)
	}

	totalCount, err := s.queries.GetKpisCount(ctx, pgtype.Text{String: req.Search, Valid: true})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to fetch kpi count: %w", err)
	}

	var pbKpis []*kpipb.Kpi
	for _, e := range kpis {
		pbKpis = append(pbKpis, &kpipb.Kpi{
			Id:     e.ID.String(),
			Name:   e.Name,
			Value:  float32(e.Value),
			Target: float32(e.Target),
			Day:    e.Day.Time.Format("2006-01-02"),
		})
	}

	return &kpipb.GetKpisResponse{
		Kpis:       pbKpis,
		TotalCount: int32(totalCount),
		Page:       page,
		PageSize:   pageSize,
	}, nil
}

// GetKpi handles fetching a single kpi by ID
func (s *KpiServiceServer) GetKpi(ctx context.Context, req *kpipb.GetKpiRequest) (*kpipb.GetKpiResponse, error) {
	// Parse and validate the UUID
	parsedUUID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UUID: %w", err)
	}

	// Create a pgtype.UUID instance
	kpiID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	kpi, err := s.queries.GetKpiByID(ctx, kpiID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "kpi not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to fetch kpi: %w", err)
	}

	return &kpipb.GetKpiResponse{
		Kpi: &kpipb.Kpi{
			Id:     kpi.ID.String(),
			Name:   kpi.Name,
			Value:  float32(kpi.Value),
			Target: float32(kpi.Target),
			Day:    kpi.Day.Time.Format("2006-01-02"),
		},
	}, nil
}

// PutKpi handles inserting a new kpi
func (s *KpiServiceServer) PutKpi(ctx context.Context, req *kpipb.PutKpiRequest) (*kpipb.PutKpiResponse, error) {
	var pgDate pgtype.Date

	// Use Scan to set the value
	err := pgDate.Scan(req.Day)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse Day field: %w", err)
	}

	kpi, err := s.queries.InsertKpi(ctx, db.InsertKpiParams{
		Name:   req.Name,
		Value:  float64(req.Value),
		Target: float64(req.Target),
		Day:    pgDate,
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to insert kpi: %w", err)
	}

	return &kpipb.PutKpiResponse{
		Kpi: &kpipb.Kpi{
			Id:     kpi.ID.String(),
			Name:   kpi.Name,
			Value:  float32(kpi.Value),
			Target: float32(kpi.Target),
			Day:    kpi.Day.Time.Format("2006-01-02"),
		},
	}, nil
}

// DeleteKpi handles deleting an kpi by ID
func (s *KpiServiceServer) DeleteKpi(ctx context.Context, req *kpipb.DeleteKpiRequest) (*kpipb.DeleteKpiResponse, error) {
	// Parse and validate the UUID                                        Kpi
	parsedUUID, err := uuid.Parse(req.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid UUID: %w", err)
	}

	// Create a pgtype.UUID instance
	kpiID := pgtype.UUID{
		Bytes: parsedUUID,
		Valid: true,
	}

	err = s.queries.DeleteKpi(ctx, kpiID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to delete kpi: %w", err)
	}
	return &kpipb.DeleteKpiResponse{}, nil
}
