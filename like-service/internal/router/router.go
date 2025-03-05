package router

import (
	"net/http"

	"liker/internal/dal"
	"liker/internal/handler"
	"liker/internal/service"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	likeDal := dal.NewLikeDal(dal.MainDB)
	likeService := service.NewLikeService(likeDal)
	likeHandler := handler.NewLikeHandler(likeService)

	mux.HandleFunc("POST /like/{post_id}", likeHandler.PostLike)
	mux.HandleFunc("GET /like/{post_id}", likeHandler.GetLike)
	mux.HandleFunc("DELETE /like/{post_id}", likeHandler.DeleteLike)
	mux.HandleFunc("/", handler.NotFoundHandler)

	return mux
}
