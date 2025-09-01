package auth_service_entity

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Platform string `json:"platform"`
}

type LoginResponse struct {
	UserID       string  `json:"user_id"`
	UserName     string  `json:"user_name"`
	AccessToken  string  `json:"access_token"`
	RefreshToken string  `json:"refresh_token"`
	Session      Session `json:"session"`
}

// HasAccessModel ...
type HasAccessModel struct {
	UserID    string `json:"user_id"`
	RoleID    string `json:"user_role_id"`
	UserRole  string `json:"user_role"`
	HasAccess bool   `json:"has_access"`
}

// HasErrorModel ...
type HasErrorModel struct {
	Error     interface{} `json:"error"`
	HasAccess bool        `json:"has_access"`
}
