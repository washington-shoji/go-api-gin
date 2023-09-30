package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/services"
)

type TableTopGameHandler struct {
	Service services.TableTopGameService
}

func NewTableTopGameHandler(service services.TableTopGameService) *TableTopGameHandler {
	return &TableTopGameHandler{
		Service: service,
	}
}

func (service *TableTopGameHandler) Create(ctx *gin.Context) {
	createTblGameReq := models.TableTopGameReq{}
	err := ctx.ShouldBindJSON(&createTblGameReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := service.Service.Create(&createTblGameReq); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	resp := SuccessMessage{
		Message: "Created successfully",
	}

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *TableTopGameHandler) Update(ctx *gin.Context) {
	uptTblGameReq := models.TableTopGameReq{}
	err := ctx.ShouldBindJSON(&uptTblGameReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	tblId := ctx.Param("id")
	id, err := uuid.Parse(tblId)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.Service.Update(id, &uptTblGameReq); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Table top game not found"}})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Updated successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *TableTopGameHandler) Delete(ctx *gin.Context) {
	uptTblGameReq := models.TableTopGameReq{}
	err := ctx.ShouldBindJSON(&uptTblGameReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	tblId := ctx.Param("id")
	id, err := uuid.Parse(tblId)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.Service.Delete(id); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Table top game not found"}})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Deleted successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *TableTopGameHandler) FindByID(ctx *gin.Context) {

	tblId := ctx.Param("id")
	id, err := uuid.Parse(tblId)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	result, err := handler.Service.FindByID(id)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Table top game not found"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *TableTopGameHandler) FindAll(ctx *gin.Context) {
	result, err := handler.Service.FindAll()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}
