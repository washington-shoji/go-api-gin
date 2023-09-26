package services

import "github.com/washington-shoji/gin-api/models"

type LoginService interface {
	Login(username string, password string) (login *models.LoginResponse, error error)
}
