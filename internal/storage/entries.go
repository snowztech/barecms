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

func (s *Storage) GetEntriesByCollectionID(collectionID string) ([]EntryDB, error) {
	var entries []EntryDB
	err := s.DB.Find(&entries, "collection_id = ?", collectionID).Error
	if err != nil {
		return nil, err
	}
	return entries, nil
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
