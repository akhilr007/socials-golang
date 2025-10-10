package store

import (
	"context"
	"database/sql"

	"github.com/akhilr007/socials/internal/repository"
)

type userRepository struct {
	db *sql.DB
}

func newUserRepositoryPG(db *sql.DB) repository.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Create(ctx context.Context) {

}
