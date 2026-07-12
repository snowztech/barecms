package storage

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func (s *Storage) CreateUser(user UserDB) error {
	created := s.DB.Create(&user)
	if created.Error != nil {
		return created.Error
	}
	return nil
}

func (s *Storage) GetUserByEmail(email string) (UserDB, error) {
	var user UserDB
	err := s.DB.First(&user, "email = ?", email).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *Storage) GetUserByID(id string) (UserDB, error) {
	var user UserDB
	err := s.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *Storage) RevokeToken(userID string) error {
	if err := s.DB.Model(&UserDB{}).Where("id = ?", userID).Update("token", "").Error; err != nil {
		return errors.Wrap(err, "failed to revoke token for user")
	}
	return nil
}

func (s *Storage) DeleteUserByID(id string) error {
	var user UserDB
	err := s.DB.First(&user, "id = ?", id).Error
	if err != nil {
		return err
	}
	deleted := s.DB.Where("id = ?", id).Delete(&UserDB{})
	if deleted.Error != nil {
		return deleted.Error
	}
	// todo: delete all user resources
	return nil
}

func (s *Storage) DeleteUserCascade(id string) ([]MediaFileDB, error) {
	var media []MediaFileDB
	err := s.DB.Transaction(func(tx *gorm.DB) error {
		siteIDs := tx.Model(&SiteDB{}).Select("id").Where("user_id = ?", id)
		collectionIDs := tx.Model(&CollectionDB{}).Select("id").Where("site_id IN (?)", siteIDs)
		if err := tx.Where("site_id IN (?)", siteIDs).Find(&media).Error; err != nil {
			return err
		}
		if err := tx.Where("collection_id IN (?)", collectionIDs).Delete(&EntryDB{}).Error; err != nil {
			return err
		}
		if err := tx.Where("site_id IN (?)", siteIDs).Delete(&CollectionDB{}).Error; err != nil {
			return err
		}
		if err := tx.Where("site_id IN (?)", siteIDs).Delete(&MediaFileDB{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", id).Delete(&SiteDB{}).Error; err != nil {
			return err
		}
		result := tx.Where("id = ?", id).Delete(&UserDB{})
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
