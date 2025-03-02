package main

import (
	"fmt"
	"log"
	"net/http"

	"gateway/internal/middleware"
	"gateway/internal/router"
)

func main() {
	Run()
}

func Run() {
	mux := router.InitServer()

	fmt.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoverMiddleware(mux)))
}
