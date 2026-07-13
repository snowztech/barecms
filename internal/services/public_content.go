package services

import (
	"barecms/internal/models"
	"encoding/json"
)

func (s *Service) GetPublicEntries(siteSlug, collectionSlug string, page, limit int) (models.PublicEntryPage, error) {
	collection, err := s.Storage.GetCollectionBySiteAndSlug(siteSlug, collectionSlug)
	if err != nil {
		return models.PublicEntryPage{}, err
	}
	entriesDB, total, err := s.Storage.GetEntriesPage(collection.ID, limit, (page-1)*limit)
	if err != nil {
		return models.PublicEntryPage{}, err
	}
	entries := make([]map[string]any, len(entriesDB))
	for index, entry := range entriesDB {
		transformed, err := publicEntry(entry.ID, entry.Data)
		if err != nil {
			return models.PublicEntryPage{}, err
		}
		entries[index] = transformed
	}
	totalPages := int((total + int64(limit) - 1) / int64(limit))
	return models.PublicEntryPage{Entries: entries, Pagination: models.Pagination{Page: page, Limit: limit, Total: total, TotalPages: totalPages}}, nil
}

func (s *Service) GetPublicEntry(siteSlug, collectionSlug, entryID string) (map[string]any, error) {
	collection, err := s.Storage.GetCollectionBySiteAndSlug(siteSlug, collectionSlug)
	if err != nil {
		return nil, err
	}
	entry, err := s.Storage.GetEntryInCollection(entryID, collection.ID)
	if err != nil {
		return nil, err
	}
	return publicEntry(entry.ID, entry.Data)
}

func publicEntry(id string, data []byte) (map[string]any, error) {
	var entryData map[string]any
	if err := json.Unmarshal(data, &entryData); err != nil {
		return nil, err
	}
	transformed, err := transformEntryData(entryData)
	if err != nil {
		return nil, err
	}
	transformed["id"] = id
	return transformed, nil
}
