package models

import "github.com/google/uuid"

type ServiceCheckRule struct {
	ID           uuid.UUID `db:"id" json:"id"`
	ServiceID    uuid.UUID `db:"service_id" json:"service_id"`
	ExpectedCode int       `db:"expected_code" json:"expected_code"`
	ExpectedBody string    `db:"expected_body" json:"expected_body"`
}
