package router

import (
	"net/http"

	"commenter/internal/dal"
	"commenter/internal/handler"
	"commenter/internal/service"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	commentDal := dal.NewCommentDal(dal.MainDB)
	commentService := service.NewCommentService(commentDal)
	commentHandler := handler.NewCommentHandler(commentService)
	//
	mux.HandleFunc("POST /comment", commentHandler.PostComment)
	mux.HandleFunc("PUT /comment/{comment_id}", commentHandler.PutComment)
	mux.HandleFunc("GET /comment/{comment_id}", commentHandler.GetComment)
	mux.HandleFunc("GET /comments", commentHandler.GetComments)
	mux.HandleFunc("DELETE /comment/{comment_id}", commentHandler.DeleteComment)
	mux.HandleFunc("/", handler.NotFoundHandler)

	return mux
}
