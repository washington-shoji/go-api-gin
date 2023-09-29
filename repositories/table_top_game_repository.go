package repositories

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type TableTopGameRepository interface {
	Create(game *models.TableTopGame) error
	Update(id uuid.UUID, game *models.TableTopGame) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.TableTopGame, error)
	FindAll() ([]*models.TableTopGame, error)
}
