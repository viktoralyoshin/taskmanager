package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	BindAddr   string
	LogLevel   string
}

func NewConfig() (*Config, error) {

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	bindAddr := os.Getenv("BIND_ADDR")
	logLevel := os.Getenv("LOG_LEVEL")

	return &Config{
		DBUser:     dbUser,
		DBPassword: dbPassword,
		DBHost:     dbHost,
		DBPort:     dbPort,
		DBName:     dbName,
		BindAddr:   bindAddr,
		LogLevel:   logLevel,
	}, nil
}
