package main

import (
	"log"
	"net/http"

	"gateway/internal/middleware"
	"gateway/internal/router"
	"gateway/pkg/logger"
)

func main() {
	Run()
}

func Run() {
	mux := router.InitServer()

	logger.LogMessage("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoverMiddleware(mux)))
}
