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

	mux.HandleFunc("POST /like", likeHandler.PostLike)
	mux.HandleFunc("GET /likes/count", likeHandler.GetLikesCount)
	mux.HandleFunc("GET /likes", likeHandler.GetLikesList)
	mux.HandleFunc("DELETE /like", likeHandler.DeleteLike)
	mux.HandleFunc("/", handler.NotFoundHandler)

	return mux
}
