package repository

import (
	"context"

	"github.com/akhilr007/socials/internal/model"
)

type CommentRepository interface {
	GetPostWithComments(context.Context, int64) ([]model.Comment, error)
}
