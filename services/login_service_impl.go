package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/repositories"
)

type LoginServiceImp struct {
	UserAccountRepository repositories.UserAccountRepository
	Validate              *validator.Validate
}

func NewLoginService(userAccRep repositories.UserAccountRepository, validate *validator.Validate) LoginService {
	return &LoginServiceImp{
		UserAccountRepository: userAccRep,
		Validate:              validate,
	}
}

// Login implements LoginService.
func (lgn *LoginServiceImp) Login(username string, password string) (login *models.LoginResponse, error error) {

	result, err := lgn.UserAccountRepository.FindByUsername(username)
	if err != nil {
		return nil, err
	}

	if err := helpers.VerifyPassword(result.Password, password); err != nil {
		return nil, err
	}

	token, err := helpers.GenerateToken(result.Username)
	if err != nil {
		return nil, err
	}

	resp := &models.LoginResponse{
		Username: result.Username,
		Token:    token,
	}

	return resp, nil
}
