package models

import (
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type CreateBookRequest struct {
	Title       string `validate:"required" json:"title"`
	Description string `validate:"required" json:"description"`
}

type UpdateBookRequest struct {
	ID          uuid.UUID  `validate:"required" json:"id"`
	Title       string     `validate:"required,min=1,max=50" json:"title"`
	Description string     `validate:"required,min=1,max=500" json:"description"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type DeleteBookRequest struct {
	ID        uuid.UUID  `validate:"required" json:"id"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type FindByIDBookRequest struct {
	ID uuid.UUID `validate:"required" json:"id"`
}
