package repositories

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type UserAccountRepository interface {
	Create(usrAcc *models.UserAccount) (error error)
	Update(usrAcc *models.UserAccount) (error error)
	Delete(usrAcc *models.UserAccount) (error error)
	FindByID(id uuid.UUID) (usrAcc *models.UserAccount, error error)
	FindByUsername(username string) (usrAcc *models.UserAccount, error error)
	FindAll() (usrAcc []*models.UserAccount, error error)
}
