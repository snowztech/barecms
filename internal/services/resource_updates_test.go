package services

import (
	"barecms/internal/models"
	"errors"
	"testing"
)

func TestUpdateSiteKeepsSlugAndEnforcesOwnership(t *testing.T) {
	service := entryTestService(t)
	if _, err := service.UpdateSite("site-1", "another-user", models.UpdateSiteRequest{Name: "Renamed"}); !errors.Is(err, ErrForbidden) {
		t.Fatalf("expected forbidden, got %v", err)
	}
	updated, err := service.UpdateSite("site-1", "owner-1", models.UpdateSiteRequest{Name: "Renamed"})
	if err != nil {
		t.Fatal(err)
	}
	if updated.Name != "Renamed" || updated.Slug != "site" {
		t.Fatalf("unexpected site: %+v", updated)
	}
}

func TestUpdateCollectionProtectsExistingEntries(t *testing.T) {
	service := entryTestService(t)
	compatible := models.UpdateCollectionRequest{Name: "Articles", Fields: []models.Field{
		{Name: "title", Type: models.FieldTypeString},
		{Name: "image", Type: models.FieldTypeImage, Optional: true},
	}}
	updated, err := service.UpdateCollection("collection-1", "owner-1", compatible)
	if err != nil {
		t.Fatal(err)
	}
	if updated.Name != "Articles" || updated.Slug != "posts" || len(updated.Fields) != 2 {
		t.Fatalf("unexpected collection: %+v", updated)
	}

	incompatible := models.UpdateCollectionRequest{Name: "Articles", Fields: []models.Field{
		{Name: "title", Type: models.FieldTypeString},
		{Name: "author", Type: models.FieldTypeString},
	}}
	if _, err := service.UpdateCollection("collection-1", "owner-1", incompatible); err == nil {
		t.Fatal("expected incompatible schema to be rejected")
	}
}
