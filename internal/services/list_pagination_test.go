package services

import (
	"barecms/internal/storage"
	"fmt"
	"testing"
)

func TestSiteAndCollectionPagesAreBounded(t *testing.T) {
	service := entryTestService(t)
	for index := 2; index <= 4; index++ {
		site := storage.SiteDB{ID: fmt.Sprintf("site-%d", index), Name: fmt.Sprintf("Site %d", index), Slug: fmt.Sprintf("site-%d", index), UserID: "owner-1"}
		if err := service.Storage.CreateSite(site); err != nil {
			t.Fatal(err)
		}
	}
	sites, err := service.GetSitesPage("owner-1", 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(sites.Sites) != 2 || sites.Pagination.Total != 4 || sites.Pagination.TotalPages != 2 {
		t.Fatalf("unexpected sites page: %+v", sites)
	}

	for index := 2; index <= 4; index++ {
		collection := storage.CollectionDB{ID: fmt.Sprintf("collection-%d", index), SiteID: "site-1", Name: fmt.Sprintf("Collection %d", index), Slug: fmt.Sprintf("collection-%d", index)}
		if err := service.Storage.CreateCollection(collection); err != nil {
			t.Fatal(err)
		}
	}
	collections, err := service.GetCollectionsPage("site-1", "owner-1", 1, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(collections.Collections) != 2 || collections.Pagination.Total != 4 || collections.Pagination.TotalPages != 2 {
		t.Fatalf("unexpected collections page: %+v", collections)
	}
}

func TestMediaPageIsBounded(t *testing.T) {
	service := mediaTestService(t)
	for index := 1; index <= 3; index++ {
		file := storage.MediaFileDB{ID: fmt.Sprintf("media-%d", index), SiteID: "site-1", StoredName: fmt.Sprintf("media-%d.png", index), OriginalName: "image.png", MIMEType: "image/png", Size: 1}
		if err := service.Storage.CreateMediaFile(&file); err != nil {
			t.Fatal(err)
		}
	}
	page, err := service.ListMediaPage("site-1", "owner-1", 2, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Files) != 1 || page.Pagination.Total != 3 || page.Pagination.TotalPages != 2 {
		t.Fatalf("unexpected media page: %+v", page)
	}
}
