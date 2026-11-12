package db

import (
	"fmt"

	"github.com/yourorg/nebula-bridge/internal/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) (*gorm.DB, error) {
	if cfg.DatabaseURL != "" {
		dsn := cfg.DatabaseURL
		return gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	// Fallback to a local sqlite file for development/tests
	dbPath := "nebula.db"
	conn := sqlite.Open(dbPath)
	return gorm.Open(conn, &gorm.Config{})
}
