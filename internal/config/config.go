package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DatabaseUrl string
}

func LoadConfig() *Config {
	godotenv.Load()
	dbUrl := os.Getenv("DATABASE_URL")
	if dbUrl == "" {
		log.Fatal("DATABASE_URL is required")
	}

	return &Config{
		DatabaseUrl: dbUrl,
	}
}
