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

type SuccessMessage struct {
	Message string `json:"message"`
}

type BookHandler struct {
	BookService services.BookService
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
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.BookService.Create(&createBookRequest); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Created successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})

}

func (handler *BookHandler) Update(ctx *gin.Context) {
	updateBookRequest := models.UpdateBookRequest{}
	err := ctx.ShouldBindJSON(&updateBookRequest)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	id := ctx.Param("id")
	bookID, err := uuid.Parse(id)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	updateBookRequest.ID = bookID
	if err := handler.BookService.Update(bookID, &updateBookRequest); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Book not found"}})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Updated successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *BookHandler) Delete(ctx *gin.Context) {
	bookID := ctx.Param("id")
	id, err := uuid.Parse(bookID)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	if err := handler.BookService.Delete(id); err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Book does not exist"}})
		return
	}

	resp := SuccessMessage{}
	resp.Message = "Deleted successfully"

	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: resp})
}

func (handler *BookHandler) GetBookByID(ctx *gin.Context) {
	findByIDRequest := models.FindByIDBookRequest{}

	bookID := ctx.Param("id")
	id, err := uuid.Parse(bookID)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Invalid input"}})
		return
	}

	findByIDRequest.ID = id

	result, err := handler.BookService.FindByID(id, &findByIDRequest)
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Book not found"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *BookHandler) GetAllBooks(ctx *gin.Context) {

	result, err := handler.BookService.FindAll()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	ctx.Header("Content-Type", "application/json")
	helpers.WebResponseSuccessHandler(ctx, helpers.ResponseSuccess{Status: http.StatusOK, Data: result})
}

func (handler *BookHandler) RenderBookForm(ctx *gin.Context) {
	books, err := handler.BookService.FindAll()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Books": books,
	})
}

func (handler *BookHandler) CreateBookForm(ctx *gin.Context) {
	if err := ctx.Request.Form; err != nil {
		fmt.Println("Error parsing form", err)
	}

	createBook := models.CreateBookRequest{
		Title:       ctx.Request.FormValue("title"),
		Description: ctx.Request.FormValue("description"),
	}

	if createBook != (models.CreateBookRequest{}) {
		if err := handler.BookService.Create(&createBook); err != nil {
			fmt.Println("Error creating book", err)
			helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadRequest, Error: []string{"Could not create book"}})
			return
		}
	}

	books, err := handler.BookService.FindAll()
	if err != nil {
		helpers.WebResponseError(ctx, helpers.ResponseError{Status: http.StatusBadGateway, Error: []string{"Server Error"}})
		return
	}

	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"Books": books,
	})
}

func (handler *BookHandler) RenderPartials(ctx *gin.Context) {
	data := gin.H{}

	ctx.HTML(http.StatusOK, "index.html", data)
}
