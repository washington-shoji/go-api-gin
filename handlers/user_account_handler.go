package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/services"
)

type UserAccountHandler struct {
	UserAccountService services.UserAccountService
}

func NewUserAccountHandler(service services.UserAccountService) *UserAccountHandler {
	return &UserAccountHandler{
		UserAccountService: service,
	}
}

func (handler *UserAccountHandler) Create(ctx *gin.Context) {
	crateUsrAccReq := models.CreateUserAccountRequest{}
	err := ctx.ShouldBindJSON(&crateUsrAccReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{
			Status: http.StatusBadRequest,
			Error:  []string{"Invalid input"},
		})
		return
	}

	if err := handler.UserAccountService.Create(&crateUsrAccReq); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{
			Status: http.StatusBadRequest,
			Error:  []string{"Invalid input"},
		})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Created successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{
		Status: http.StatusOK, Data: resp,
	})
}
