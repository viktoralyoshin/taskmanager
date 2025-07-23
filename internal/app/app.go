package app

import (
	"net/http"

	"github.com/viktoralyoshin/taskmanager/internal/config"
	"github.com/viktoralyoshin/taskmanager/internal/db"
)

func Start(config *config.Config) error {
	s := NewServer(config)

	db, err := db.ConnectDB(config)
	if err != nil {
		s.logger.Errorf("Database connection error: %v", err)
	}
	defer db.Close()

	s.logger.Infof("Starting server on port: %v", config.BindAddr)

	if err := http.ListenAndServe(config.BindAddr, s.router); err != nil {
		s.logger.Fatalf("Error starting server: %v", err)
	}

	return nil
}
