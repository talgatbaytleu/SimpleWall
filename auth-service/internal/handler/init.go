package handler

import (
	"net/http"

	"auth-service/internal/apperrors"
	"auth-service/internal/dal"
	"auth-service/internal/logic"
)

func InitServer() {
	userDal := dal.NewUserDal()
	userLogic := logic.NewUserLogic(userDal)
	userHandler := NewUserHandler(userLogic)

	http.HandleFunc("POST /registrate", userHandler.RegistrateUser)
	http.HandleFunc("POST /login", userHandler.LoginUser)
	http.HandleFunc("GET /validate", userHandler.CheckToken)
	http.HandleFunc("/", NotFoundHandler)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	apperrors.ResponseErrorJson(apperrors.ErrNotFound, w)
}
