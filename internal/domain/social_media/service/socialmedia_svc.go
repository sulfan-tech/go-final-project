package service

import (
	"context"
	"go-final-project/internal/delivery/middlewares/logger"
	"go-final-project/internal/domain/social_media/model"
	"go-final-project/internal/domain/social_media/repo"

	"sync"
)

var (
	instanceSocialMediaService *SocialMediaService
	onceSocialMediaService     sync.Once
)

type SocialMediaService struct {
	socialMediaRepo repo.SocialMediaRepositoryImpl
	logger          logger.Logger
}

type SocialMediaServiceImpl interface {
	CreateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error)
	GetSocialMedias(ctx context.Context) ([]model.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error)
	DeleteSocialMedia(ctx context.Context, id uint) error
}

func NewSocialMediaService() SocialMediaServiceImpl {
	onceSocialMediaService.Do(func() {
		repoSocialMedia, err := repo.NewSocialMediaRepository()
		if err != nil {
			panic(err)
		}
		logger := logger.NewLogger(logger.DebugLevel)

		instanceSocialMediaService = &SocialMediaService{
			socialMediaRepo: repoSocialMedia,
			logger:          logger,
		}
	})

	return instanceSocialMediaService
}

func (s *SocialMediaService) CreateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := socialMedia.Validate(); err != nil {
		return nil, err
	}
	return s.socialMediaRepo.CreateSocialMedia(ctx, socialMedia)
}

func (s *SocialMediaService) GetSocialMedias(ctx context.Context) ([]model.SocialMedia, error) {
	socialMedia, err := s.socialMediaRepo.GetSocialMedias(ctx)
	if err != nil {
		return nil, err
	}

	for i := range socialMedia {
		if err := s.socialMediaRepo.PreloadUser(ctx, &socialMedia[i]); err != nil {
			return nil, err
		}
	}

	return socialMedia, nil
}

func (s *SocialMediaService) UpdateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := socialMedia.Validate(); err != nil {
		return nil, err
	}
	return s.socialMediaRepo.UpdateSocialMedia(ctx, socialMedia)
}

func (s *SocialMediaService) DeleteSocialMedia(ctx context.Context, id uint) error {
	return s.socialMediaRepo.Delete(ctx, id)
}
