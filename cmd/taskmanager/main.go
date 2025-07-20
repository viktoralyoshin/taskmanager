package main

import (
	"log"

	"github.com/viktoralyoshin/taskmanager/internal/app"
)

func main() {
	config, err := app.NewConfig()
	if err != nil {
		log.Fatalf("configuration error: %v", err)
	}

	app.Start(config)
}
