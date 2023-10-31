package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type EventService interface {
	Create(event *models.EventReq) error
	Update(id uuid.UUID, event *models.EventReq) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.EventRes, error)
	FindAll() (events []*models.EventRes, error error)
}
