package repository

import (
	"context"

	"github.com/akhilr007/socials/internal/model"
)

type PostRepository interface {
	Create(ctx context.Context, post *model.Post) error
}
