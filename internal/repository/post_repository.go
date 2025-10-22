package repository

import (
	"context"

	"github.com/akhilr007/socials/internal/model"
)

type PostRepository interface {
	Create(context.Context, *model.Post) error
	GetByID(context.Context, int64) (*model.Post, error)
	Delete(context.Context, int64) error
}
