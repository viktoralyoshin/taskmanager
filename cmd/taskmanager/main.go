package main

import (
	"log"

	"github.com/viktoralyoshin/taskmanager/internal/app"
	"github.com/viktoralyoshin/taskmanager/internal/config"
)

func main() {
	config, err := config.NewConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	app.Start(config)
}
