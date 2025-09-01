package auth_service_entity

import (
	"time"
)

type Audit struct {
	ID        string    `json:"id" swaggerignore:"true"`
	CreatedAt time.Time `json:"created_at" swaggerignore:"true"`
	CreatedBy string    `json:"created_by" swaggerignore:"true"`
	UpdatedAt time.Time `json:"updated_at" swaggerignore:"true"`
	UpdatedBy string    `json:"updated_by" swaggerignore:"true"`
}
