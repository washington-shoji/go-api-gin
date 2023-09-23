package databases

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func CreateTables(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS account (
		id SERIAL PRIMARY KEY NOT NULL,
		user_name VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		first_name VARCHAR(50),
		last_name VARCHAR(50),
		encrypted_password VARCHAR(256) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS user_account (
		id uuid PRIMARY KEY NOT NULL,
		user_name VARCHAR(50) UNIQUE NOT NULL,
		email VARCHAR(50) UNIQUE NOT NULL,
		full_name VARCHAR(50) NOT NULL,
		encrypted_password VARCHAR(256) NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	return err
}
