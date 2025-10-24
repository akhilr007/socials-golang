package store

import (
	"database/sql"
	"errors"

	"github.com/akhilr007/socials/internal/repository"
)

var (
	ErrNotFound        = errors.New("resource not found")
	ErrVersionConflict = errors.New("update version conflict")
)

type Storage interface {
	Users() repository.UserRepository
	Posts() repository.PostRepository
	Comments() repository.CommentRepository
}

type postgresStorage struct {
	db                *sql.DB
	userRepository    repository.UserRepository
	postRepository    repository.PostRepository
	commentRepository repository.CommentRepository
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

func (s *postgresStorage) Comments() repository.CommentRepository {
	if s.commentRepository == nil {
		s.commentRepository = newCommentRepositoryPG(s.db)
	}
	return s.commentRepository
}
