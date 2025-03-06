package handler

import (
	"encoding/json"
	"net/http"

	"poster/internal/service"
	"poster/pkg/utils"
)

type postHandler struct {
	postService service.PostServiceInterface
}

func NewPostService(postService service.PostServiceInterface) *postHandler {
	return &postHandler{postService: postService}
}

func (h *postHandler) PostPost(w http.ResponseWriter, r *http.Request) {
	description, imageFile, err := utils.ParseFormData(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	user_idStr := r.Header.Get("X-User-ID")

	err = h.postService.CreatePost(imageFile, description, user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *postHandler) PutPost(w http.ResponseWriter, r *http.Request) {
	post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	description, imageFile, err := utils.ParseFormData(r)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	user_idStr := r.Header.Get("X-User-ID")

	err = h.postService.UpdatePost(imageFile, description, user_idStr, post_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *postHandler) GetPost(w http.ResponseWriter, r *http.Request) {
	post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	jsonData, err := h.postService.RetrievePost(post_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func (h *postHandler) DeletePost(w http.ResponseWriter, r *http.Request) {
	post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}
	user_idStr := r.Header.Get("X-User-ID")

	err = h.postService.RemovePost(user_idStr, post_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	w.WriteHeader(http.StatusNoContent)
}
