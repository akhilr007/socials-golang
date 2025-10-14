package store

import (
	"context"
	"database/sql"
	"errors"

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

func (p *postRepository) GetByID(ctx context.Context, id int64) (*model.Post, error) {
	query := `
		SELECT id, title, content, tags, user_id, created_at, updated_at
		FROM posts
		WHERE id = $1
	`

	var post model.Post
	err := p.db.QueryRowContext(ctx, query, id).Scan(
		&post.ID,
		&post.Title,
		&post.Content,
		pq.Array(post.Tags),
		&post.UserID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return &post, nil
}
