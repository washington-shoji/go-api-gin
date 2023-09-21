package repositories

import (
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type BookRepository interface {
	Create(book *models.Book) (error error)
	Update(book *models.Book) (error error)
	Delete(book *models.Book) (error error)
	FindByID(id uuid.UUID) (book *models.Book, error error)
	FindAll() (books []*models.Book, error error)
}
