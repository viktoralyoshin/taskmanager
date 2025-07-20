package app

import (
	"net/http"
)

func Start(config *Config) error {
	s := NewServer(config)

	s.logger.Infof("Starting server on port: %v", config.BindAddr)

	if err := http.ListenAndServe(config.BindAddr, s.router); err != nil {
		s.logger.Fatalf("error starting server: %v", err)
	}

	return nil
}
