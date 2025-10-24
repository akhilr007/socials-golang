package store

import (
	"context"
	"database/sql"

	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/repository"
)

type commentRepository struct {
	db *sql.DB
}

func newCommentRepositoryPG(db *sql.DB) repository.CommentRepository {
	return &commentRepository{
		db: db,
	}
}

func (c *commentRepository) GetPostWithComments(ctx context.Context, id int64) ([]model.Comment, error) {
	query := `
		SELECT c.id, c.post_id, c.user_id, c.content, c.created_at, u.username, u.id FROM comments c
		JOIN users u on u.id = c.user_id
		WHERE c.post_id = $1
		ORDER BY c.created_at DESC;
	`

	ctx, cancel := context.WithTimeout(ctx, TimeoutQueryDuration)
	defer cancel()

	rows, err := c.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []model.Comment{}
	for rows.Next() {
		var c model.Comment
		c.User = model.User{}
		err := rows.Scan(&c.ID, &c.PostID, &c.UserID, &c.Content, &c.CreatedAt, &c.User.Username, &c.User.ID)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}

func (c *commentRepository) Create(ctx context.Context, comment *model.Comment) error {
	query := `
	INSERT INTO comments (post_id, user_id, content)
	VALUES ($1, $2, $3)
	RETURNING id, created_at`

	ctx, cancel := context.WithTimeout(ctx, TimeoutQueryDuration)
	defer cancel()

	err := c.db.QueryRowContext(
		ctx,
		query,
		comment.PostID,
		comment.UserID,
		comment.Content,
	).Scan(
		&comment.ID,
		&comment.CreatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}
