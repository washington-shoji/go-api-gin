package models

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID               uuid.UUID  `json:"id"`
	Title            string     `json:"title"`
	ShortDescription string     `json:"short-description"`
	Description      string     `json:"description"`
	ImageUrl         string     `json:"image-url"`
	Date             time.Time  `json:"data"`
	Registration     time.Time  `json:"registration"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}

type EventReq struct {
	Title            string                `json:"title" form:"title"`
	ShortDescription string                `json:"short-description" form:"short-description"`
	Description      string                `json:"description" form:"description"`
	ImageFile        multipart.File        `json:"imageFile" form:"imageFile"`
	ImageHeader      *multipart.FileHeader `json:"imageHeader" form:"imageHeader"`
	ImageUrl         string                `json:"image-url" form:"image-url"`
	Date             time.Time             `json:"data" form:"data"`
	Registration     time.Time             `json:"registration" form:"registration"`
}

type EventRes struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"short-description"`
	Description      string    `json:"description"`
	ImageUrl         string    `json:"image-url"`
	Date             time.Time `json:"data"`
	Registration     time.Time `json:"registration"`
}
