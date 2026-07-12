package storage

import (
	"testing"

	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type legacyCollection struct {
	ID     string         `gorm:"primaryKey"`
	SiteID string         `gorm:"not null"`
	Slug   string         `gorm:"uniqueIndex;not null"`
	Name   string         `gorm:"not null"`
	Fields datatypes.JSON `gorm:"type:jsonb"`
}

func (legacyCollection) TableName() string { return "collections" }

func migrationTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func TestMigrationsAreIdempotent(t *testing.T) {
	db := migrationTestDB(t)
	if err := runMigrations(db); err != nil {
		t.Fatal(err)
	}
	if err := runMigrations(db); err != nil {
		t.Fatalf("second migration run failed: %v", err)
	}
	var count int64
	if err := db.Model(&SchemaMigrationDB{}).Count(&count).Error; err != nil {
		t.Fatal(err)
	}
	if count != int64(len(migrations)) {
		t.Fatalf("expected %d migrations, got %d", len(migrations), count)
	}
}

func TestCollectionSlugIsUniqueWithinSite(t *testing.T) {
	db := migrationTestDB(t)
	if err := runMigrations(db); err != nil {
		t.Fatal(err)
	}
	for _, site := range []SiteDB{{ID: "one", Name: "One", Slug: "one", UserID: "user"}, {ID: "two", Name: "Two", Slug: "two", UserID: "user"}} {
		if err := db.Create(&site).Error; err != nil {
			t.Fatal(err)
		}
	}
	if err := db.Create(&CollectionDB{ID: "one-posts", SiteID: "one", Name: "Posts", Slug: "posts"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&CollectionDB{ID: "two-posts", SiteID: "two", Name: "Posts", Slug: "posts"}).Error; err != nil {
		t.Fatalf("same slug in another site should be allowed: %v", err)
	}
	if err := db.Create(&CollectionDB{ID: "duplicate", SiteID: "one", Name: "Other", Slug: "posts"}).Error; err == nil {
		t.Fatal("duplicate slug in one site should be rejected")
	}
}

func TestMigrationsUpgradeLegacyGlobalCollectionIndex(t *testing.T) {
	db := migrationTestDB(t)
	if err := db.AutoMigrate(&SiteDB{}, &legacyCollection{}); err != nil {
		t.Fatal(err)
	}
	if !db.Migrator().HasIndex(&legacyCollection{}, "idx_collections_slug") {
		t.Fatal("legacy index was not created")
	}
	if err := runMigrations(db); err != nil {
		t.Fatal(err)
	}
	if db.Migrator().HasIndex(&CollectionDB{}, "idx_collections_slug") {
		t.Fatal("legacy global index still exists")
	}
	if !db.Migrator().HasIndex(&CollectionDB{}, "idx_collections_site_slug") {
		t.Fatal("scoped index was not created")
	}
}
