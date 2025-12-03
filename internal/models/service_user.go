package models

import (
	"time"

	"github.com/google/uuid"
)

type ServiceUser struct {
	UserID    int       `db:"user_id" json:"user_id"`
	ServiceID uuid.UUID `db:"service_id" json:"service_id"`
	Role      string    `db:"role" json:"role"`
	Status    string    `db:"status" json:"status"`
	InvitedBy int       `db:"invited_by" json:"invited_by"`
	JoinedAt  time.Time `db:"joined_at" json:"joined_at"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}
