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
	ImagePublicId    string     `json:"publicId"`
	Date             time.Time  `json:"date"`
	Registration     time.Time  `json:"registration"`
	CreatedAt        time.Time  `json:"createdAt"`
	UpdatedAt        *time.Time `json:"updatedAt"`
}

type EventReqJson struct {
	Title            string    `json:"title"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	ImageUrl         *string   `json:"imageUrl"`
	ImagePublicId    *string   `json:"publicId"`
	Date             time.Time `json:"date"`
	Registration     time.Time `json:"registration"`
}

type EventReq struct {
	ImageHeader  *multipart.FileHeader `form:"imageHeader"`
	ImageFile    multipart.File        `form:"imageFile"`
	EventDetails EventReqJson          `form:"eventDetails"`
}

type EventRes struct {
	ID               uuid.UUID `json:"id"`
	Title            string    `json:"title"`
	ShortDescription string    `json:"shortDescription"`
	Description      string    `json:"description"`
	ImageUrl         string    `json:"imageUrl"`
	ImagePublicId    *string   `json:"publicId"`
	Date             time.Time `json:"date"`
	Registration     time.Time `json:"registration"`
}
