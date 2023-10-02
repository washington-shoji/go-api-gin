package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type DynamicDataService interface {
	Create(*models.DynamicDataReq) error
	Update(id uuid.UUID, data *models.DynamicDataReq) error
	Delete(id uuid.UUID) error
	FindById(id uuid.UUID) (*models.DynamicDataRes, error)
	FindAll() ([]*models.DynamicDataRes, error)
}
