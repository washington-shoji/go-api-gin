package databases

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/washington-shoji/gin-api/config"
)

func DatabaseConnection() (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.EnvConfig("DB_USER"), config.EnvConfig("DB_PASSWORD"), config.EnvConfig("DB_NAME"))
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
