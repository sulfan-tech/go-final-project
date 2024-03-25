package repo

import (
	"context"
	"go-final-project/config/psql"
	"go-final-project/internal/delivery/middlewares/logger"
	"go-final-project/internal/domain/social_media/model"
	"sync"

	"gorm.io/gorm"
)

var (
	instanceSocialMediaRepo *SocialMediaRepository
	onceSocialMediaRepo     sync.Once
)

type SocialMediaRepository struct {
	db      *gorm.DB
	logging logger.Logger
}

type SocialMediaRepositoryImpl interface {
	CreateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error)
	GetSocialMedias(ctx context.Context) ([]model.SocialMedia, error)
	UpdateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error)
	PreloadUser(ctx context.Context, socialMedia *model.SocialMedia) error
	Delete(ctx context.Context, id uint) error
}

func NewSocialMediaRepository() (SocialMediaRepositoryImpl, error) {
	onceSocialMediaRepo.Do(func() {
		db, err := psql.SetupDatabase()
		if err != nil {
			panic(err)
		}

		logger := logger.NewLogger(logger.InfoLevel)

		db.AutoMigrate(&model.SocialMedia{})
		instanceSocialMediaRepo = &SocialMediaRepository{
			db:      db,
			logging: logger,
		}
	})
	return instanceSocialMediaRepo, nil
}

func (r *SocialMediaRepository) CreateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := r.db.Create(socialMedia).Error; err != nil {
		return nil, err
	}
	return socialMedia, nil
}

func (s *SocialMediaRepository) GetSocialMedias(ctx context.Context) ([]model.SocialMedia, error) {
	var socialMedia []model.SocialMedia
	if err := s.db.Find(&socialMedia).Error; err != nil {
		return nil, err
	}

	if err := s.db.Preload("User").Find(&socialMedia).Error; err != nil {
		return nil, err
	}

	return socialMedia, nil
}

func (r *SocialMediaRepository) UpdateSocialMedia(ctx context.Context, socialMedia *model.SocialMedia) (*model.SocialMedia, error) {
	if err := r.db.Save(socialMedia).Error; err != nil {
		return nil, err
	}
	return socialMedia, nil
}

func (r *SocialMediaRepository) PreloadUser(ctx context.Context, socialMedia *model.SocialMedia) error {
	return r.db.Preload("User").First(socialMedia, socialMedia.ID).Error
}

func (r *SocialMediaRepository) Delete(ctx context.Context, id uint) error {
	if err := r.db.Delete(&model.SocialMedia{}, id).Error; err != nil {
		return err
	}
	return nil
}
