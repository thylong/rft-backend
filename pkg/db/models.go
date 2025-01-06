// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Event struct {
	EventID      pgtype.UUID      `json:"event_id"`
	EventPrivacy int32            `json:"event_privacy"`
	Name         string           `json:"name"`
	Description  string           `json:"description"`
	Type         string           `json:"type"`
	Department   string           `json:"department"`
	Regions      []string         `json:"regions"`
	Tags         []string         `json:"tags"`
	StartAt      pgtype.Timestamp `json:"start_at"`
}

type Okr struct {
	ID          pgtype.UUID `json:"id"`
	Name        string      `json:"name"`
	Number      int32       `json:"number"`
	Year        int32       `json:"year"`
	Description string      `json:"description"`
}

type OkrKr struct {
	ID          pgtype.UUID `json:"id"`
	OkrID       pgtype.UUID `json:"okr_id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Sponsor     string      `json:"sponsor"`
	Kpis        string      `json:"kpis"`
	Scope       string      `json:"scope"`
	Initiatives string      `json:"initiatives"`
}
