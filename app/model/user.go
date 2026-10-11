package model

import "time"

// User merepresentasikan entitas user di database.
type User struct {
	ID        int       `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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
}

// AssignRoleRequest dipakai endpoint PATCH /users/:id/role.
type AssignRoleRequest struct {
	Role string `json:"role"`
}
