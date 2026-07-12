package storage

import "gorm.io/gorm"

func (s *Storage) CreateSite(site SiteDB) error {
	created := s.DB.Create(&site)
	if created.Error != nil {
		return created.Error
	}
	return nil
}

func (s *Storage) GetSites() ([]SiteDB, error) {
	var sites []SiteDB
	retrieved := s.DB.Find(&sites)
	if retrieved.Error != nil {
		return nil, retrieved.Error
	}
	return sites, nil
}

func (s *Storage) GetSite(id string) (SiteDB, error) {
	var site SiteDB
	retrieved := s.DB.Where("id = ?", id).First(&site)
	if retrieved.Error != nil {
		return SiteDB{}, retrieved.Error
	}
	return site, nil
}

func (s *Storage) DeleteSite(id string) error {
	deleted := s.DB.Where("id = ?", id).Delete(&SiteDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (s *Storage) DeleteSiteCascade(id string) ([]MediaFileDB, error) {
	var media []MediaFileDB
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("site_id = ?", id).Find(&media).Error; err != nil {
			return err
		}
		collectionIDs := tx.Model(&CollectionDB{}).Select("id").Where("site_id = ?", id)
		if err := tx.Where("collection_id IN (?)", collectionIDs).Delete(&EntryDB{}).Error; err != nil {
			return err
		}
		if err := tx.Where("site_id = ?", id).Delete(&CollectionDB{}).Error; err != nil {
			return err
		}
		if err := tx.Where("site_id = ?", id).Delete(&MediaFileDB{}).Error; err != nil {
			return err
		}
		result := tx.Where("id = ?", id).Delete(&SiteDB{})
		if result.Error != nil {
			return result.Error
		}
		if result.RowsAffected != 1 {
			return gorm.ErrRecordNotFound
		}
		return nil
	})
	return media, err
}

func (s *Storage) DeleteSitesByUserID(userID string) error {
	deleted := s.DB.Where("user_id = ?", userID).Delete(&SiteDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (s *Storage) GetSitesByUserID(userID string) ([]SiteDB, error) {
	var sites []SiteDB
	retrieved := s.DB.Where("user_id = ?", userID).Find(&sites)
	if retrieved.Error != nil {
		return nil, retrieved.Error
	}
	return sites, nil
}

func (s *Storage) UserOwnsSite(userID, siteID string) (bool, error) {
	var count int64
	err := s.DB.Model(&SiteDB{}).
		Where("id = ? AND user_id = ?", siteID, userID).
		Count(&count).Error
	return count > 0, err
}

func (s *Storage) UserOwnsCollection(userID, collectionID string) (bool, error) {
	var count int64
	err := s.DB.Table("collections").
		Joins("JOIN sites ON sites.id = collections.site_id").
		Where("collections.id = ? AND sites.user_id = ?", collectionID, userID).
		Count(&count).Error
	return count > 0, err
}

func (s *Storage) UserOwnsEntry(userID, entryID string) (bool, error) {
	var count int64
	err := s.DB.Table("entries").
		Joins("JOIN collections ON collections.id = entries.collection_id").
		Joins("JOIN sites ON sites.id = collections.site_id").
		Where("entries.id = ? AND sites.user_id = ?", entryID, userID).
		Count(&count).Error
	return count > 0, err
}

func (s *Storage) GetSiteBySlug(slug string) (SiteDB, error) {
	var site SiteDB
	retrieved := s.DB.Where("slug = ?", slug).First(&site)
	if retrieved.Error != nil {
		return SiteDB{}, retrieved.Error
	}
	return site, nil
}
