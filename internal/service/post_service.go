package service

import (
	"context"

	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/repository"
)

type PostService interface {
	CreatePost(ctx context.Context, post *model.Post) error
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
