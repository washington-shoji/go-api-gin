package handlers

import (
	"fmt"
	"net/http"
	"time"

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
	createEventReq := models.EventReq{}
	err := ctx.ShouldBindJSON(&createEventReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.EventService.Create(&createEventReq); err != nil {
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
	updateEventReq := models.EventReq{}
	err := ctx.ShouldBindJSON(&updateEventReq)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	eventID := ctx.Param("id")
	id, err := uuid.Parse(eventID)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.EventService.Update(id, &updateEventReq); err != nil {
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
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.EventService.Delete(id); err != nil {
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
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	result, err := handler.EventService.FindByID(id)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Event not found"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *EventHandler) GetAllEvents(ctx *gin.Context) {
	result, err := handler.EventService.FindAll()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *EventHandler) RenderCreateEventForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "event", gin.H{})
}

func (handler *EventHandler) PostCreateEventForm(ctx *gin.Context) {
	if err := ctx.Request.Form; err != nil {
		fmt.Println("request form", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	file, header, err := ctx.Request.FormFile("imageFile")
	if err != nil {
		fmt.Println("file, header", err)
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid file input"}})
		return
	}

	const layout = "2006-01-02T15:04"
	date := ctx.Request.FormValue("date")
	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid date input"}})
		return
	}

	registration := ctx.Request.FormValue("registration")
	parsedRegistration, err := time.Parse(layout, registration)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid registration input"}})
		return
	}

	createEvent := models.EventReq{
		Title:            ctx.Request.FormValue("title"),
		ShortDescription: ctx.Request.FormValue("shortDescription"),
		Description:      ctx.Request.FormValue("description"),
		ImageFile:        file,
		ImageHeader:      header,
		Date:             parsedDate,
		Registration:     parsedRegistration,
	}

	fmt.Println("createEvent", &createEvent)

	if createEvent != (models.EventReq{}) {
		if err := handler.EventService.Create(&createEvent); err != nil {
			fmt.Println("Error creating book", err)
			helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Could not create event"}})
			return
		}
	}

	// result, err := handler.EventService.FindAll()
	// if err != nil {
	// 	helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
	// 	return
	// }

	ctx.HTML(http.StatusOK, "home", gin.H{})
}
