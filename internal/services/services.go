package services

import (
	"barecms/configs"
	"barecms/internal/storage"
)

type Service struct {
	Storage *storage.Storage
	Config  configs.AppConfig
}

func NewService(config configs.AppConfig) (*Service, error) {
	storage, err := storage.NewStorage(config.DatabaseURL)
	if err != nil {
		return nil, err
	}
	return &Service{Storage: storage, Config: config}, nil
}
