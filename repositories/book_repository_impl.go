package repositories

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/washington-shoji/gin-api/models"
)

type BookRepositoryImp struct {
	Database *sql.DB
}

func NewBookRepositoryImp(Db *sql.DB) BookRepository {
	return &BookRepositoryImp{
		Database: Db,
	}
}

// Create implements BookRepository.
func (b *BookRepositoryImp) Create(book *models.Book) (error error) {
	query := `
	INSERT INTO book (id, title, description, created_at)
	VALUES ($1, $2, $3, $4)
	`

	_, err := b.Database.Query(
		query,
		book.ID,
		book.Title,
		book.Description,
		book.CreatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements BookRepository.
func (b *BookRepositoryImp) FindAll() ([]*models.Book, error) {
	rows, err := b.Database.Query(`SELECT * FROM book WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}

	bks := []*models.Book{}
	for rows.Next() {
		bk, err := scanIntoBook(rows)
		if err != nil {
			return nil, err
		}
		bks = append(bks, bk)
	}

	return bks, nil
}

// Delete implements BookRepository.
func (b *BookRepositoryImp) Delete(book *models.Book) (error error) {
	_, err := b.Database.Query(`UPDATE book SET deleted_at = $1 WHERE id = $2`, book.DeletedAt, book.ID)
	if err != nil {
		return err
	}
	return nil
}

// FindByID implements BookRepository.
func (b *BookRepositoryImp) FindByID(id uuid.UUID) (book *models.Book, error error) {
	rows, err := b.Database.Query(`SELECT * FROM book WHERE id = $1 AND deleted_at IS NULL`, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoBook(rows)
	}

	return nil, fmt.Errorf("book %d not found", id)
}

// Update implements BookRepository.
func (b *BookRepositoryImp) Update(book *models.Book) (error error) {
	query := `UPDATE book
	SET 
	title = $2,
	description = $3,
	updated_at = $4
	WHERE id = $1
	`

	_, err := b.Database.Query(
		query,
		book.ID,
		book.Title,
		book.Description,
		book.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

func scanIntoBook(rows *sql.Rows) (*models.Book, error) {
	book := &models.Book{}
	err := rows.Scan(
		&book.ID,
		&book.Title,
		&book.Description,
		&book.CreatedAt,
		&book.UpdatedAt,
		&book.DeletedAt,
	)

	return book, err
}
