package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	_ "github.com/lib/pq"

	"github.com/google/uuid"
	"github.com/washington-shoji/gin-api/models"
)

type DynamicDataRepositoryImpl struct {
	Database *sql.DB
}

func NewDynamicDataRepositoryImpl(Db *sql.DB) DynamicDataRepository {
	return &DynamicDataRepositoryImpl{
		Database: Db,
	}
}

// Create implements DynamicDataRepository.
func (repo *DynamicDataRepositoryImpl) Create(dyn *models.DynamicData) error {

	query := `INSERT INTO dynamic_data (id, data, created_at)
	VALUES ($1, $2, $3)
	`
	id := uuid.New()
	time := time.Now()

	_, err := repo.Database.Query(query, id, dyn.Data, time)
	if err != nil {
		return err
	}

	return nil
}

// Update implements DynamicDataRepository.
func (repo *DynamicDataRepositoryImpl) Update(id uuid.UUID, dyn *models.DynamicData) error {
	query := `UPDATE dynamic_data
	SET data = $1, updated_at = $2
	WHERE id = $3
	`

	time := time.Now()
	_, err := repo.Database.Query(query, dyn.Data, time, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements DynamicDataRepository.
func (repo *DynamicDataRepositoryImpl) Delete(id uuid.UUID) error {
	query := `UPDATE dynamic_data
	SET deleted_at = $1
	WHERE id = $2
	`
	time := time.Now()
	_, err := repo.Database.Query(query, time, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements DynamicDataRepository.
func (repo *DynamicDataRepositoryImpl) FindAll() ([]*models.DynamicData, error) {
	query := `SELECT * from dynamic_data WHERE deleted_at IS NULL`

	rows, err := repo.Database.Query(query)
	if err != nil {
		return nil, err
	}

	resp := []*models.DynamicData{}

	for rows.Next() {
		data, err := scanIntoDynamicData(rows)
		if err != nil {
			return nil, err
		}

		resp = append(resp, data)
	}

	return resp, nil
}

// FindById implements DynamicDataRepository.
func (repo *DynamicDataRepositoryImpl) FindById(id uuid.UUID) (*models.DynamicData, error) {
	query := `SELECT * from dynamic_data WHERE id = $1 AND deleted_at IS NULL`

	rows, err := repo.Database.Query(query, id)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoDynamicData(rows)
	}

	return nil, fmt.Errorf("data for %d not found", id)
}

func scanIntoDynamicData(rows *sql.Rows) (*models.DynamicData, error) {
	data := &models.DynamicData{}
	var jsonColumnData []byte
	err := rows.Scan(
		&data.ID,
		&jsonColumnData,
		&data.CreatedAt,
		&data.UpdatedAt,
		&data.DeletedAt,
	)

	if err := json.Unmarshal(jsonColumnData, &data); err != nil {
		fmt.Println("Error unmarshalling JSONB data:", err)
		return nil, err
	}

	return data, err
}
