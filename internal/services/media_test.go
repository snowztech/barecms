package services

import (
	"barecms/configs"
	"barecms/internal/storage"
	"bytes"
	"errors"
	"os"
	"testing"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func mediaTestService(t *testing.T) *Service {
	t.Helper()
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	if err := db.AutoMigrate(&storage.SiteDB{}, &storage.MediaFileDB{}); err != nil {
		t.Fatal(err)
	}
	if err := db.Create(&storage.SiteDB{ID: "site-1", Name: "Site", Slug: "site", UserID: "owner-1"}).Error; err != nil {
		t.Fatal(err)
	}
	return &Service{
		Storage: &storage.Storage{DB: db},
		Config:  configs.AppConfig{UploadsDir: t.TempDir(), MaxFileSize: 1024},
	}
}

func pngBytes() []byte {
	return append([]byte{0x89, 0x50, 0x4e, 0x47, 0x0d, 0x0a, 0x1a, 0x0a}, make([]byte, 504)...)
}

func TestMediaLifecycle(t *testing.T) {
	service := mediaTestService(t)
	file, err := service.UploadMedia("site-1", "owner-1", "../../logo.png", bytes.NewReader(pngBytes()))
	if err != nil {
		t.Fatalf("upload media: %v", err)
	}
	if file.OriginalName != "logo.png" || file.MIMEType != "image/png" {
		t.Fatalf("unexpected metadata: %+v", file)
	}

	files, err := service.ListMedia("site-1", "owner-1")
	if err != nil || len(files) != 1 {
		t.Fatalf("list media: files=%d err=%v", len(files), err)
	}
	_, path, err := service.GetMedia(file.ID)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("stored file missing: %v", err)
	}

	if err := service.DeleteMedia(file.ID, "owner-1"); err != nil {
		t.Fatalf("delete media: %v", err)
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatalf("expected stored file deletion, got %v", err)
	}
}

func TestMediaRejectsCrossTenantAndUnsafeUploads(t *testing.T) {
	service := mediaTestService(t)
	if _, err := service.UploadMedia("site-1", "another-user", "logo.png", bytes.NewReader(pngBytes())); !errors.Is(err, ErrForbidden) {
		t.Fatalf("expected forbidden, got %v", err)
	}
	if _, err := service.UploadMedia("site-1", "owner-1", "script.html", bytes.NewReader([]byte("<script>alert(1)</script>"))); !errors.Is(err, ErrUnsupportedFile) {
		t.Fatalf("expected unsupported file error, got %v", err)
	}

	service.Config.MaxFileSize = 4
	if _, err := service.UploadMedia("site-1", "owner-1", "large.png", bytes.NewReader(pngBytes())); !errors.Is(err, ErrFileTooLarge) {
		t.Fatalf("expected file-too-large error, got %v", err)
	}
}
