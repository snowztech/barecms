package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"encoding/json"
	"errors"
	"testing"

	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func entryTestService(t *testing.T) *Service {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&storage.SiteDB{}, &storage.CollectionDB{}, &storage.EntryDB{}); err != nil {
		t.Fatal(err)
	}
	fields, _ := json.Marshal([]models.Field{{Name: "title", Type: models.FieldTypeString}})
	for _, record := range []any{
		&storage.SiteDB{ID: "site-1", Name: "Site", Slug: "site", UserID: "owner-1"},
		&storage.CollectionDB{ID: "collection-1", SiteID: "site-1", Name: "Posts", Slug: "posts", Fields: datatypes.JSON(fields)},
		&storage.EntryDB{ID: "entry-1", CollectionID: "collection-1", Data: datatypes.JSON(`{"title":{"value":"Old","type":"string"}}`)},
	} {
		if err := db.Create(record).Error; err != nil {
			t.Fatal(err)
		}
	}
	return &Service{Storage: &storage.Storage{DB: db}}
}

func TestUpdateEntryValidatesOwnershipAndSchema(t *testing.T) {
	service := entryTestService(t)
	valid := models.UpdateEntryRequest{Data: json.RawMessage(`{"title":{"value":"New","type":"string"}}`)}
	if _, err := service.UpdateEntry("entry-1", "another-user", valid); !errors.Is(err, ErrForbidden) {
		t.Fatalf("expected forbidden, got %v", err)
	}
	invalid := models.UpdateEntryRequest{Data: json.RawMessage(`{"title":{"value":"","type":"string"}}`)}
	if _, err := service.UpdateEntry("entry-1", "owner-1", invalid); err == nil {
		t.Fatal("expected validation error")
	}
	updated, err := service.UpdateEntry("entry-1", "owner-1", valid)
	if err != nil {
		t.Fatal(err)
	}
	if string(updated.Data) != string(valid.Data) {
		t.Fatalf("unexpected data: %s", updated.Data)
	}
}
