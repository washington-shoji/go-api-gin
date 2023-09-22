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

	if _, err := b.BookRepository.FindByID(id); err != nil {
		return err
	}

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
func (b *BookServiceImpl) FindAll() (books []*models.BookResponse, error error) {
	result, err := b.BookRepository.FindAll()
	if err != nil {
		return nil, err
	}

	resp := books
	// ignore the first argument (index)
	// iterate over the results and append
	for _, rst_item := range result {
		// append the BookResponse to resp slice (func response)
		resp = append(resp, &models.BookResponse{
			ID:          rst_item.ID,
			Title:       rst_item.Title,
			Description: rst_item.Description,
		})
	}
	return resp, nil

}

// FindByID implements BookService.
func (b *BookServiceImpl) FindByID(id uuid.UUID, bk *models.FindByIDBookRequest) (book *models.BookResponse, error error) {

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

	resp := &models.BookResponse{
		ID:          result.ID,
		Title:       result.Title,
		Description: result.Description,
	}

	return resp, nil
}

// Update implements BookService.
func (b *BookServiceImpl) Update(id uuid.UUID, book *models.UpdateBookRequest) (error error) {

	err := b.Validate.Struct(book)
	if err != nil {
		return err
	}

	if _, err := b.BookRepository.FindByID(id); err != nil {
		return err
	}

	time := time.Now()
	bookModel := models.Book{
		ID:          book.ID,
		Title:       book.Title,
		Description: book.Description,
		UpdatedAt:   &time,
	}

	if err := b.BookRepository.Update(&bookModel); err != nil {
		return err
	}
	return nil
}
