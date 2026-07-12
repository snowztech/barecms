package storage

import (
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func authorizationTestStorage(t *testing.T) *Storage {
	t.Helper()

	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("open test database: %v", err)
	}
	if err := db.AutoMigrate(&SiteDB{}, &CollectionDB{}, &EntryDB{}); err != nil {
		t.Fatalf("migrate test database: %v", err)
	}

	return &Storage{DB: db}
}

func TestOwnershipFollowsResourceHierarchy(t *testing.T) {
	storage := authorizationTestStorage(t)

	if err := storage.DB.Create(&SiteDB{ID: "site-1", Name: "Site", Slug: "site", UserID: "owner-1"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := storage.DB.Create(&CollectionDB{ID: "collection-1", Name: "Posts", Slug: "posts", SiteID: "site-1"}).Error; err != nil {
		t.Fatal(err)
	}
	if err := storage.DB.Create(&EntryDB{ID: "entry-1", CollectionID: "collection-1", Data: []byte(`{}`)}).Error; err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name  string
		check func(string) (bool, error)
	}{
		{name: "site", check: func(userID string) (bool, error) { return storage.UserOwnsSite(userID, "site-1") }},
		{name: "collection", check: func(userID string) (bool, error) { return storage.UserOwnsCollection(userID, "collection-1") }},
		{name: "entry", check: func(userID string) (bool, error) { return storage.UserOwnsEntry(userID, "entry-1") }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			owned, err := test.check("owner-1")
			if err != nil || !owned {
				t.Fatalf("owner should have access: owned=%v err=%v", owned, err)
			}

			owned, err = test.check("another-user")
			if err != nil {
				t.Fatalf("check another user: %v", err)
			}
			if owned {
				t.Fatal("another user must not have access")
			}
		})
	}
}
