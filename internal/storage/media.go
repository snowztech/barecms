package storage

func (s *Storage) CreateMediaFile(file *MediaFileDB) error {
	return s.DB.Create(file).Error
}

func (s *Storage) GetMediaFile(id string) (MediaFileDB, error) {
	var file MediaFileDB
	err := s.DB.Where("id = ?", id).First(&file).Error
	return file, err
}

func (s *Storage) GetMediaFilesBySiteID(siteID string) ([]MediaFileDB, error) {
	var files []MediaFileDB
	err := s.DB.Where("site_id = ?", siteID).Order("created_at DESC").Find(&files).Error
	return files, err
}

func (s *Storage) DeleteMediaFile(id string) error {
	return s.DB.Where("id = ?", id).Delete(&MediaFileDB{}).Error
}
