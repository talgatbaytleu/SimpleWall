package router

import (
	"net/http"

	"gateway/internal/handler"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()
	// userDal := dal.NewUserDal(dal.MainDB)
	// userLogic := logic.NewUserLogic(userDal)
	// userHandler := NewUserHandler(userLogic)
	//
	// mux.HandleFunc("POST /registrate", userHandler.RegistrateUser)
	// mux.HandleFunc("POST /login", userHandler.LoginUser)
	// mux.HandleFunc("GET /validate", userHandler.CheckToken)

	mux.HandleFunc("POST /registrate", handler.HandleAuthService)
	mux.HandleFunc("POST /login", handler.HandleAuthService)
	mux.HandleFunc("/post", handler.HandlePostService)
	mux.HandleFunc("/post/{pattern}", handler.HandlePostService)
	mux.HandleFunc("/like", handler.HandleLikeService)
	mux.HandleFunc("/comment/{pattern}", handler.HandleCommentService)
	mux.HandleFunc("/wall", handler.HandleWallService)
	mux.HandleFunc("/wall/{pattern}", handler.HandleWallService)
	mux.HandleFunc("/", handler.NotFoundHandler)

	return mux
}
