package main

import (
	"log"
	"net/http"

	"poster/internal/dal"
	"poster/internal/middleware"
	"poster/internal/router"
	"poster/pkg/logger"
)

func main() {
	dal.InitDB()
	defer dal.CloseDB()

	mux := router.InitServer()

	logger.LogMessage("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoverMiddleware(mux)))
}
