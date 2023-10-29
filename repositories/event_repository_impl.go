package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type EventRepositoryImp struct {
	Database *sql.DB
}

func NewEventRepositoryImp(Db *sql.DB) EventRepository {
	return &EventRepositoryImp{
		Database: Db,
	}
}

func (repo *EventRepositoryImp) Create(event *models.Event) error {
	query := `INSERT event_table (
		id, 
		title, 
		short_description, 
		description, 
		image_url, 
		date, 
		registration, 
		created_at)
		VALUES($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := repo.Database.Query(
		query,
		event.ID,
		event.Title,
		event.ShortDescription,
		event.Description,
		event.ImageUrl,
		event.Date,
		event.Registration,
		event.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *EventRepositoryImp) Update(event *models.Event) error {
	query := `UPDATE event_table 
	SET
	title = $2, 
	short_description = $3, 
	description $4, 
	image_url = $5, 
	date = $6, 
	registration $7,
	updated_at = $8
	WHERE id = $1  
	`

	_, err := repo.Database.Query(
		query,
		event.ID,
		event.Title,
		event.ShortDescription,
		event.Description,
		event.ImageUrl,
		event.Date,
		event.Registration,
		event.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}

func (repo *EventRepositoryImp) Delete(id *uuid.UUID) error {
	panic("unimplemented")
}

func (repo *EventRepositoryImp) FindAll() ([]*models.Event, error) {
	panic("unimplemented")
}

func (repo *EventRepositoryImp) FindByID(id *uuid.UUID) (*models.Event, error) {
	panic("unimplemented")
}
