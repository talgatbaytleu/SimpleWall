package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"commenter/internal/service"
	"commenter/pkg/apperrors"
	"commenter/pkg/logger"
	"commenter/pkg/utils"
)

type commentHandler struct {
	commentService service.CommentServiceInterface
}

func NewCommentHandler(commentService service.CommentServiceInterface) *commentHandler {
	return &commentHandler{commentService: commentService}
}

func (h *commentHandler) PostComment(w http.ResponseWriter, r *http.Request) {
	user_idStr := r.Header.Get("X-User-ID")

	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	err = h.commentService.CreateComment(bytes.NewReader(bodyBytes), user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	logger.LogMessage("CreateComment: " + user_idStr + ": successful")

	err = h.commentService.SendNotification(bytes.NewReader(bodyBytes), user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	logger.LogMessage("SendNotification: " + user_idStr + ": successful")

	w.WriteHeader(http.StatusOK)
}

func (h *commentHandler) PutComment(w http.ResponseWriter, r *http.Request) {
	comment_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	user_idStr := r.Header.Get("X-User-ID")

	err = h.commentService.UpdateComment(r.Body, comment_idStr, user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	logger.LogMessage("UpdateComment: " + user_idStr + ": successful")

	w.WriteHeader(http.StatusOK)
}

func (h *commentHandler) GetComment(w http.ResponseWriter, r *http.Request) {
	comment_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	jsonComment, err := h.commentService.GetCommentById(comment_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	logger.LogMessage("GetCommentById: " + comment_idStr + ": successful")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonComment)
}

func (h *commentHandler) GetComments(w http.ResponseWriter, r *http.Request) {
	post_idStr := r.URL.Query().Get("post_id")
	if post_idStr == "" {
		utils.ResponseErrorJson(apperrors.ErrNoPostId, w)
		return
	}

	jsonComment, err := h.commentService.GetPostComments(post_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	logger.LogMessage("GetPostComments: " + post_idStr + ": successful")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonComment)
}

func (h *commentHandler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	comment_idStr, err := utils.GetURLVar(2, r.URL.Path)
	if err != nil {
		utils.ResponseErrorJson(err, w)
	}
	user_idStr := r.Header.Get("X-User-ID")

	err = h.commentService.DeleteComment(comment_idStr, user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}
	logger.LogMessage("DeleteComment: " + comment_idStr + ": successful")

	w.WriteHeader(http.StatusNoContent)
}
