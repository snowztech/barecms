package storage

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func cascadeTestStorage(t *testing.T) *Storage {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&SiteDB{}, &CollectionDB{}, &EntryDB{}, &MediaFileDB{}); err != nil {
		t.Fatal(err)
	}
	return &Storage{DB: db}
}

func TestDeleteSiteCascadeRemovesAllDatabaseResources(t *testing.T) {
	storage := cascadeTestStorage(t)
	for _, record := range []any{
		&SiteDB{ID: "site", Name: "Site", Slug: "site", UserID: "user"},
		&CollectionDB{ID: "collection", Name: "Posts", Slug: "posts", SiteID: "site"},
		&EntryDB{ID: "entry", CollectionID: "collection", Data: []byte(`{}`)},
		&MediaFileDB{ID: "media", SiteID: "site", StoredName: "media.png", OriginalName: "image.png", MIMEType: "image/png", Size: 1},
	} {
		if err := storage.DB.Create(record).Error; err != nil {
			t.Fatal(err)
		}
	}
	media, err := storage.DeleteSiteCascade("site")
	if err != nil {
		t.Fatal(err)
	}
	if len(media) != 1 || media[0].StoredName != "media.png" {
		t.Fatalf("unexpected media: %+v", media)
	}
	for name, model := range map[string]any{"sites": &SiteDB{}, "collections": &CollectionDB{}, "entries": &EntryDB{}, "media": &MediaFileDB{}} {
		var count int64
		if err := storage.DB.Model(model).Count(&count).Error; err != nil {
			t.Fatal(err)
		}
		if count != 0 {
			t.Errorf("expected no %s, got %d", name, count)
		}
	}
}

func TestDeleteUserCascadeRemovesOwnedGraph(t *testing.T) {
	storage := cascadeTestStorage(t)
	if err := storage.DB.AutoMigrate(&UserDB{}); err != nil {
		t.Fatal(err)
	}
	for _, record := range []any{
		&UserDB{ID: "user", Email: "user@example.com", Username: "user", Password: "hash"},
		&SiteDB{ID: "site", Name: "Site", Slug: "site", UserID: "user"},
		&CollectionDB{ID: "collection", Name: "Posts", Slug: "posts", SiteID: "site"},
		&EntryDB{ID: "entry", CollectionID: "collection", Data: []byte(`{}`)},
		&MediaFileDB{ID: "media", SiteID: "site", StoredName: "media.png", OriginalName: "image.png", MIMEType: "image/png", Size: 1},
	} {
		if err := storage.DB.Create(record).Error; err != nil {
			t.Fatal(err)
		}
	}
	media, err := storage.DeleteUserCascade("user")
	if err != nil {
		t.Fatal(err)
	}
	if len(media) != 1 {
		t.Fatalf("unexpected media: %+v", media)
	}
	for name, model := range map[string]any{"users": &UserDB{}, "sites": &SiteDB{}, "collections": &CollectionDB{}, "entries": &EntryDB{}, "media": &MediaFileDB{}} {
		var count int64
		if err := storage.DB.Model(model).Count(&count).Error; err != nil {
			t.Fatal(err)
		}
		if count != 0 {
			t.Errorf("expected no %s, got %d", name, count)
		}
	}
}

func TestDeleteSiteCascadeRollsBackOnFailure(t *testing.T) {
	storage := cascadeTestStorage(t)
	for _, record := range []any{
		&SiteDB{ID: "site", Name: "Site", Slug: "site", UserID: "user"},
		&CollectionDB{ID: "collection", Name: "Posts", Slug: "posts", SiteID: "site"},
		&EntryDB{ID: "entry", CollectionID: "collection", Data: []byte(`{}`)},
	} {
		if err := storage.DB.Create(record).Error; err != nil {
			t.Fatal(err)
		}
	}
	if err := storage.DB.Exec(`CREATE TRIGGER prevent_collection_delete BEFORE DELETE ON collections BEGIN SELECT RAISE(FAIL, 'blocked'); END`).Error; err != nil {
		t.Fatal(err)
	}
	if _, err := storage.DeleteSiteCascade("site"); err == nil {
		t.Fatal("expected cascade failure")
	}
	for name, model := range map[string]any{"sites": &SiteDB{}, "collections": &CollectionDB{}, "entries": &EntryDB{}} {
		var count int64
		if err := storage.DB.Model(model).Count(&count).Error; err != nil {
			t.Fatal(err)
		}
		if count != 1 {
			t.Errorf("expected rollback to preserve %s, got %d", name, count)
		}
	}
}
