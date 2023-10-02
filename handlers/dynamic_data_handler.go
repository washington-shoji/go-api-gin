package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/services"
)

type DynamicDataHandler struct {
	Service services.DynamicDataService
}

func NewDynamicDataHandler(service services.DynamicDataService) *DynamicDataHandler {
	return &DynamicDataHandler{
		Service: service,
	}
}

func (handler *DynamicDataHandler) Create(ctx *gin.Context) {
	rawJson, err := ctx.GetRawData()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	dynDataReq := models.DynamicDataReq{
		Data: rawJson,
	}

	if err := handler.Service.Create(&dynDataReq); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	resp := SuccessMessage{
		Message: "Created successfully",
	}

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *DynamicDataHandler) FindAll(ctx *gin.Context) {
	result, err := handler.Service.FindAll()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}
