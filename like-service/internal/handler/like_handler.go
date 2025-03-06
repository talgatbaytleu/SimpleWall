package handler

import (
	"encoding/json"
	"net/http"

	"liker/internal/service"
	"liker/pkg/apperrors"
	"liker/pkg/models"
	"liker/pkg/utils"
)

type likeHandler struct {
	likeService service.LikeServiceInterface
}

func NewLikeHandler(likeService service.LikeServiceInterface) *likeHandler {
	return &likeHandler{likeService: likeService}
}

func (h *likeHandler) PostLike(w http.ResponseWriter, r *http.Request) {
	// post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	// if err != nil {
	// 	utils.ResponseErrorJson(err, w)
	// }

	var like models.LikeType

	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	user_idStr := r.Header.Get("X-User-ID")

	err = h.likeService.ToLike(like.PostID, user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *likeHandler) GetLikesCount(w http.ResponseWriter, r *http.Request) {
	post_idStr := r.URL.Query().Get("post_id")
	if post_idStr == "" {
		utils.ResponseErrorJson(apperrors.ErrNoPostID, w)
		return
	}

	// post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	// if err != nil {
	// 	utils.ResponseErrorJson(err, w)
	// }

	jsonData, err := h.likeService.GetLikesCount(post_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func (h *likeHandler) GetLikesList(w http.ResponseWriter, r *http.Request) {
	post_idStr := r.URL.Query().Get("post_id")
	if post_idStr == "" {
		utils.ResponseErrorJson(apperrors.ErrNoPostID, w)
		return
	}

	// post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	// if err != nil {
	// 	utils.ResponseErrorJson(err, w)
	// }

	jsonData, err := h.likeService.GetLikesList(post_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonData)
}

func (h *likeHandler) DeleteLike(w http.ResponseWriter, r *http.Request) {
	// post_idStr, err := utils.GetURLVar(2, r.URL.Path)
	// if err != nil {
	// 	utils.ResponseErrorJson(err, w)
	// }
	var like models.LikeType

	err := json.NewDecoder(r.Body).Decode(&like)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	user_idStr := r.Header.Get("X-User-ID")

	err = h.likeService.ToUnlike(like.PostID, user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
