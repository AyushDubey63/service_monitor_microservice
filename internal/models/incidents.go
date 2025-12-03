package models

import (
	"time"

	"github.com/google/uuid"
)

type Incident struct {
	ID         uuid.UUID `db:"id" json:"id"`
	ServiceID  uuid.UUID `db:"service_id" json:"service_id"`
	StartedAt  time.Time `db:"started_at" json:"started_at"`
	ResolvedAt *time.Time `db:"resolved_at" json:"resolved_at"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
