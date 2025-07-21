package app

import (
	"net/http"

	"github.com/viktoralyoshin/taskmanager/internal/config"
	"github.com/viktoralyoshin/taskmanager/internal/db"
)

func Start(config *config.Config) error {
	s := NewServer(config)

	db, err := db.ConnectTestDB()
	if err != nil {
		s.logger.Errorf("Database connection error: %v", err)
	}

	s.logger.Infof("Starting server on port: %v", config.BindAddr)

	if err := http.ListenAndServe(config.BindAddr, s.router); err != nil {
		s.logger.Fatalf("error starting server: %v", err)
	}

	return nil
}
