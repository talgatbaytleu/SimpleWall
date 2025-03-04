package router

import (
	"net/http"

	"poster/internal/dal"
	"poster/internal/handler"
	"poster/internal/service"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	postDal := dal.NewPostDal(dal.MainDB)
	postService := service.NewPostService(postDal)
	postHandler := handler.NewPostService(postService)
	//
	// mux.HandleFunc("POST /registrate", userHandler.RegistrateUser)
	// mux.HandleFunc("POST /login", userHandler.LoginUser)
	// mux.HandleFunc("GET /validate", userHandler.CheckToken)

	// mux.HandleFunc("POST /registrate", handler.HandleAuthService)
	// mux.HandleFunc("POST /login", handler.HandleAuthService)
	mux.HandleFunc("POST /post", postHandler.PostPost)
	mux.HandleFunc("PUT /post/{post_id}", postHandler.PutPost)
	mux.HandleFunc("GET /post/{post_id}", postHandler.GetPost)
	mux.HandleFunc("DELETE /post/{post_id}", postHandler.DeletePost)
	mux.HandleFunc("/", handler.NotFoundHandler)

	return mux
}
