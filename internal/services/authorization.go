package services

import "github.com/pkg/errors"

var ErrForbidden = errors.New("you do not have access to this resource")

func forbiddenUnless(allowed bool, err error) error {
	if err != nil {
		return err
	}
	if !allowed {
		return ErrForbidden
	}
	return nil
}

func (s *Service) requireSiteOwner(userID, siteID string) error {
	allowed, err := s.Storage.UserOwnsSite(userID, siteID)
	return forbiddenUnless(allowed, err)
}

func (s *Service) requireCollectionOwner(userID, collectionID string) error {
	allowed, err := s.Storage.UserOwnsCollection(userID, collectionID)
	return forbiddenUnless(allowed, err)
}

func (s *Service) requireEntryOwner(userID, entryID string) error {
	allowed, err := s.Storage.UserOwnsEntry(userID, entryID)
	return forbiddenUnless(allowed, err)
}
