package store

import (
	"context"
	"database/sql"

	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/repository"
	"github.com/lib/pq"
)

type postRepository struct {
	db *sql.DB
}

func newPostRepositoryPG(db *sql.DB) repository.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) Create(ctx context.Context, post *model.Post) error {
	query := `
		INSERT INTO posts (content, title, user_id, tags)
		VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at
	`

	err := p.db.QueryRowContext(
		ctx,
		query,
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}
	return nil
}
