package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type EventService interface {
	Create(event *models.EventRequest) error
	Update(id uuid.UUID, event *models.EventRequest) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.EventResponse, error)
	FindAll() (events []*models.EventResponse, error error)
}
