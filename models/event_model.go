package models

import (
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
	DeletedAt        *time.Time `json:"deletedAt"`
}
