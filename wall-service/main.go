package main

import (
	"fmt"
	"log"
	"net/http"

	psqladapter "wall/internal/adapters/driven-adapters/psql-adapter"
	redisadapter "wall/internal/adapters/driven-adapters/redis-adapter"
	"wall/internal/middleware"
	"wall/internal/router"
)

func main() {
	dbURL := "postgres://tbaitleu:talgat9595@localhost:5432/sw_posts_db"

	psqladapter.InitDB(dbURL)
	defer psqladapter.CloseDB()

	redisadapter.InitRedis()

	mux := router.InitServer()

	fmt.Println("Server started on port: 8085")
	log.Fatal(http.ListenAndServe(":8085", middleware.RecoverMiddleware(mux)))
}
