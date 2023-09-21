package services

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
	"github.com/washington-shoji/gin-api/repositories"
)

type BookServiceImpl struct {
	BookRepository repositories.BookRepository
	Validate       *validator.Validate
}

func NewBookService(bookRepository repositories.BookRepository, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		Validate:       validate,
	}
}

// Create implements BookService.
func (b *BookServiceImpl) Create(book *models.CreateBookRequest) (error error) {
	err := b.Validate.Struct(book)
	if err != nil {
		return err
	}

	id := uuid.New()
	time := time.Now()

	bookModel := models.Book{
		ID:          id,
		Title:       book.Title,
		Description: book.Description,
		CreatedAt:   time,
	}

	b.BookRepository.Create(&bookModel)
	return nil
}

// Delete implements BookService.
func (b *BookServiceImpl) Delete(id uuid.UUID) (error error) {

	time := time.Now()
	bookModel := models.Book{
		ID:        id,
		DeletedAt: &time,
	}

	if err := b.BookRepository.Delete(&bookModel); err != nil {
		return err
	}
	return nil
}

// FindAll implements BookService.
func (b *BookServiceImpl) FindAll() (books []*models.Book, error error) {
	books, err := b.BookRepository.FindAll()
	if err != nil {
		return nil, err
	}
	return books, nil

}

// FindByID implements BookService.
func (b *BookServiceImpl) FindByID(id uuid.UUID, bk *models.FindByIDBookRequest) (book *models.Book, error error) {

	err := b.Validate.Struct(bk)
	if err != nil {
		return nil, err
	}

	bookModel := &models.FindByIDBookRequest{
		ID: id,
	}

	result, err := b.BookRepository.FindByID(bookModel.ID)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Update implements BookService.
func (b *BookServiceImpl) Update(id uuid.UUID, book *models.UpdateBookRequest) (error error) {
	err := b.Validate.Struct(book)
	if err != nil {
		return err
	}

	time := time.Now()
	bookModel := models.Book{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		UpdatedAt:   &time,
	}

	b.BookRepository.Update(&bookModel)
	return nil
}
