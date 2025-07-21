package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectTestDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		return nil, err
	}

	return db, nil
}
