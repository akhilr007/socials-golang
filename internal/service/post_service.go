package service

import (
	"context"

	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/repository"
)

type PostService interface {
	CreatePost(ctx context.Context, post *model.Post) error
	GetByID(ctx context.Context, id int64) (*model.Post, error)
	DeletePost(ctx context.Context, id int64) error
	UpdatePost(ctx context.Context, post *model.Post) error
}

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) PostService {
	return &postService{
		repository: repository,
	}
}

func (s *postService) CreatePost(ctx context.Context, post *model.Post) error {
	return s.repository.Create(ctx, post)
}

func (s *postService) GetByID(ctx context.Context, id int64) (*model.Post, error) {
	return s.repository.GetByID(ctx, id)
}

func (s *postService) DeletePost(ctx context.Context, id int64) error {
	return s.repository.Delete(ctx, id)
}

func (s *postService) UpdatePost(ctx context.Context, post *model.Post) error {
	return s.repository.Update(ctx, post)
}
