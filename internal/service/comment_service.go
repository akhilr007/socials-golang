package service

import (
	"context"

	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/repository"
)

type CommentService interface {
	GetPostWithComments(context.Context, int64) ([]model.Comment, error)
	CreateComment(context.Context, *model.Comment) error
}

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) CommentService {
	return &commentService{
		repository: repository,
	}
}

func (c *commentService) GetPostWithComments(ctx context.Context, id int64) ([]model.Comment, error) {
	return c.repository.GetPostWithComments(ctx, id)
}

func (c *commentService) CreateComment(ctx context.Context, comment *model.Comment) error {
	return c.repository.Create(ctx, comment)
}
