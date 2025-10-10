package store

import (
	"context"
	"database/sql"

	"github.com/akhilr007/socials/internal/repository"
)

type postRepository struct {
	db *sql.DB
}

func newPostRepositoryPG(db *sql.DB) repository.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (r *postRepository) Create(ctx context.Context) {

}
