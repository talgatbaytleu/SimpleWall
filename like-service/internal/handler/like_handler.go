package handler

import (
	"encoding/json"
	"net/http"

	"liker/internal/service"
	"liker/pkg/utils"
)

type likeHandler struct {
	likeService service.LikeServiceInterface
}

func NewLikeHandler(likeService service.LikeServiceInterface) *likeHandler {
	return &likeHandler{likeService: likeService}
}

func (h *likeHandler) PostLike(w http.ResponseWriter, r *http.Request) {
	post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	user_idStr := r.Header.Get("X-User-ID")

	err = h.likeService.ToLike(post_idStr, user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	w.WriteHeader(http.StatusOK)
}

func (h *likeHandler) GetLike(w http.ResponseWriter, r *http.Request) {
	post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	jsonData, err := h.likeService.GetLikeCount(post_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func (h *likeHandler) DeleteLike(w http.ResponseWriter, r *http.Request) {
	post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}
	user_idStr := r.Header.Get("X-User-ID")

	err = h.likeService.ToUnlike(post_idStr, user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}

	w.WriteHeader(http.StatusNoContent)
}
