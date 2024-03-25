package commentservice

import (
	"context"
	"go-final-project/internal/delivery/middlewares/logger"
	"go-final-project/internal/domain/comment/model"
	"go-final-project/internal/domain/comment/repo"
	photoRepo "go-final-project/internal/domain/photo/repo"
	"sync"
)

type CommentService struct {
	commentRepo repo.CommentRepositoryImpl
	photoRepo   photoRepo.PhotoRepositoryImpl
	logger      logger.Logger
}

var (
	instance *CommentService
	once     sync.Once
)

type CommentServiceImpl interface {
	CreateComment(ctx context.Context, photoId int, comment *model.Comment) (*model.Comment, error)
	GetComments(ctx context.Context) ([]model.Comment, error)
	UpdateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error)
	DeleteComment(ctx context.Context, id uint) error
}

func NewCommentService() CommentServiceImpl {
	once.Do(func() {
		repoComment, err := repo.NewInstanceCommentRepository()
		if err != nil {
			panic(err)
		}

		photoRepo, err := photoRepo.NewInstancePhotoRepository()
		if err != nil {
			panic(err)
		}

		logger := logger.NewLogger(logger.DebugLevel)

		instance = &CommentService{
			commentRepo: repoComment,
			photoRepo:   photoRepo,
			logger:      logger,
		}
	})

	return instance
}

func (s *CommentService) CreateComment(ctx context.Context, photoId int, comment *model.Comment) (*model.Comment, error) {
	photo, err := s.photoRepo.GetPhotoByID(photoId)
	if err != nil {
		return nil, err
	}

	createdComment, err := s.commentRepo.CreateComment(ctx, photo.ID, comment)
	if err != nil {
		return nil, err
	}
	return createdComment, nil
}

func (s *CommentService) GetComments(ctx context.Context) ([]model.Comment, error) {
	comment, err := s.commentRepo.GetComments(ctx)
	if err != nil {
		return nil, err
	}
	return comment, nil
}

func (s *CommentService) UpdateComment(ctx context.Context, comment *model.Comment) (*model.Comment, error) {
	updatedComment, err := s.commentRepo.UpdateComment(ctx, comment)
	if err != nil {
		return nil, err
	}
	return updatedComment, nil
}

func (s *CommentService) DeleteComment(ctx context.Context, id uint) error {
	err := s.commentRepo.DeleteComment(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
