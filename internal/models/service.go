package models

import (
	"time"

	"github.com/google/uuid"
)

type MonitorService struct {
	ID              uuid.UUID `db:"id" json:"id"`
	Name            string    `db:"name" json:"name"`
	Description     string    `db:"description" json:"description"`
	Endpoint        string    `db:"endpoint" json:"endpoint"`
	HttpMethod      string    `db:"http_method" json:"http_method"`
	Environment     string    `db:"environment" json:"environment"`
	IntervalSeconds int       `db:"interval_seconds" json:"interval_seconds"`
	RetryCount      int       `db:"retry_count" json:"retry_count"`
	TimeoutMS  int       `db:"timeout_ms" json:"timeout_ms"`
	IsPublic        bool      `db:"is_public" json:"is_public"`
	AlertEnabled    bool      `db:"alert_enabled" json:"alert_enabled"`
	AlertChannel    string    `db:"alert_channel" json:"alert_channel"`
	CreatedBy       int       `db:"created_by" json:"created_by"`
	TenantID        uuid.UUID `db:"tenant_id" json:"tenant_id"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
