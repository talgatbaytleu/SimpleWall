package main

import (
	"fmt"
	"log"
	"net/http"

	"poster/internal/dal"
	"poster/internal/middleware"
	"poster/internal/router"
)

func main() {
	dbURL := "postgres://tbaitleu:talgat9595@localhost:5432/sw_posts_db"

	dal.InitDB(dbURL)
	defer dal.CloseDB()

	mux := router.InitServer()

	fmt.Println("Server started on port: 8082")
	log.Fatal(http.ListenAndServe(":8082", middleware.RecoverMiddleware(mux)))
}
