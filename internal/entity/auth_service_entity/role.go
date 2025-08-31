package auth_service_entity

type Role struct {
	Audit
	Name        string       `json:"name"`
	Status      bool         `json:"status"`
	Description string       `json:"description"`
	Permissions []Permission `json:"permissions"`
}

type RolePK struct {
	ID string `json:"id"`
}

type GetRoles struct {
	Count int    `json:"count"`
	Roles []Role `json:"roles"`
}

type UpdateRole struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
	UpdatedBy   string `json:"updated_by"`
}
