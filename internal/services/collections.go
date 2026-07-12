package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"barecms/internal/utils"
	"encoding/json"

	"github.com/pkg/errors"
	"gorm.io/datatypes"
)

func (s *Service) CreateCollection(request models.CreateCollectionRequest, userID string) error {
	if err := s.requireSiteOwner(userID, request.SiteID); err != nil {
		return err
	}

	// Validate field types
	for _, field := range request.Fields {
		if !models.IsValidFieldType(string(field.Type)) {
			return errors.New("invalid field type: " + string(field.Type))
		}
	}

	// Convert fields to JSON for storage
	fieldsJSON, err := json.Marshal(request.Fields)
	if err != nil {
		return err
	}

	collection := models.Collection{
		ID:     utils.GenerateUniqueID(),
		Name:   request.Name,
		Slug:   utils.Slugify(request.Name),
		SiteID: request.SiteID,
		Fields: request.Fields,
	}
	collectionDB := mapToCollectionDB(collection, fieldsJSON)
	if err := s.Storage.CreateCollection(collectionDB); err != nil {
		return err
	}
	return nil
}

func (s *Service) GetCollectionByID(collectionID, userID string) (models.Collection, error) {
	if err := s.requireCollectionOwner(userID, collectionID); err != nil {
		return models.Collection{}, err
	}
	collectionDB, err := s.Storage.GetCollection(collectionID)
	if err != nil {
		return models.Collection{}, err
	}
	return mapToCollection(collectionDB), nil
}

func (s *Service) GetCollectionsBySiteID(siteID, userID string) ([]models.Collection, error) {
	if err := s.requireSiteOwner(userID, siteID); err != nil {
		return nil, err
	}
	collectionsDB, err := s.Storage.GetCollectionsBySiteID(siteID)
	if err != nil {
		return nil, err
	}
	var collections []models.Collection
	for _, collectionDB := range collectionsDB {
		collections = append(collections, mapToCollection(collectionDB))
	}
	return collections, nil

}

func (s *Service) DeleteCollection(collectionID, userID string) error {
	if err := s.requireCollectionOwner(userID, collectionID); err != nil {
		return err
	}

	if err := s.Storage.DeleteEntriesByCollectionID(collectionID); err != nil {
		return err
	}

	if err := s.Storage.DeleteCollection(collectionID); err != nil {
		return err
	}

	return nil
}

func mapToCollectionDB(collection models.Collection, fieldsJSON []byte) storage.CollectionDB {
	return storage.CollectionDB{
		ID:     collection.ID,
		Name:   collection.Name,
		Slug:   collection.Slug,
		SiteID: collection.SiteID,
		Fields: datatypes.JSON(fieldsJSON),
	}
}

func mapToCollection(collectionDB storage.CollectionDB) models.Collection {
	var fields []models.Field
	if err := json.Unmarshal(collectionDB.Fields, &fields); err != nil {
		return models.Collection{}
	}

	return models.Collection{
		ID:     collectionDB.ID,
		Name:   collectionDB.Name,
		Slug:   collectionDB.Slug,
		SiteID: collectionDB.SiteID,
		Fields: fields,
	}
}
