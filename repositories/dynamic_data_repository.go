package repositories

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type DynamicDataRepository interface {
	Create(*models.DynamicData) error
	Update(id uuid.UUID, data *models.DynamicData) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*models.DynamicData, error)
	FindAll() ([]*models.DynamicData, error)
}
