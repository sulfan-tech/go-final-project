package service

import (
	"context"
	"errors"
	"go-final-project/internal/delivery/middlewares/logger"
	"go-final-project/internal/domain/photo/model"
	photorepo "go-final-project/internal/domain/photo/repo"
	"sync"
)

var (
	instance *PhotoService
	once     sync.Once
)

type PhotoService struct {
	repoPhoto photorepo.PhotoRepositoryImpl
	logger    logger.Logger
}

type PhotoServiceImpl interface {
	CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error)
	GetPhoto() ([]model.Photo, error)
	GetPhotoByID(ctx context.Context, id int) (model.Photo, error)
	UpdatePhoto(ctx context.Context, photo *model.Photo) (model.Photo, error)
	DeletePhoto(ctx context.Context, id string) error
}

func NewInstancePhotoService() PhotoServiceImpl {
	once.Do(func() {
		repoPhoto, err := photorepo.NewInstancePhotoRepository()
		if err != nil {
			panic(err)
		}
		logger := logger.NewLogger(logger.DebugLevel)

		instance = &PhotoService{
			repoPhoto: repoPhoto,
			logger:    logger,
		}
	})

	return instance
}

func (s *PhotoService) CreatePhoto(ctx context.Context, photo model.Photo) (model.Photo, error) {
	createdPhoto, err := s.repoPhoto.CreatePhoto(ctx, &photo)
	if err != nil {
		s.logger.Error("Failed to create photo: " + err.Error())
		return model.Photo{}, err
	}

	return *createdPhoto, nil
}

func (s *PhotoService) GetPhoto() ([]model.Photo, error) {
	photos, err := s.repoPhoto.GetPhoto()
	if err != nil {
		s.logger.Error("Failed to get photos:" + err.Error())
		return nil, errors.New("failed to get photos")
	}
	return photos, nil
}

func (s *PhotoService) GetPhotoByID(ctx context.Context, id int) (model.Photo, error) {
	photo, err := s.repoPhoto.GetPhotoByID(id)
	if err != nil {
		s.logger.Error("Failed to get photo by ID:" + err.Error())
		return model.Photo{}, errors.New("failed to get photo by ID")
	}
	return *photo, nil
}

func (s *PhotoService) UpdatePhoto(ctx context.Context, photo *model.Photo) (model.Photo, error) {
	updatedPhoto, err := s.repoPhoto.UpdatePhoto(*photo)
	if err != nil {
		s.logger.Error("Failed to update photo: " + err.Error())
		return model.Photo{}, errors.New("failed to update photo")
	}
	return *updatedPhoto, nil
}

func (s *PhotoService) DeletePhoto(ctx context.Context, id string) error {
	err := s.repoPhoto.DeletePhoto(id)
	if err != nil {
		s.logger.Error("Failed to delete photo:" + err.Error())
		return errors.New("failed to delete photo")
	}
	return nil
}
