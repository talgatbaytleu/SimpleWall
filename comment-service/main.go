package main

import (
	"log"
	"net/http"

	"commenter/internal/dal"
	"commenter/internal/middleware"
	"commenter/internal/router"
	"commenter/pkg/logger"
)

func main() {
	dal.InitDB()
	defer dal.CloseDB()

	mux := router.InitServer()

	logger.LogMessage("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoverMiddleware(mux)))
}
