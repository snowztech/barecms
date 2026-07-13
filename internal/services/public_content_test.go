package services

import "testing"

func TestPublicContentIsPaginatedAndFlattened(t *testing.T) {
	service := entryTestService(t)
	page, err := service.GetPublicEntries("site", "posts", 1, 20)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Entries) != 1 || page.Pagination.Total != 1 || page.Pagination.TotalPages != 1 {
		t.Fatalf("unexpected page: %+v", page)
	}
	if page.Entries[0]["id"] != "entry-1" || page.Entries[0]["title"] != "Old" {
		t.Fatalf("unexpected public entry: %+v", page.Entries[0])
	}
}

func TestPublicEntryIsScopedToSiteAndCollection(t *testing.T) {
	service := entryTestService(t)
	entry, err := service.GetPublicEntry("site", "posts", "entry-1")
	if err != nil {
		t.Fatal(err)
	}
	if entry["id"] != "entry-1" {
		t.Fatalf("unexpected entry: %+v", entry)
	}
	if _, err := service.GetPublicEntry("site", "missing", "entry-1"); err == nil {
		t.Fatal("expected missing collection to be rejected")
	}
}
