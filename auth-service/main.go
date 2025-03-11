package main

import (
	"log"
	"net/http"

	"auth-service/internal/dal"
	"auth-service/internal/middleware"
	"auth-service/internal/router"
	"auth-service/pkg/logger"
)

func main() {
	Run()
}

func Run() {
	dal.InitDB()
	defer dal.CloseDB()

	mux := router.InitServer()

	logger.LogMessage("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoverMiddleware(mux)))
}
