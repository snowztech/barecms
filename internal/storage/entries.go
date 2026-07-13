package storage

func (s *Storage) CreateEntry(entry EntryDB) error {
	created := s.DB.Create(&entry)
	if created.Error != nil {
		return created.Error
	}
	return nil
}

func (s *Storage) GetEntryByID(id string) (EntryDB, error) {
	var entry EntryDB
	err := s.DB.First(&entry, "id = ?", id).Error
	if err != nil {
		return entry, err
	}
	return entry, nil
}

func (s *Storage) GetEntryInCollection(id, collectionID string) (EntryDB, error) {
	var entry EntryDB
	err := s.DB.Where("id = ? AND collection_id = ?", id, collectionID).First(&entry).Error
	return entry, err
}

func (s *Storage) GetEntriesByCollectionID(collectionID string) ([]EntryDB, error) {
	var entries []EntryDB
	err := s.DB.Find(&entries, "collection_id = ?", collectionID).Error
	if err != nil {
		return nil, err
	}
	return entries, nil
}

func (s *Storage) GetEntriesPage(collectionID string, limit, offset int) ([]EntryDB, int64, error) {
	var total int64
	if err := s.DB.Model(&EntryDB{}).Where("collection_id = ?", collectionID).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	var entries []EntryDB
	err := s.DB.Where("collection_id = ?", collectionID).
		Order("id ASC").Limit(limit).Offset(offset).Find(&entries).Error
	return entries, total, err
}

func (s *Storage) UpdateEntryData(id string, data []byte) error {
	return s.DB.Model(&EntryDB{}).Where("id = ?", id).Update("data", data).Error
}

func (s *Storage) DeleteEntry(id string) error {
	deleted := s.DB.Where("id = ?", id).Delete(&EntryDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (s *Storage) DeleteEntriesByCollectionIDs(collectionIDs []string) error {
	deleted := s.DB.Where("collection_id IN (?)", collectionIDs).Delete(&EntryDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}

func (s *Storage) DeleteEntriesByCollectionID(collectionID string) error {
	deleted := s.DB.Where("collection_id = ?", collectionID).Delete(&EntryDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	return nil
}
