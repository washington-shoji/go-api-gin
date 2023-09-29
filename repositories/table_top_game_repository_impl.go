package repositories

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/washington-shoji/gin-api/models"
)

type TableTopGameRepositoryImpl struct {
	Database *sql.DB
}

func NewTableTopGameRepositoryImpl(Db *sql.DB) TableTopGameRepository {
	return &TableTopGameRepositoryImpl{
		Database: Db,
	}
}

// Create implements TableTopGameRepository.
func (tbg *TableTopGameRepositoryImpl) Create(game *models.TableTopGame) error {
	query := `INSERT INTO table_top_game (id, name, game_detail, created_at)
	VALUES ($1, $2, $3, $4)
	`

	jsonbData, err := json.Marshal(game.GameDetail)
	if err != nil {
		return err
	}

	time := time.Now()
	if _, err := tbg.Database.Query(query, game.ID, game.Name, jsonbData, time); err != nil {
		return err
	}

	return nil
}

// Update implements TableTopGameRepository.
func (tbg *TableTopGameRepositoryImpl) Update(id uuid.UUID, game *models.TableTopGame) error {
	query := `UPDATE table_top_game
			SET name = $1, game_detail = $1, updated_at = now()
			WHERE id = $3
	`

	_, err := tbg.Database.Query(query, game.Name, game.GameDetail, id)
	if err != nil {
		return err
	}

	return nil
}

// Delete implements TableTopGameRepository.
func (tbg *TableTopGameRepositoryImpl) Delete(id uuid.UUID) error {
	query := `UPDATE table_top_game
		SET deleted_at = now()
		WHERE id = $1
	`

	_, err := tbg.Database.Query(query, id)
	if err != nil {
		return err
	}

	return nil
}

// FindAll implements TableTopGameRepository.
func (tbg *TableTopGameRepositoryImpl) FindAll() ([]*models.TableTopGame, error) {
	query := `SELECT * from table_top_game WHERE deleted_at IS NULL`

	rows, err := tbg.Database.Query(query)
	if err != nil {
		fmt.Println("err 1", err)
		return nil, err
	}

	for rows.Next() {
		var jsonData []byte

		if err := rows.Scan(&jsonData); err != nil {
			fmt.Println("Error scanning JSONB data:", err)
			return nil, err
		}

		var games models.TableTopGameResp

		if err := json.Unmarshal(jsonData, &games); err != nil {
			fmt.Println("Error unmarshalling JSONB data:", err)
			return nil, err
		}

		fmt.Println("games", games)
	}

	// games := []*models.TableTopGame{}

	// for rows.Next() {
	// 	gm, err := scanIntoGame(rows)
	// 	if err != nil {
	// 		fmt.Println("err 2", err)
	// 		return nil, err
	// 	}

	// 	games = append(games, gm)
	// }

	// return games, nil

	return nil, nil

}

// FindByID implements TableTopGameRepository.
func (tbg *TableTopGameRepositoryImpl) FindByID(id uuid.UUID) (*models.TableTopGame, error) {
	query := `SELECT * from table_top_game WHERE id = $1 AND deleted_at IS NULL`

	rows, err := tbg.Database.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoGame(rows)
	}

	return nil, fmt.Errorf("game %d not found", id)
}

func scanIntoGame(rows *sql.Rows) (*models.TableTopGame, error) {
	game := &models.TableTopGame{}
	err := rows.Scan(
		&game.ID,
		&game.Name,
		&game.GameDetail,
		&game.CreatedAt,
		&game.UpdatedAt,
		&game.DeletedAt,
	)

	return game, err
}
