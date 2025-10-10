package repository

import (
	"context"

	"github.com/akhilr007/socials/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) error
}
