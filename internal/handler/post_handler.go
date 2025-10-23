package handler

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/akhilr007/socials/internal/middleware"
	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/service"
	"github.com/akhilr007/socials/internal/store"
	"github.com/akhilr007/socials/internal/util"
	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	postService    service.PostService
	commentService service.CommentService
}

func NewPostHandler(p service.PostService, c service.CommentService) *PostHandler {
	return &PostHandler{
		postService:    p,
		commentService: c,
	}
}

type CreatePostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags" validate:"dive,max=30"`
}

type UpdatePostPayload struct {
	Title   *string `json:"title" validate:"omitempty,max=100"`
	Content *string `json:"content" validate:"omitempty,max=1000"`
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	var payload CreatePostPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		util.BadRequestError(w, r, err)
		return
	}

	if err := util.Validate.Struct(payload); err != nil {
		util.BadRequestError(w, r, err)
		return
	}

	post := &model.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: change after auth
		UserID: 1,
	}
	ctx := r.Context()

	if err := h.postService.CreatePost(ctx, post); err != nil {
		util.InternalServerError(w, r, err)
		return
	}

	if err := util.WriteJSON(w, http.StatusCreated, post); err != nil {
		util.InternalServerError(w, r, err)
		return
	}
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)

	comments, err := h.commentService.GetPostWithComments(r.Context(), post.ID)
	if err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	post.Comments = comments

	if err := util.WriteJSON(w, http.StatusOK, post); err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *PostHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.InternalServerError(w, r, err)
		return
	}

	ctx := r.Context()
	err = h.postService.DeletePost(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			util.NotFoundError(w, r, err)
		default:
			util.InternalServerError(w, r, err)
		}
		return
	}

	if err := util.WriteJSON(w, http.StatusOK, "deleted successfully"); err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *PostHandler) UpdatePost(w http.ResponseWriter, r *http.Request) {
	post := getPostFromContext(r)
	log.Print(post)
	var payload UpdatePostPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		util.InternalServerError(w, r, err)
		return
	}

	if err := util.Validate.Struct(payload); err != nil {
		util.BadRequestError(w, r, err)
		return
	}

	if payload.Content != nil {
		post.Content = *payload.Content
	}

	if payload.Title != nil {
		post.Title = *payload.Title
	}

	if err := h.postService.UpdatePost(r.Context(), post); err != nil {
		util.InternalServerError(w, r, err)
		return
	}

	if err := util.WriteJSON(w, http.StatusOK, post); err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func getPostFromContext(r *http.Request) *model.Post {
	post, _ := r.Context().Value(middleware.PostContextKey).(*model.Post)
	return post
}
