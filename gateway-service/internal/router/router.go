package router

import (
	"net/http"

	"gateway/internal/handler"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	// Auth-service routes
	mux.HandleFunc("POST /registrate", handler.HandleAuthService)
	mux.HandleFunc("POST /login", handler.HandleAuthService)
	// Post-service routes
	mux.HandleFunc("/post", handler.HandlePostService)
	mux.HandleFunc("/post/{pattern}", handler.HandlePostService)
	// Like-service routes
	mux.HandleFunc("/like", handler.HandleLikeService)
	mux.HandleFunc("/like/{pattern}", handler.HandleLikeService)
	mux.HandleFunc("/likes/{pattern}", handler.HandleLikeService)
	// Comment-service routes
	mux.HandleFunc("/comment", handler.HandleCommentService)
	mux.HandleFunc("/comment/{pattern}", handler.HandleCommentService)
	mux.HandleFunc("/comments", handler.HandleCommentService)
	// Wall-service routes
	mux.HandleFunc("/wall", handler.HandleWallService)

	// Not Allowed Routes
	mux.HandleFunc("/", handler.NotFoundHandler)

	return mux
}
