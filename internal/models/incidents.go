package models

import (
	"time"

	"github.com/google/uuid"
)

type Incident struct {
	ID         uuid.UUID `db:"id" json:"id"`
	ServiceID  uuid.UUID `db:"service_id" json:"service_id"`
	Status     string `db:"status" json:"status"`
	ErrorMessage string `db:"error_message" json:"error_message"`
	TriggerStatusCode int `db:"trigger_status_code" json:"trigger_status_code"`
	TriggerLatencyMS int `db:"trigger_latency_ms" json:"trigger_latency_ms"`
	DurationMS int `db:"duration_ms" json:"duration_ms"`
	StartedAt  time.Time `db:"started_at" json:"started_at"`
	ResolvedAt *time.Time `db:"resolved_at" json:"resolved_at"`
	Notes string `db:"notes" json:"notes"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
}
