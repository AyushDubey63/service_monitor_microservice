package models

import (
	"time"

	"github.com/google/uuid"
)

type ServiceHealthLog struct {
	ID        uuid.UUID `db:"id" json:"id"`
	ServiceID uuid.UUID `db:"service_id" json:"service_id"`
	Status    string    `db:"status" json:"status"` 
	StatusCode int 		`db:"status_code" json:"status_code"`
	LatencyMs int       `db:"latency_ms" json:"latency_ms"`
	Error     string    `db:"error" json:"error"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
