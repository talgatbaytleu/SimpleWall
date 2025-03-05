package main

import (
	"fmt"
	"log"
	"net/http"

	"liker/internal/dal"
	"liker/internal/middleware"
	"liker/internal/router"
)

func main() {
	dbURL := "postgres://tbaitleu:talgat9595@localhost:5432/sw_posts_db"

	dal.InitDB(dbURL)
	defer dal.CloseDB()

	mux := router.InitServer()

	fmt.Println("Server started on port: 8083")
	log.Fatal(http.ListenAndServe(":8083", middleware.RecoverMiddleware(mux)))
}
