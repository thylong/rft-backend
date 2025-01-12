// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	// Delete an event by ID
	DeleteEvent(ctx context.Context, eventID pgtype.UUID) error
	// Delete a kpi by ID
	DeleteKpi(ctx context.Context, id pgtype.UUID) error
	// Delete an okr by ID
	DeleteOkr(ctx context.Context, id pgtype.UUID) error
	// Get a single event by ID
	GetEventByID(ctx context.Context, eventID pgtype.UUID) (Event, error)
	// events.sql
	// Get events with pagination and optional search
	GetEvents(ctx context.Context, arg GetEventsParams) ([]Event, error)
	// Get total count for pagination
	GetEventsCount(ctx context.Context, dollar_1 pgtype.Text) (int64, error)
	// Get a single kpi by ID
	GetKpiByID(ctx context.Context, id pgtype.UUID) (Kpi, error)
	// kpis.sql
	// Get kpis
	GetKpis(ctx context.Context, arg GetKpisParams) ([]Kpi, error)
	// Get total count for pagination
	GetKpisCount(ctx context.Context, dollar_1 pgtype.Text) (int64, error)
	// Get a single okr by ID
	GetOkrByID(ctx context.Context, id pgtype.UUID) (GetOkrByIDRow, error)
	// okrs.sql
	// Get okrs with pagination and optional search
	GetOkrs(ctx context.Context, arg GetOkrsParams) ([]GetOkrsRow, error)
	// Get total count for pagination
	GetOkrsCount(ctx context.Context, dollar_1 pgtype.Text) (int64, error)
	// Insert a new event
	InsertEvent(ctx context.Context, arg InsertEventParams) (Event, error)
	// Insert a new key result
	InsertKeyResult(ctx context.Context, arg InsertKeyResultParams) (OkrKr, error)
	// Insert a new kpi
	InsertKpi(ctx context.Context, arg InsertKpiParams) (Kpi, error)
	// Insert a new okr
	InsertOkr(ctx context.Context, arg InsertOkrParams) (InsertOkrRow, error)
}

var _ Querier = (*Queries)(nil)
