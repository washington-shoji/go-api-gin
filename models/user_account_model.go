package models

import (
	"time"

	"github.com/google/uuid"
)

type UserAccount struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username"`
	Password  string     `json:"password"`
	Email     string     `json:"email"`
	FullName  string     `json:"fullName"`
	Role      string     `json:"role"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type UserAccountResponse struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	FullName string    `json:"fullName"`
}

type CreateUserAccountRequest struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
	Email    string `validate:"required" json:"email"`
	FullName string `validate:"required" json:"fullName"`
}

type UpdateUserAccountRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}
