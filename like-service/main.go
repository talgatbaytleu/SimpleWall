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
	dal.InitDB()
	defer dal.CloseDB()

	mux := router.InitServer()

	fmt.Println("Server started on port: 8083")
	log.Fatal(http.ListenAndServe(":8083", middleware.RecoverMiddleware(mux)))
}
