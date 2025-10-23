package middleware

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/akhilr007/socials/internal/service"
	"github.com/akhilr007/socials/internal/store"
	"github.com/akhilr007/socials/internal/util"
	"github.com/go-chi/chi/v5"
)

type PostMiddleware struct {
	postService service.PostService
}

func NewPostMiddleware(postService service.PostService) *PostMiddleware {
	return &PostMiddleware{
		postService: postService,
	}
}

type postCtxKey string

const PostContextKey postCtxKey = "post"

func (pm *PostMiddleware) PostsContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		idParam := chi.URLParam(r, "postID")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			util.InternalServerError(w, r, err)
			return
		}

		ctx := r.Context()

		post, err := pm.postService.GetByID(ctx, id)
		if err != nil {
			switch {
			case errors.Is(err, store.ErrNotFound):
				util.NotFoundError(w, r, err)
			default:
				util.InternalServerError(w, r, err)
			}
			return
		}

		ctx = context.WithValue(ctx, PostContextKey, post)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
