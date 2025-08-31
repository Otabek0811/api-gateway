package auth_service_entity

import "time"

type Session struct {
	Audit
	UserID       string    `json:"user_id"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	IsActive     bool      `json:"is_active"`
	ExpiresAt    time.Time `json:"expires_at"`
	LastActiveAt time.Time `json:"last_active_at"`
	Platform     string    `json:"platform"`
}

type SessionList struct {
	Items []Session `json:"sessions"`
	Count int       `json:"count"`
}

type SessionPK struct {
	ID string `json:"id"`
}

type DeactivateSession struct {
	ID        string `json:"id"`
	UpdatedBy string `json:"updated_by"`
}
