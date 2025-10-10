package store

import (
	"context"
	"database/sql"

	"github.com/akhilr007/socials/internal/model"
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

func (u *userRepository) Create(ctx context.Context, user *model.User) error {
	query := `
		INSERT INTO users (username, password, email) VALUES ($1, $2, $3) 
		RETURNING id, created_at
	`

	err := u.db.QueryRowContext(
		ctx,
		query,
		user.Username,
		user.Password,
		user.Email,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)
	if err != nil {
		return err
	}
	return nil
}
