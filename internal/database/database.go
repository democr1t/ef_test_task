package database

import (
	"effective_mobile_test_task/internal/models"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log/slog"
)

func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	slog.Debug("database connect success")

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := db.AutoMigrate(&models.Person{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}
	slog.Debug("database migrate success")

	return db, nil
}
