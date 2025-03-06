package main

import (
	"fmt"
	"log"
	"net/http"

	"commenter/internal/dal"
	"commenter/internal/middleware"
	"commenter/internal/router"
)

func main() {
	dbURL := "postgres://tbaitleu:talgat9595@localhost:5432/sw_posts_db"

	dal.InitDB(dbURL)
	defer dal.CloseDB()

	mux := router.InitServer()

	fmt.Println("Server started on port: 8084")
	log.Fatal(http.ListenAndServe(":8084", middleware.RecoverMiddleware(mux)))
}
