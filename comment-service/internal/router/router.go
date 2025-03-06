package router

import (
	"net/http"

	"commenter/internal/handler"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	// postDal := dal.NewPostDal(dal.MainDB)
	// postService := service.NewPostService(postDal)
	// postHandler := handler.NewPostService(postService)
	//
	// mux.HandleFunc("POST /registrate", userHandler.RegistrateUser)
	// mux.HandleFunc("POST /login", userHandler.LoginUser)
	// mux.HandleFunc("GET /validate", userHandler.CheckToken)

	// mux.HandleFunc("POST /registrate", handler.HandleAuthService)
	// mux.HandleFunc("POST /login", handler.HandleAuthService)
	mux.HandleFunc("POST /comment/{post_id}", postHandler.PostPost)
	mux.HandleFunc("PUT /comment/{post_id}", postHandler.PutPost)
	mux.HandleFunc("GET /comment/{post_id}", postHandler.GetPost)
	mux.HandleFunc("DELETE /comment/{post_id}", postHandler.DeletePost)
	mux.HandleFunc("/", handler.NotFoundHandler)

	return mux
}
