package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"barecms/internal/utils"
	"encoding/json"
	"fmt"
)

func (s *Service) GetSites(userID string) ([]models.Site, error) {
	sitesDB, err := s.Storage.GetSitesByUserID(userID)
	if err != nil {
		return nil, err
	}

	sites := make([]models.Site, 0)
	for _, siteDB := range sitesDB {
		site := mapToSite(siteDB)
		sites = append(sites, site)
	}

	return sites, nil
}

func (s *Service) GetSite(id, userID string) (models.Site, error) {
	if err := s.requireSiteOwner(userID, id); err != nil {
		return models.Site{}, err
	}
	siteDB, err := s.Storage.GetSite(id)
	if err != nil {
		return models.Site{}, err
	}

	site := mapToSite(siteDB)
	return site, nil
}

func (s *Service) GetSiteWithCollections(id, userID string) (map[string]interface{}, error) {
	// Get site
	site, err := s.GetSite(id, userID)
	if err != nil {
		return nil, err
	}

	// Get collections for this site
	collections, err := s.GetCollectionsBySiteID(id, userID)
	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"site":        site,
		"collections": collections,
	}, nil
}

func (s *Service) CreateSite(request models.CreateSiteRequest, userID string) error {
	newSite := models.Site{
		ID:     utils.GenerateUniqueID(),
		Name:   request.Name,
		Slug:   utils.Slugify(request.Name),
		UserID: userID,
	}
	siteDB := mapToSiteDB(newSite)
	if err := s.Storage.CreateSite(siteDB); err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteSite(id, userID string) error {
	if err := s.requireSiteOwner(userID, id); err != nil {
		return err
	}

	if err := s.Storage.DeleteSite(id); err != nil {
		return err
	}

	// Delete collections associated with the site ID
	siteCollections, err := s.Storage.GetCollectionsBySiteID(id)
	if err != nil {
		return err
	}
	var collectionIds []string
	for _, collection := range siteCollections {
		collectionIds = append(collectionIds, collection.ID)
	}

	// Delete entries associated with the collection IDs
	if err := s.Storage.DeleteEntriesByCollectionIDs(collectionIds); err != nil {
		return err
	}

	// Delete collections
	if err := s.Storage.DeleteCollectionsBySiteID(id); err != nil {
		return err
	}
	return nil
}

func mapToSiteDB(site models.Site) storage.SiteDB {
	return storage.SiteDB{
		ID:     site.ID,
		Name:   site.Name,
		Slug:   site.Slug,
		UserID: site.UserID,
	}
}

func mapToSite(siteDB storage.SiteDB) models.Site {
	return models.Site{
		ID:   siteDB.ID,
		Name: siteDB.Name,
		Slug: siteDB.Slug,
	}
}

func (s *Service) GetSiteData(slug string) (*models.SiteData, error) {
	// Get site by slug
	siteDB, err := s.Storage.GetSiteBySlug(slug)
	if err != nil {
		return nil, err
	}

	siteData := models.SiteData{
		ID:   siteDB.ID,
		Name: siteDB.Name,
		Slug: siteDB.Slug,
	}

	// Get collections by site ID
	collectionsDB, err := s.Storage.GetCollectionsBySiteID(siteDB.ID)
	if err != nil {
		return nil, err
	}

	mData := make(map[string]interface{})
	for _, collectionDB := range collectionsDB {
		// Get entries by collection ID
		entriesDB, err := s.Storage.GetEntriesByCollectionID(collectionDB.ID)
		if err != nil {
			return nil, err
		}

		var entries []map[string]interface{}
		for _, entryDB := range entriesDB {
			var entryData map[string]interface{}
			if err := json.Unmarshal(entryDB.Data, &entryData); err != nil {
				return nil, err
			}

			transformedData, err := transformEntryData(entryData)
			if err != nil {
				return nil, err
			}

			entries = append(entries, transformedData)
		}

		mData[collectionDB.Slug] = entries
	}

	siteData.Data = mData

	return &siteData, nil
}

func transformEntryData(entryData map[string]interface{}) (map[string]interface{}, error) {
	transformedData := make(map[string]interface{})

	for key, value := range entryData {
		field, ok := value.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("invalid format for field %s", key)
		}

		val, ok := field["value"]
		if !ok {
			return nil, fmt.Errorf("missing 'value' for field %s", key)
		}

		transformedData[key] = val
	}

	return transformedData, nil
}
