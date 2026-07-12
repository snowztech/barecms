package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"barecms/internal/utils"
	"encoding/json"

	"gorm.io/datatypes"
)

func (s *Service) CreateCollection(request models.CreateCollectionRequest, userID string) error {
	if err := s.requireSiteOwner(userID, request.SiteID); err != nil {
		return err
	}

	if err := validateCollectionSchema(request.Name, request.Fields); err != nil {
		return err
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

func (s *Service) UpdateCollection(id, userID string, request models.UpdateCollectionRequest) (models.Collection, error) {
	if err := s.requireCollectionOwner(userID, id); err != nil {
		return models.Collection{}, err
	}
	if err := validateCollectionSchema(request.Name, request.Fields); err != nil {
		return models.Collection{}, err
	}
	entries, err := s.Storage.GetEntriesByCollectionID(id)
	if err != nil {
		return models.Collection{}, err
	}
	for _, entry := range entries {
		if err := validateEntryData(json.RawMessage(entry.Data), request.Fields); err != nil {
			return models.Collection{}, &ValidationError{Fields: map[string]string{"fields": "schema is incompatible with existing entry " + entry.ID}}
		}
	}
	fieldsJSON, err := json.Marshal(request.Fields)
	if err != nil {
		return models.Collection{}, err
	}
	if err := s.Storage.UpdateCollection(id, request.Name, datatypes.JSON(fieldsJSON)); err != nil {
		return models.Collection{}, err
	}
	return s.GetCollectionByID(id, userID)
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

	return s.Storage.DeleteCollection(collectionID)
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
