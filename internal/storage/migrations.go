package storage

import (
	"fmt"
	"log/slog"

	"gorm.io/gorm"
)

type migration struct {
	version string
	up      func(*gorm.DB) error
}

var migrations = []migration{
	{version: "001_baseline", up: func(db *gorm.DB) error {
		return db.AutoMigrate(&SiteDB{}, &CollectionDB{}, &EntryDB{}, &UserDB{}, &MediaFileDB{})
	}},
	{version: "002_scoped_collection_indexes", up: migrateCollectionIndexes},
}

func runMigrations(db *gorm.DB) error {
	if err := db.AutoMigrate(&SchemaMigrationDB{}); err != nil {
		return fmt.Errorf("create migration ledger: %w", err)
	}
	for _, item := range migrations {
		var count int64
		if err := db.Model(&SchemaMigrationDB{}).Where("version = ?", item.version).Count(&count).Error; err != nil {
			return err
		}
		if count > 0 {
			continue
		}
		if err := db.Transaction(func(tx *gorm.DB) error {
			if err := item.up(tx); err != nil {
				return err
			}
			return tx.Create(&SchemaMigrationDB{Version: item.version}).Error
		}); err != nil {
			return fmt.Errorf("apply migration %s: %w", item.version, err)
		}
		slog.Info("Database migration applied", "version", item.version)
	}
	return nil
}

func migrateCollectionIndexes(db *gorm.DB) error {
	migrator := db.Migrator()
	if migrator.HasIndex(&CollectionDB{}, "idx_collections_slug") {
		if err := migrator.DropIndex(&CollectionDB{}, "idx_collections_slug"); err != nil {
			return err
		}
	}
	if !migrator.HasIndex(&CollectionDB{}, "idx_collections_site_slug") {
		if err := migrator.CreateIndex(&CollectionDB{}, "idx_collections_site_slug"); err != nil {
			return err
		}
	}
	return nil
}
