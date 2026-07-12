package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"barecms/internal/utils"
	"errors"
	"fmt"
	"io"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

var (
	ErrFileTooLarge    = errors.New("file is too large")
	ErrUnsupportedFile = errors.New("unsupported file type")
)

var allowedMediaTypes = map[string]bool{
	"image/jpeg":      true,
	"image/png":       true,
	"image/gif":       true,
	"image/webp":      true,
	"application/pdf": true,
}

func (s *Service) UploadMedia(siteID, userID, originalName string, source io.Reader) (models.MediaFile, error) {
	if err := s.requireSiteOwner(userID, siteID); err != nil {
		return models.MediaFile{}, err
	}

	contents, err := io.ReadAll(io.LimitReader(source, s.Config.MaxFileSize+1))
	if err != nil {
		return models.MediaFile{}, err
	}
	if int64(len(contents)) > s.Config.MaxFileSize {
		return models.MediaFile{}, fmt.Errorf("%w: maximum is %d bytes", ErrFileTooLarge, s.Config.MaxFileSize)
	}
	if len(contents) == 0 {
		return models.MediaFile{}, fmt.Errorf("file cannot be empty")
	}

	mimeType := http.DetectContentType(contents)
	if !allowedMediaTypes[mimeType] {
		return models.MediaFile{}, fmt.Errorf("%w: %s", ErrUnsupportedFile, mimeType)
	}

	id := utils.GenerateUniqueID()
	extension := mediaExtension(mimeType)
	storedName := id + extension
	if err := os.MkdirAll(s.Config.UploadsDir, 0o750); err != nil {
		return models.MediaFile{}, err
	}
	path := filepath.Join(s.Config.UploadsDir, storedName)
	if err := os.WriteFile(path, contents, 0o640); err != nil {
		return models.MediaFile{}, err
	}

	fileDB := storage.MediaFileDB{ID: id, SiteID: siteID, StoredName: storedName, OriginalName: filepath.Base(originalName), MIMEType: mimeType, Size: int64(len(contents))}
	if err := s.Storage.CreateMediaFile(&fileDB); err != nil {
		_ = os.Remove(path)
		return models.MediaFile{}, err
	}
	return mapToMediaFile(fileDB), nil
}

func (s *Service) ListMedia(siteID, userID string) ([]models.MediaFile, error) {
	if err := s.requireSiteOwner(userID, siteID); err != nil {
		return nil, err
	}
	filesDB, err := s.Storage.GetMediaFilesBySiteID(siteID)
	if err != nil {
		return nil, err
	}
	files := make([]models.MediaFile, len(filesDB))
	for i, file := range filesDB {
		files[i] = mapToMediaFile(file)
	}
	return files, nil
}

func (s *Service) GetMedia(id string) (models.MediaFile, string, error) {
	file, err := s.Storage.GetMediaFile(id)
	if err != nil {
		return models.MediaFile{}, "", err
	}
	return mapToMediaFile(file), filepath.Join(s.Config.UploadsDir, file.StoredName), nil
}

func (s *Service) DeleteMedia(id, userID string) error {
	file, err := s.Storage.GetMediaFile(id)
	if err != nil {
		return err
	}
	if err := s.requireSiteOwner(userID, file.SiteID); err != nil {
		return err
	}
	if err := s.Storage.DeleteMediaFile(id); err != nil {
		return err
	}
	if err := os.Remove(filepath.Join(s.Config.UploadsDir, file.StoredName)); err != nil && !os.IsNotExist(err) {
		return err
	}
	return nil
}

func mediaExtension(mimeType string) string {
	extensions, _ := mime.ExtensionsByType(mimeType)
	if len(extensions) > 0 {
		return extensions[0]
	}
	return ""
}

func mapToMediaFile(file storage.MediaFileDB) models.MediaFile {
	return models.MediaFile{ID: file.ID, SiteID: file.SiteID, OriginalName: file.OriginalName, MIMEType: file.MIMEType, Size: file.Size, URL: "/api/files/" + file.ID, CreatedAt: file.CreatedAt}
}
