package models

import (
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID               uuid.UUID  `json:"id"`
	Title            string     `json:"title"`
	ShortDescription string     `json:"shortDescription"`
	Description      string     `json:"description"`
	ImageUrl         string     `json:"imageUrl"`
	Date             time.Time  `json:"date"`
	Registration     time.Time  `json:"registration"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}

type EventReq struct {
	Title            string                `json:"title" form:"title"`
	ShortDescription string                `json:"shortDescription" form:"shortDescription"`
	Description      string                `json:"description" form:"description"`
	ImageFile        multipart.File        `json:"imageFile" form:"imageFile"`
	ImageHeader      *multipart.FileHeader `json:"imageHeader" form:"imageHeader"`
	ImageUrl         string                `json:"imageUrl" form:"imageUrl"`
	Date             time.Time             `json:"date" form:"date"`
	Registration     time.Time             `json:"registration" form:"registration"`
}

type EventRes struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	ImageUrl         string    `json:"imageUrl"`
	Date             time.Time `json:"date"`
	Registration     time.Time `json:"registration"`
}
