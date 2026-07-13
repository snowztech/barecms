package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"barecms/internal/utils"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
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

func (s *Service) GetSitesPage(userID string, page, limit int) (models.SitePage, error) {
	sitesDB, total, err := s.Storage.GetSitesPage(userID, limit, (page-1)*limit)
	if err != nil {
		return models.SitePage{}, err
	}
	sites := make([]models.Site, len(sitesDB))
	for index, site := range sitesDB {
		sites[index] = mapToSite(site)
	}
	return models.SitePage{Sites: sites, Pagination: pagination(total, page, limit)}, nil
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
	if strings.TrimSpace(request.Name) == "" {
		return &ValidationError{Fields: map[string]string{"name": "is required"}}
	}
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

func (s *Service) UpdateSite(id, userID string, request models.UpdateSiteRequest) (models.Site, error) {
	if err := s.requireSiteOwner(userID, id); err != nil {
		return models.Site{}, err
	}
	name := strings.TrimSpace(request.Name)
	if name == "" {
		return models.Site{}, &ValidationError{Fields: map[string]string{"name": "is required"}}
	}
	if err := s.Storage.UpdateSiteName(id, name); err != nil {
		return models.Site{}, err
	}
	site, err := s.Storage.GetSite(id)
	if err != nil {
		return models.Site{}, err
	}
	return mapToSite(site), nil
}

func (s *Service) DeleteSite(id, userID string) error {
	if err := s.requireSiteOwner(userID, id); err != nil {
		return err
	}

	media, err := s.Storage.DeleteSiteCascade(id)
	if err != nil {
		return err
	}
	for _, file := range media {
		if err := os.Remove(filepath.Join(s.Config.UploadsDir, file.StoredName)); err != nil && !os.IsNotExist(err) {
			slog.Warn("Could not remove deleted site's media file", "file", file.StoredName, "error", err)
		}
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
