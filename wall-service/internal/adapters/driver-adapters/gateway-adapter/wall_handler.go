package gatewayadapter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"wall/internal/ports"
	"wall/pkg/apperrors"
	"wall/pkg/utils"
)

type wallHandler struct {
	wallService ports.WallService
}

func NewWallHandler(wallService ports.WallService) *wallHandler {
	return &wallHandler{wallService: wallService}
}

func (h *wallHandler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}

func (h *wallHandler) GetUserWall(w http.ResponseWriter, r *http.Request) {
	user_idStr := r.URL.Query().Get("user_id")
	if user_idStr == "" {
		utils.ResponseErrorJson(apperrors.ErrNoPostId, w)
		return
	}

	jsonWall, err := h.wallService.GetUserWall(user_idStr)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jsonWall)
}
