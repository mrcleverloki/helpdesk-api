package model

import "time"

// User merepresentasikan entitas user di database.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ListQuery menampung parameter query untuk pagination, pencarian,dan filtering.
type ListQuery struct {
	Page     int    `query:"page"`
	PerPage  int    `query:"per_page"`
	Limit    int    `query:"limit"`
	Offset   int    `query:"offset"`
	Search   string `query:"search"`
	IsActive *bool  `query:"is_active"`
	Sort     string `query:"sort"`
	Order    string `query:"order"`
}

// WebResponse membungkus format JSON response standar.
type WebResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    interface{}       `json:"data,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
	Meta    *Meta             `json:"meta,omitempty"`
}

// ErrorResponse membungkus format JSON response kegagalan yang seragam.
type ErrorResponse struct {
	Success   bool              `json:"success"`
	Code      string            `json:"code"`
	Message   string            `json:"message"`
	Fields    map[string]string `json:"fields,omitempty"`
	RequestID string            `json:"request_id,omitempty"`
}

// Meta menampung informasi pagination.
type Meta struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

// CreateUserRequest dipakai saat membuat user baru.
type CreateUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Email    string `json:"email"    validate:"required,email,max=120"`
	Password string `json:"password" validate:"required,min=8,max=72,strongpassword"`
	Role     string `json:"role"     validate:"required,oneof=admin staff user"`
}

// ReplaceUserRequest dipakai saat memperbarui seluruh data user (PUT).
type ReplaceUserRequest struct {
	Username string `json:"username" validate:"required,min=3,max=30,alphanum"`
	Email    string `json:"email"    validate:"required,email,max=120"`
}

// PatchUserRequest dipakai saat memperbarui sebagian data user (PATCH).
type PatchUserRequest struct {
	Username *string `json:"username,omitempty" validate:"omitnil,min=3,max=30,alphanum"`
	Email    *string `json:"email,omitempty"    validate:"omitnil,email,max=120"`
	IsActive *bool   `json:"is_active,omitempty"`
}

// AssignRoleRequest dipakai endpoint PATCH /users/:id/role.
type AssignRoleRequest struct {
	Role string `json:"role"`
}
