package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func createSchema(db *sql.DB) error {
	schema := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT UNIQUE NOT NULL,
        	name VARCHAR UNIQUE NOT NULL,
        	password_hash TEXT NOT NULL,
        	created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	_, err := db.Exec(schema)
	if err != nil {
		return err
	}

	return nil
}

func StartTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	if err := createSchema(db); err != nil {
		return nil, err
	}

	return db, nil
}
