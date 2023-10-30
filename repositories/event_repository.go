package repositories

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type EventRepository interface {
	Create(event models.Event) error
	Update(event models.Event) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.Event, error)
	FindAll() ([]*models.Event, error)
}
