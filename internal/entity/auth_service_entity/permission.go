package auth_service_entity

type Permission struct {
	Audit
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
}

type PermissionPK struct {
	ID string `json:"id"`
}

type GetPermissions struct {
	Count       int          `json:"count"`
	Permissions []Permission `json:"permissions"`
}

type UpdatePermission struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Description string `json:"description"`
	UpdatedBy   string `json:"updated_by"`
}

type RolePermission struct {
	Audit
	RoleID string `json:"role_id"`
	PermID string `json:"perm_id"`
}

type UserPermission struct {
	Audit
	UserID string `json:"user_id"`
	PermID string `json:"perm_id"`
}
