package photorepo

import (
	"go-final-project/config/psql"
	"go-final-project/internal/delivery/middlewares/logger"
	"go-final-project/internal/domain/photo/model"
	"sync"

	"gorm.io/gorm"
)

var (
	instance *PhotoRepository
	once     sync.Once
)

type PhotoRepository struct {
	db      *gorm.DB
	logging logger.Logger
}

type PhotoRepositoryImpl interface {
	CreatePhoto(photo *model.Photo) (*model.Photo, error)
	GetPhoto() ([]model.Photo, error)
}

func NewInstancePhotoRepository() (PhotoRepositoryImpl, error) {
	once.Do(func() {
		db, err := psql.SetupDatabase()
		if err != nil {
			panic(err)
		}
		logger := logger.NewLogger(logger.InfoLevel)

		db.AutoMigrate(&model.Photo{})
		instance = &PhotoRepository{
			db:      db,
			logging: logger,
		}
	})

	return instance, nil
}

func (p *PhotoRepository) CreatePhoto(photo *model.Photo) (*model.Photo, error) {
	if err := p.db.Table("photo").Create(photo).Error; err != nil {
		p.logging.Error("Failed to create photo: " + err.Error())
		return nil, err
	}
	return photo, nil
}

func (p *PhotoRepository) GetPhoto() ([]model.Photo, error) {
	var photos []model.Photo
	if err := p.db.Find(&photos).Error; err != nil {
		p.logging.Error("Failed to get photo: " + err.Error())
		return nil, err
	}
	return nil, nil
}

func (p *PhotoRepository) GetPhotoByID(id string) (*model.Photo, error) {
	var photo model.Photo
	result := p.db.First(photo, id)
	if result.Error != nil {
		return nil, result.Error
	}

	return &photo, nil
}

func (p *PhotoRepository) UpdatePhoto(photo model.Photo) (*model.Photo, error) {
	result := p.db.Save(photo)
	if result.Error != nil {
		return nil, result.Error
	}
	return &photo, result.Error
}

func (p *PhotoRepository) DeletePhoto(id string) error {
	result := p.db.Delete(&model.Photo{}, id)
	return result.Error
}
