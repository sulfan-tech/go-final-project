package repo

import (
	"context"
	"go-final-project/config/psql"
	"go-final-project/internal/delivery/middlewares/logger"
	"go-final-project/internal/domain/comment/model"
	mp "go-final-project/internal/domain/photo/model"
	"sync"

	"gorm.io/gorm"
)

var (
	instance *CommentRepository
	once     sync.Once
)

type CommentRepository struct {
	db      *gorm.DB
	logging logger.Logger
}

type CommentRepositoryImpl interface {
	CreateComment(ctx context.Context, photoID uint, comment *model.Comment) (*model.Comment, error)
	GetComments(ctx context.Context) ([]model.Comment, error)
	UpdateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	DeleteComment(ctx context.Context, id uint) error
}

// NewCommentRepository creates a new instance of CommentRepository.
func NewInstanceCommentRepository() (CommentRepositoryImpl, error) {
	once.Do(func() {
		db, err := psql.SetupDatabase()
		if err != nil {
			panic(err)
		}
		logger := logger.NewLogger(logger.InfoLevel)

		db.AutoMigrate(&model.Comment{})
		instance = &CommentRepository{
			db:      db,
			logging: logger,
		}
	})

	return instance, nil
}

func (r *CommentRepository) CreateComment(ctx context.Context, photoID uint, comment *model.Comment) (*model.Comment, error) {
	// Retrieve the photo associated with the given photoID
	photo := mp.Photo{}
	if err := r.db.First(&photo, photoID).Error; err != nil {
		return nil, err
	}

	comment.PhotoID = photoID
	comment.Photo = photo

	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}

	return comment, nil
}

func (r *CommentRepository) GetComments(ctx context.Context) ([]model.Comment, error) {
	var comment []model.Comment
	if err := r.db.Find(&comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepository) UpdateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	if err := r.db.Save(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *CommentRepository) DeleteComment(ctx context.Context, id uint) error {
	if err := r.db.Delete(&model.Comment{}, id).Error; err != nil {
		return err
	}
	return nil
}
