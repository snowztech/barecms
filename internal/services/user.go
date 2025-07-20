package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"

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
	// Delete user resources first
	if err := s.DeleteUserResources(userID); err != nil {
		return errors.Wrap(err, "failed to delete user resources")
	}

	// Finally, delete the user itself
	if err := s.Storage.DeleteUserByID(userID); err != nil {
		return errors.Wrap(err, "failed to delete user")
	}

	return nil
}

func (s *Service) DeleteUserResources(userID string) error {
	// Get sites owned by the user
	userSites, err := s.Storage.GetSitesByUserID(userID)
	if err != nil {
		return errors.Wrap(err, "failed to get user sites")
	}
	var siteIDs []string
	for _, site := range userSites {
		siteIDs = append(siteIDs, site.ID)
	}

	// Get collections from the sites
	siteCollections, err := s.Storage.GetCollectionsFromSitesIDs(siteIDs)
	if err != nil {
		return errors.Wrap(err, "failed to get collections from sites")
	}
	var collectionIDs []string
	for _, collection := range siteCollections {
		collectionIDs = append(collectionIDs, collection.ID)
	}

	// Delete entries by collection IDs
	if len(collectionIDs) > 0 {
		if err := s.Storage.DeleteEntriesByCollectionIDs(collectionIDs); err != nil {
			return errors.Wrap(err, "failed to delete entries")
		}
	}

	// Delete collections by site IDs
	if len(siteIDs) > 0 {
		if err := s.Storage.DeleteCollectionsBySiteIDs(siteIDs); err != nil {
			return errors.Wrap(err, "failed to delete collections")
		}
	}

	// Delete sites by user ID
	if err := s.Storage.DeleteSitesByUserID(userID); err != nil {
		return errors.Wrap(err, "failed to delete sites")
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

func mapToUserDB(user models.User) storage.UserDB {
	return storage.UserDB{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
	}
}
