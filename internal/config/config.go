package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	DatabaseURL string
	LogLevel    string
}

func Load() *Config {
	_ = godotenv.Load()

	cfg := &Config{
		Port:        os.Getenv("PORT"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		LogLevel:    os.Getenv("LOG_LEVEL"),
	}

	if cfg.LogLevel == "" {
		cfg.LogLevel = "info"
	}

	log.Printf("config loaded: PORT=%s, DB_PRESENT=%v", cfg.Port, cfg.DatabaseURL != "")
	return cfg
}
