package handler

import (
	"net/http"

	"github.com/akhilr007/socials/internal/model"
	"github.com/akhilr007/socials/internal/service"
	"github.com/akhilr007/socials/internal/util"
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
