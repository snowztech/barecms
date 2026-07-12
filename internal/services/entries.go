package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"barecms/internal/utils"
	"encoding/json"

	"gorm.io/datatypes"
)

func (s *Service) CreateEntry(request *models.CreateEntryRequest, userID string) error {
	if err := s.requireCollectionOwner(userID, request.CollectionID); err != nil {
		return err
	}
	collectionDB, err := s.Storage.GetCollection(request.CollectionID)
	if err != nil {
		return err
	}
	collection := mapToCollection(collectionDB)
	if err := validateEntryData(request.Data, collection.Fields); err != nil {
		return err
	}
	entry := models.Entry{
		ID:           utils.GenerateUniqueID(),
		CollectionID: request.CollectionID,
		Data:         request.Data,
	}
	entryDB := mapToEntryDB(entry)

	if err := s.Storage.CreateEntry(entryDB); err != nil {
		return err
	}

	return nil
}

func (s *Service) GetEntryByID(id, userID string) (models.Entry, error) {
	if err := s.requireEntryOwner(userID, id); err != nil {
		return models.Entry{}, err
	}
	entryDB, err := s.Storage.GetEntryByID(id)
	if err != nil {
		return models.Entry{}, err
	}
	return mapToEntry(entryDB), nil
}

func (s *Service) GetEntriesByCollectionID(collectionID, userID string) ([]models.Entry, error) {
	if err := s.requireCollectionOwner(userID, collectionID); err != nil {
		return nil, err
	}
	entriesDB, err := s.Storage.GetEntriesByCollectionID(collectionID)
	if err != nil {
		return nil, err
	}
	entries := make([]models.Entry, len(entriesDB))
	for i, entryDB := range entriesDB {
		entries[i] = mapToEntry(entryDB)
	}
	return entries, nil
}

func (s *Service) DeleteEntry(id, userID string) error {
	if err := s.requireEntryOwner(userID, id); err != nil {
		return err
	}
	return s.Storage.DeleteEntry(id)
}

func (s *Service) UpdateEntry(id, userID string, request models.UpdateEntryRequest) (models.Entry, error) {
	if err := s.requireEntryOwner(userID, id); err != nil {
		return models.Entry{}, err
	}
	entryDB, err := s.Storage.GetEntryByID(id)
	if err != nil {
		return models.Entry{}, err
	}
	collectionDB, err := s.Storage.GetCollection(entryDB.CollectionID)
	if err != nil {
		return models.Entry{}, err
	}
	if err := validateEntryData(request.Data, mapToCollection(collectionDB).Fields); err != nil {
		return models.Entry{}, err
	}
	if err := s.Storage.UpdateEntryData(id, request.Data); err != nil {
		return models.Entry{}, err
	}
	entryDB.Data = datatypes.JSON(request.Data)
	return mapToEntry(entryDB), nil
}

func mapToEntryDB(entry models.Entry) storage.EntryDB {
	return storage.EntryDB{
		ID:           entry.ID,
		CollectionID: entry.CollectionID,
		Data:         datatypes.JSON(entry.Data),
	}
}

func mapToEntry(entryDB storage.EntryDB) models.Entry {
	return models.Entry{
		ID:           entryDB.ID,
		CollectionID: entryDB.CollectionID,
		Data:         json.RawMessage(entryDB.Data),
	}
}
