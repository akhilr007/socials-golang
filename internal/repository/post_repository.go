package repository

import "context"

type PostRepository interface {
	Create(ctx context.Context)
}
