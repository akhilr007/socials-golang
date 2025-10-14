package handler

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/service"
	"github.com/akhilr007/socials/internal/store"
	"github.com/akhilr007/socials/internal/util"
	"github.com/go-chi/chi/v5"
)

type PostHandler struct {
	service service.PostService
}

func NewPostHandler(s service.PostService) *PostHandler {
	return &PostHandler{
		service: s,
	}
}

type CreatePostPayload struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
}

func (h *PostHandler) CreatePost(w http.ResponseWriter, r *http.Request) {

	var payload CreatePostPayload
	if err := util.ReadJSON(w, r, &payload); err != nil {
		util.WriteJSONError(w, http.StatusBadRequest, err.Error())
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

	if err := h.service.CreatePost(ctx, post); err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := util.WriteJSON(w, http.StatusCreated, post); err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}

func (h *PostHandler) GetPostByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}

	ctx := r.Context()

	post, err := h.service.GetByID(ctx, id)
	if err != nil {
		switch {
		case errors.Is(err, store.ErrNotFound):
			util.WriteJSONError(w, http.StatusNotFound, err.Error())
		default:
			util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}

	if err := util.WriteJSON(w, http.StatusOK, post); err != nil {
		util.WriteJSONError(w, http.StatusInternalServerError, err.Error())
		return
	}
}
