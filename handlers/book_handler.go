package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/helpers"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/services"
)

type BookHandler struct {
	BookService services.BookService
}

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

func NewBookHandler(service services.BookService) *BookHandler {
	return &BookHandler{
		BookService: service,
	}
}

func (handler *BookHandler) Create(ctx *gin.Context) {
	createBookRequest := models.CreateBookRequest{}
	err := ctx.ShouldBindJSON(&createBookRequest)
	if err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.BookService.Create(&createBookRequest); err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	resp := helpers.WebResponse(http.StatusOK, "Created successfully")

	ctx.JSON(http.StatusOK, resp)
}

func (handler *BookHandler) Update(ctx *gin.Context) {
	updateBookRequest := models.UpdateBookRequest{}
	err := ctx.ShouldBindJSON(&updateBookRequest)
	if err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	id := ctx.Param("id")
	bookID, err := uuid.Parse(id)
	if err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	updateBookRequest.ID = bookID
	if err := handler.BookService.Update(bookID, &updateBookRequest); err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusNotFound, Error: []string{"Book not found"}})
		return
	}

	resp := helpers.WebResponse(http.StatusOK, "Updated successfully")

	ctx.JSON(http.StatusOK, resp)
}

func (handler *BookHandler) Delete(ctx *gin.Context) {
	bookID := ctx.Param("id")
	id, err := uuid.Parse(bookID)
	if err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.BookService.Delete(id); err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusNotFound, Error: []string{"Book does not exist"}})
		return
	}

	resp := helpers.WebResponse(http.StatusOK, "Deleted successfully")

	ctx.JSON(http.StatusOK, resp)
}

func (handler *BookHandler) GetBookByID(ctx *gin.Context) {
	findByIDRequest := models.FindByIDBookRequest{}

	bookID := ctx.Param("id")
	id, err := uuid.Parse(bookID)
	if err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	findByIDRequest.ID = id

	result, err := handler.BookService.FindByID(id, &findByIDRequest)
	if err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusNotFound, Error: []string{"Book not found"}})
		return
	}

	resp := helpers.WebResponse(http.StatusOK, result)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}

func (handler *BookHandler) GetAllBooks(ctx *gin.Context) {

	result, err := handler.BookService.FindAll()
	if err != nil {
		helpers.WebResponseHandler(ctx, helpers.Response{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	resp := helpers.WebResponse(http.StatusOK, result)

	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, resp)
}
