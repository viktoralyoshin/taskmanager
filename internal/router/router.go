package router

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func NewRouter(logger *logrus.Logger) *mux.Router {
	r := mux.NewRouter()

	return r
}
