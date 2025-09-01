package auth_service_entity

import (
	"time"
)

type Audit struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	CreatedBy string    `json:"created_by"`
	UpdatedAt time.Time `json:"updated_at"`
	UpdatedBy string    `json:"updated_by"`
}
