package handlers

import (
	"barecms/internal/services"
	"barecms/internal/storage"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/datatypes"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func publicContentHandler(t *testing.T) *Handler {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&storage.SiteDB{}, &storage.CollectionDB{}, &storage.EntryDB{}); err != nil {
		t.Fatal(err)
	}
	for _, record := range []any{
		&storage.SiteDB{ID: "site", Name: "Site", Slug: "site", UserID: "user"},
		&storage.CollectionDB{ID: "posts", SiteID: "site", Name: "Posts", Slug: "posts"},
		&storage.EntryDB{ID: "entry", CollectionID: "posts", Data: datatypes.JSON(`{"title":{"value":"Hello","type":"string"}}`)},
	} {
		if err := db.Create(record).Error; err != nil {
			t.Fatal(err)
		}
	}
	return &Handler{Service: &services.Service{Storage: &storage.Storage{DB: db}}}
}

func TestPublicEntriesSetsCachePolicy(t *testing.T) {
	e := echo.New()
	handler := publicContentHandler(t)
	e.GET("/content/:siteSlug/:collectionSlug", handler.GetPublicEntries)
	response := httptest.NewRecorder()
	e.ServeHTTP(response, httptest.NewRequest(http.MethodGet, "/content/site/posts", nil))
	if response.Code != http.StatusOK {
		t.Fatalf("expected 200, got %d", response.Code)
	}
	if response.Header().Get(echo.HeaderCacheControl) != publicContentCacheControl {
		t.Fatalf("unexpected cache policy: %q", response.Header().Get(echo.HeaderCacheControl))
	}
}
