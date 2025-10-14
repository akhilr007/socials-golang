package store

import (
	"database/sql"
	"errors"

	"github.com/akhilr007/socials/internal/repository"
)

var (
	ErrNotFound = errors.New("resource not found")
)

type Storage interface {
	Users() repository.UserRepository
	Posts() repository.PostRepository
}

type postgresStorage struct {
	db             *sql.DB
	userRepository repository.UserRepository
	postRepository repository.PostRepository
}

func NewPostgresStorage(db *sql.DB) Storage {
	return &postgresStorage{
		db: db,
	}
}

func (s *postgresStorage) Users() repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = newUserRepositoryPG(s.db)
	}
	return s.userRepository
}

func (s *postgresStorage) Posts() repository.PostRepository {
	if s.postRepository == nil {
		s.postRepository = newPostRepositoryPG(s.db)
	}
	return s.postRepository
}
