package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type UserAccountService interface {
	Create(usrAcc *models.CreateUserAccountRequest) (error error)
	Update(id uuid.UUID, usrAcc *models.UpdateUserAccountRequest) (error error)
	Delete(id uuid.UUID) (error error)
	FindByID(id uuid.UUID) (usrAcc *models.UserAccountResponse, error error)
	FindAll() (usrAcc []*models.UserAccountResponse, error error)
}
