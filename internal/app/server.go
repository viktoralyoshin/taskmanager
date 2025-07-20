package app

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/viktoralyoshin/taskmanager/internal/router"
)

type Server struct {
	router *mux.Router
	logger *logrus.Logger
}

func NewServer(config *Config) *Server {
	logger := logrus.New()
	return &Server{
		logger: logger,
		router: router.NewRouter(logger),
	}
}

func NewLogger(config *Config) *logrus.Logger {
	return nil
}
