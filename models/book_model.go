package models

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Book struct {
	ID          uuid.UUID  `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	ImageUrl    string     `json:"imageUrl"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
	DeletedAt   *time.Time `json:"deletedAt"`
}

type BookResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	ImageUrl    string    `json:"imageUrl"`
}

type CreateBookRequest struct {
	Title       string                `validate:"required" json:"title"`
	Description string                `validate:"required" json:"description"`
	ImageFile   multipart.File        `json:"imageFile"`
	ImageHeader *multipart.FileHeader `json:"imageHeader"`
	ImageUrl    string                `json:"imageUrl"`
}

type UpdateBookRequest struct {
	ID          uuid.UUID             `validate:"required" json:"id"`
	Title       string                `validate:"required,min=1,max=50" json:"title"`
	Description string                `validate:"required,min=1,max=500" json:"description"`
	UpdatedAt   *time.Time            `json:"updatedAt"`
	ImageFile   multipart.File        `json:"imageFile"`
	ImageHeader *multipart.FileHeader `json:"imageHeader"`
	ImageUrl    string                `json:"imageUrl"`
}

type DeleteBookRequest struct {
	ID        uuid.UUID  `validate:"required" json:"id"`
	DeletedAt *time.Time `json:"deletedAt"`
}

type FindByIDBookRequest struct {
	ID uuid.UUID `validate:"required" json:"id"`
}
