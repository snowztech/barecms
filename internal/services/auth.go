package services

import (
	"barecms/internal/models"
	"barecms/internal/storage"
	"barecms/internal/utils"

	"github.com/pkg/errors"
)

func (s *Service) Register(request models.RegisterRequest) error {

	hashedPassword, err := utils.GenerateHashedPassword(request.Password)
	if err != nil {
		return err
	}

	userDB := storage.UserDB{
		ID:       utils.GenerateUniqueID(),
		Email:    request.Email,
		Username: request.Username,
		Password: hashedPassword,
	}

	created := s.Storage.CreateUser(userDB)
	if created != nil {
		return created
	}

	return nil
}

func (s *Service) Login(email, password string) (models.User, error) {
	userDB, err := s.Storage.GetUserByEmail(email)
	if err != nil {
		return models.User{}, err
	}

	if err := utils.CompareHashAndPassword(userDB.Password, password); err != nil {
		return models.User{}, errors.New("invalid password")
	}

	user := mapToUser(userDB)

	return user, nil
}

func (s *Service) Logout(userID string) error {
	if err := s.Storage.RevokeToken(userID); err != nil {
		return errors.Wrap(err, "failed to logout user")
	}

	return nil
}
