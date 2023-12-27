package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/services"
)

type EventHandler struct {
	EventService services.EventService
}

func NewEventHandler(service services.EventService) *EventHandler {
	return &EventHandler{
		EventService: service,
	}
}

func (handler *EventHandler) Create(ctx *gin.Context) {
	if err := ctx.Request.Form; err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	file, header, err := ctx.Request.FormFile("imageHeader")
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid file input"}})
		return
	}

	eventReqJson := models.EventReq{}
	if err := ctx.ShouldBind(&eventReqJson); err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	createEventReq := models.EventReq{
		ImageHeader:  header,
		ImageFile:    file,
		EventDetails: eventReqJson.EventDetails,
	}

	if err := handler.EventService.Create(&createEventReq); err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}
	resp := SuccessMessage{}
	resp.Message = "Created successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{
		Status: http.StatusOK, Data: resp,
	})
}

func (handler *EventHandler) Update(ctx *gin.Context) {
	eventID := ctx.Param("id")
	id, err := uuid.Parse(eventID)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	var eventReqJson models.EventReq

	if err := ctx.ShouldBind(&eventReqJson); err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	updateEventReq := models.EventReq{
		EventDetails: eventReqJson.EventDetails,
	}

	// Handle file upload if present
	if file, header, err := ctx.Request.FormFile("imageHeader"); err == nil {
		defer file.Close()
		updateEventReq.ImageHeader = header
		updateEventReq.ImageFile = file
	} else if err != http.ErrMissingFile {
		// Handle file upload errors other than missing file
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid file input"}})
		return
	}

	if err := handler.EventService.Update(id, &updateEventReq); err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Event not found"}})
		return
	}

	resp := SuccessMessage{
		Message: "Updated successfully",
	}
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})

}

func (handler *EventHandler) Delete(ctx *gin.Context) {
	eventID := ctx.Param("id")
	id, err := uuid.Parse(eventID)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.EventService.Delete(id); err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Event does not exist"}})
		return
	}

	resp := SuccessMessage{
		Message: "Deleted successfully",
	}

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *EventHandler) GetEventByID(ctx *gin.Context) {
	eventID := ctx.Param("id")
	id, err := uuid.Parse(eventID)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	result, err := handler.EventService.FindByID(id)
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Event not found"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *EventHandler) GetAllEvents(ctx *gin.Context) {
	result, err := handler.EventService.FindAll()
	if err != nil {
		log.Printf("Error in Handler: %s", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}
