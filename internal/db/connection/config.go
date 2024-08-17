package connection

import (
	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v3/log"
	"github.com/joho/godotenv"
)

type DatabaseConfig struct {
	Host     string
	Username string
	Password string
	DBName   string
	Port     string
	AppName  string
	SSLMode  string
	Timezone string
}

func NewEnvConfig() *DatabaseConfig {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load .env: %e", err)
	}

	config := DatabaseConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from .env: %e", err)
	}

	return &config
}
