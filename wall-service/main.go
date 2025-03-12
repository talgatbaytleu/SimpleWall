package main

import (
	"log"
	"net/http"

	psqladapter "wall/internal/adapters/driven-adapters/psql-adapter"
	redisadapter "wall/internal/adapters/driven-adapters/redis-adapter"
	"wall/internal/middleware"
	"wall/internal/router"
	"wall/pkg/logger"
)

func main() {
	psqladapter.InitDB()
	defer psqladapter.CloseDB()

	redisadapter.InitRedis()

	mux := router.InitServer()

	logger.LogMessage("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoverMiddleware(mux)))
}
