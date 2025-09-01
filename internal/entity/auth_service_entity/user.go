package auth_service_entity

type User struct {
	Audit
	Name        string       `json:"name"`
	PhoneNumber string       `json:"phone_number"`
	Email       string       `json:"email"`
	Username    string       `json:"username"`
	Password    string       `json:"password"`
	Image       string       `json:"image"`
	Address     string       `json:"address"`
	Role        Role         `json:"role"`
	Status      string       `json:"status"`
	Permissions []Permission `json:"permissions"`
}

type GetUsers struct {
	Count int    `json:"count"`
	Users []User `json:"users"`
}

type UserPK struct {
	ID string `json:"id"`
}

type ActivationModel struct {
	ID        string `json:"id"`
	Status    string `json:"status"`
	UpdatedBy string `json:"updated_by"`
}

type UpdateUser struct {
	Audit
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Image       string `json:"image"`
	Address     string `json:"address"`
	Role        Role   `json:"role"`
	Status      string `json:"status"`
}

type UserLogin struct {
	Audit
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
