package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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

func (handler *UserAccountHandler) GetUserAccountByID(ctx *gin.Context) {
	usrId := ctx.Param("id")
	id, err := uuid.Parse(usrId)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	result, err := handler.UserAccountService.FindByID(id)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"user account not found"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *UserAccountHandler) GetAllUserAccounts(ctx *gin.Context) {
	result, err := handler.UserAccountService.FindAll()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{
			Status: http.StatusBadGateway,
			Error:  []string{"Server Error"},
		})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{
		Status: http.StatusOK,
		Data:   result,
	})
}

func (handler *UserAccountHandler) Update(ctx *gin.Context) {
	updUsrAccReq := models.UpdateUserAccountRequest{}
	err := ctx.ShouldBindJSON(&updUsrAccReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	id := ctx.Param("id")
	usrId, err := uuid.Parse(id)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.UserAccountService.Update(usrId, &updUsrAccReq); err != nil {
		fmt.Println(err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"User account not found"}})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Updated successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *UserAccountHandler) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	usrId, err := uuid.Parse(id)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.UserAccountService.Delete(usrId); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"User account does not exist"}})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Deleted successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}
