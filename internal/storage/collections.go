package storage

import "gorm.io/datatypes"

func (s *Storage) CreateCollection(collection CollectionDB) error {
	created := s.DB.Create(&collection)
	if created.Error != nil {
		return created.Error
	}
	return nil
}

func (s *Storage) GetCollections() ([]CollectionDB, error) {
	var collections []CollectionDB
	retrieved := s.DB.Find(&collections)
	if retrieved.Error != nil {
		return nil, retrieved.Error
	}
	return collections, nil
}

func (s *Storage) GetCollection(id string) (CollectionDB, error) {
	var collection CollectionDB
	retrieved := s.DB.Where("id = ?", id).First(&collection)
	if retrieved.Error != nil {
		return CollectionDB{}, retrieved.Error
	}
	return collection, nil
}

func (s *Storage) UpdateCollection(id, name string, fields datatypes.JSON) error {
	return s.DB.Model(&CollectionDB{}).Where("id = ?", id).Updates(map[string]any{"name": name, "fields": fields}).Error
}

func (s *Storage) GetCollectionsBySiteID(siteID string) ([]CollectionDB, error) {
	var collections []CollectionDB
	retrieved := s.DB.Where("site_id = ?", siteID).Find(&collections)
	if retrieved.Error != nil {
		return nil, retrieved.Error
	}
	return collections, nil
}

func (s *Storage) DeleteCollection(id string) error {
	tx := s.DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// delete all entries for this collection
	if err := tx.Where("collection_id = ?", id).Delete(&EntryDB{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	// delete the collection
	if err := tx.Where("id = ?", id).Delete(&CollectionDB{}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (s *Storage) DeleteCollectionsBySiteID(siteId string) error {
	deleted := s.DB.Where("site_id = ?", siteId).Delete(&CollectionDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (s *Storage) GetCollectionsFromSitesIDs(siteIDs []string) ([]CollectionDB, error) {
	var collections []CollectionDB
	retrieved := s.DB.Where("site_id IN (?)", siteIDs).Find(&collections)
	if retrieved.Error != nil {
		return nil, retrieved.Error
	}
	return collections, nil
}

func (s *Storage) DeleteCollectionsBySiteIDs(siteIds []string) error {
	deleted := s.DB.Where("site_id IN (?)", siteIds).Delete(&CollectionDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
