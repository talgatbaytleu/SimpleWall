package main

import (
	"fmt"
	"log"
	"net/http"

	"auth-service/internal/dal"
	"auth-service/internal/handler"
	"auth-service/internal/middleware"
)

func main() {
	Run()
}

func Run() {
	dbURL := "postgres://tbaitleu:talgat9595@localhost:5432/sw_users_auth"

	dal.InitializeDB(dbURL)
	defer dal.CloseDB()

	mux := handler.InitServer()

	fmt.Println("Server started on port: 8081")
	log.Fatal(http.ListenAndServe(":8081", middleware.RecoverMiddleware(mux)))
}
