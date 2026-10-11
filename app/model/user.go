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

// ListQuery menampung parameter query untuk pagination, pencarian, dan filtering.
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

// Meta menampung informasi pagination.
type Meta struct {
	Page      int `json:"page"`
	PerPage   int `json:"per_page"`
	TotalData int `json:"total_data"`
	TotalPage int `json:"total_page"`
}

// CreateUserRequest dipakai saat membuat user baru.
type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// ReplaceUserRequest dipakai saat memperbarui seluruh data user (PUT).
type ReplaceUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

// PatchUserRequest dipakai saat memperbarui sebagian data user (PATCH).
type PatchUserRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	IsActive *bool   `json:"is_active"`
}

// AssignRoleRequest dipakai endpoint PATCH /users/:id/role.
type AssignRoleRequest struct {
	Role string `json:"role"`
}
