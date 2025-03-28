package router

import (
	"net/http"

	"auth-service/internal/dal"
	"auth-service/internal/handler"
	"auth-service/internal/logic"
	"auth-service/pkg/apperrors"
	"auth-service/pkg/utils"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()
	userDal := dal.NewUserDal(dal.MainDB)
	userLogic := logic.NewUserLogic(userDal)
	userHandler := handler.NewUserHandler(userLogic)

	mux.HandleFunc("POST /registrate", userHandler.RegistrateUser)
	mux.HandleFunc("POST /login", userHandler.LoginUser)
	mux.HandleFunc("GET /validate", userHandler.CheckToken)
	mux.HandleFunc("/", NotFoundHandler)

	return mux
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}
