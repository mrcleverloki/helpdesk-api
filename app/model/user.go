package model

import (
	"time"
)

type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`      // Tidak pernah keluar sebagai JSON
	Role      string    `json:"role"`   // Role user (user/admin)
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ReplaceUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
}

type PatchUserRequest struct {
	Username *string `json:"username"`
	Email    *string `json:"email"`
	IsActive *bool   `json:"is_active"`
}

type ListQuery struct {
	Page     int
	Limit    int
	Search   string
	Sort     string
	Order    string
	IsActive *bool
}

type Meta struct {
	Page       int `json:"page"`
	Limit      int `json:"limit"`
	Total      int `json:"total"`
	TotalPages int `json:"total_pages"`
}

type WebResponse struct {
	Success bool              `json:"success"`
	Message string            `json:"message"`
	Data    any               `json:"data,omitempty"`
	Meta    *Meta             `json:"meta,omitempty"`
	Errors  map[string]string `json:"errors,omitempty"`
}
