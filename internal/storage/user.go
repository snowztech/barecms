package storage

import "github.com/pkg/errors"

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
