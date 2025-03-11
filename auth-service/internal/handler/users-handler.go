package handler

import (
	"encoding/json"
	"net/http"
	"strings"

	"auth-service/internal/logic"
	"auth-service/pkg/logger"
	"auth-service/pkg/utils"
)

type userHandler struct {
	userLogic logic.UserLogicInterface
}

func NewUserHandler(userLogic logic.UserLogicInterface) *userHandler {
	return &userHandler{userLogic: userLogic}
}

func (h *userHandler) RegistrateUser(w http.ResponseWriter, r *http.Request) {
	err := h.userLogic.CreateUser(r.Body)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	logger.LogMessage("User seccessfully registrated")

	w.WriteHeader(200)
}

func (h *userHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	token, err := h.userLogic.LoginUser(r.Body)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	logger.LogMessage("User successfully logged in")

	json.NewEncoder(w).Encode(map[string]string{"token": token})
}

func (h *userHandler) CheckToken(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")
	token := strings.TrimPrefix(authHeader, "Bearer ")

	user_id, err := h.userLogic.CheckToken(token)
	if err != nil {
		utils.ResponseErrorJson(err, w)
		return
	}

	logger.LogMessage("Token successfully checked")

	w.Header().Set("X-User-ID", user_id)
	w.WriteHeader(200)
}
