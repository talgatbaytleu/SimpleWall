package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	Run()
}

func Run() {
	mux := handler.InitServer()

	fmt.Println("Server started on port: 8081")
	log.Fatal(http.ListenAndServe(":8081", middleware.RecoverMiddleware(mux)))
}
