package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type BookService interface {
	Create(book *models.CreateBookRequest) (error error)
	Update(id uuid.UUID, book *models.UpdateBookRequest) (error error)
	Delete(id uuid.UUID) (error error)
	FindByID(id uuid.UUID, bkr *models.FindByIDBookRequest) (book *models.BookResponse, error error)
	FindAll() (books []*models.BookResponse, error error)
}
