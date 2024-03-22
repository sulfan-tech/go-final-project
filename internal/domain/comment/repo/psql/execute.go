package repocomment

import (
	"context"
	"go-final-project/config/psql"
	commentmodel "go-final-project/internal/domain/comment/model"

	"gorm.io/gorm"
)

type CommentPsqlRepository struct {
	db *gorm.DB
}

type PsqlUserRepositoryImpl interface {
	CreateComment(ctx context.Context, comment commentmodel.Comment) error
}

func NewInstanceUserRepository() (PsqlUserRepositoryImpl, error) {
	db, err := psql.SetupDatabase()
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&commentmodel.Comment{})
	return &CommentPsqlRepository{
		db: db,
	}, nil
}

func (r *CommentPsqlRepository) CreateComment(ctx context.Context, comment commentmodel.Comment) error {
	return r.db.Create(comment).Error
}
