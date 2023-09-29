package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type TableTopGameService interface {
	Create(game *models.TableTopGameReq) error
	Update(id uuid.UUID, game *models.TableTopGameReq) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.TableTopGameResp, error)
	FindAll() ([]*models.TableTopGameResp, error)
}
