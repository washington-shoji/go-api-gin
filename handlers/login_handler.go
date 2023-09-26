package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/services"
)

type LoginHandler struct {
	LoginService services.LoginService
}

func NewLoginHandler(service services.LoginService) *LoginHandler {
	return &LoginHandler{
		LoginService: service,
	}
}

func (handler *LoginHandler) Login(ctx *gin.Context) {
	loginReq := models.LoginRequest{}
	err := ctx.ShouldBindJSON(&loginReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{
			Status: http.StatusBadRequest,
			Error:  []string{"Invalid input"},
		})
		return
	}

	result, err := handler.LoginService.Login(loginReq.Username, loginReq.Password)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusForbidden, Error: []string{"Not authorized"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}
