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
	CREATE TABLE IF NOT EXISTS table_top_game (
		id uuid PRIMARY KEY NOT NULL,
		name VARCHAR(100) UNIQUE NOT NULL,
		game_detail JSONB,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS book (
		id uuid PRIMARY KEY NOT NULL,
		title VARCHAR(50) UNIQUE NOT NULL,
		description VARCHAR(500),
		image_url VARCHAR(500),
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS dynamic_data (
		id uuid PRIMARY KEY NOT NULL,
		data JSONB,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS table_exp (
		id uuid PRIMARY KEY NOT NULL,
		exp_json JSONB,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP,
		deleted_at TIMESTAMP
	);
	CREATE TABLE IF NOT EXISTS event_table (
		id uuid PRIMARY KEY NOT NULL,
		title VARCHAR(100) UNIQUE NOT NULL,
		short_description VARCHAR(200) NOT NULL,
		description VARCHAR(1000) NOT NULL,
		image_url VARCHAR(500) NOT NULL,
		image_public_id VARCHAR(500) NOT NULL,
		date TIMESTAMP NOT NULL,
		registration TIMESTAMP NOT NULL,
		created_at TIMESTAMP NOT NULL,
		updated_at TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	return err
}
