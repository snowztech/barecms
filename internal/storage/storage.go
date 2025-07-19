package storage

import (
	"log/slog"

	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func NewStorage(databaseURL string) (*Storage, error) {
	database, err := openDatabase(databaseURL)
	if err != nil {
		slog.Error("Failed to open database", "error", err)
		return nil, err
	}

	slog.Info("Database connection established.")

	return &Storage{DB: database}, nil
}

func openDatabase(url string) (*gorm.DB, error) {
	database, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to database")
	}

	// Migrate the schema
	if err := database.AutoMigrate(&SiteDB{}, &CollectionDB{}, &EntryDB{}, &UserDB{}); err != nil {
		return nil, errors.Wrap(err, "failed to migrate database schema")
	}

	slog.Info("Database migrated successfully.")

	return database, nil
}
