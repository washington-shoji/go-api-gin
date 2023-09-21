package services

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type BookService interface {
	Create(book *models.CreateBookRequest) (error error)
	Update(id uuid.UUID, book *models.UpdateBookRequest) (error error)
	Delete(id uuid.UUID, book *models.DeleteBookRequest) (error error)
	FindByID(id uuid.UUID, bkr *models.FindByIDBookRequest) (book *models.Book, error error)
	FindAll() (books []*models.Book, error error)
}
