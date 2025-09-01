package auth_service_entity

type Id struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
}

type OrderBy struct {
	Column string `json:"column"`
	Order  string `json:"order"`
}

type Filter struct {
	Column string `json:"column"`
	Type   string `json:"type"` // eq, ne, gt, gte, lt, lte, search
	Value  string `json:"value"`
}

type GetListFilter struct {
	Page    int       `json:"page"`
	Limit   int       `json:"limit"`
	Filters []Filter  `json:"filters"`
	OrderBy []OrderBy `json:"order_by"`
}

type UpdateFieldItem struct {
	Column string      `json:"column"`
	Value  interface{} `json:"value"`
}

type UpdateFieldRequest struct {
	Filter []Filter          `json:"filter"`
	Items  []UpdateFieldItem `json:"items"`
}

type RowsEffected struct {
	RowsEffected int `json:"rows_effected"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// ErrorResponse ...
type ErrorResponse struct {
	Code    string      `json:"code"`
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}
