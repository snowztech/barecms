package storage

import (
	"time"

	"gorm.io/datatypes"
)

type SiteDB struct {
	ID     string `gorm:"primaryKey"`
	Name   string `gorm:"uniqueIndex;not null"`
	Slug   string `gorm:"uniqueIndex;not null"`
	UserID string `gorm:"not null"`
}

func (SiteDB) TableName() string {
	return "sites"
}

type CollectionDB struct {
	ID      string         `gorm:"primaryKey"`
	Name    string         `gorm:"not null"`
	Slug    string         `gorm:"not null;uniqueIndex:idx_collections_site_slug,priority:2"`
	SiteID  string         `gorm:"not null;index;uniqueIndex:idx_collections_site_slug,priority:1"`
	Fields  datatypes.JSON `gorm:"type:jsonb"`
	Entries []EntryDB      `gorm:"foreignKey:CollectionID"`
}

func (CollectionDB) TableName() string {
	return "collections"
}

type EntryDB struct {
	ID           string         `gorm:"primaryKey"`
	CollectionID string         `gorm:"not null;index"`
	Data         datatypes.JSON `gorm:"type:jsonb"`
}

func (EntryDB) TableName() string {
	return "entries"
}

type UserDB struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null"`
	Username string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

type MediaFileDB struct {
	ID           string    `gorm:"primaryKey"`
	SiteID       string    `gorm:"not null;index"`
	StoredName   string    `gorm:"not null;uniqueIndex"`
	OriginalName string    `gorm:"not null"`
	MIMEType     string    `gorm:"not null"`
	Size         int64     `gorm:"not null"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

func (MediaFileDB) TableName() string { return "media_files" }

func (UserDB) TableName() string {
	return "users"
}

type SchemaMigrationDB struct {
	Version   string    `gorm:"primaryKey"`
	AppliedAt time.Time `gorm:"autoCreateTime"`
}

func (SchemaMigrationDB) TableName() string { return "schema_migrations" }
