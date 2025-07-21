package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/viktoralyoshin/taskmanager/internal/config"
)

func ConnectDB(config *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%v port=%v user=%v password=%v dbname=%v",
		config.DBHost,
		config.DBPort,
		config.DBUser,
		config.DBPassword,
		config.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, err
}
