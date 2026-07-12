package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func (s *Service) GetUser(userID string) (models.User, error) {
	userDB, err := s.Storage.GetUserByID(userID)
	if err != nil {
		return models.User{}, err
	}

	user := mapToUser(userDB)

	return user, nil
}

func (s *Service) DeleteUser(userID string) error {
	media, err := s.Storage.DeleteUserCascade(userID)
	if err != nil {
		return errors.Wrap(err, "failed to delete user")
	}
	for _, file := range media {
		if err := os.Remove(filepath.Join(s.Config.UploadsDir, file.StoredName)); err != nil && !os.IsNotExist(err) {
			slog.Warn("Could not remove deleted user's media file", "file", file.StoredName, "error", err)
		}
	}
	return nil
}

func mapToUser(userDB storage.UserDB) models.User {
	return models.User{
		ID:       userDB.ID,
		Email:    userDB.Email,
		Username: userDB.Username,
	}
}
