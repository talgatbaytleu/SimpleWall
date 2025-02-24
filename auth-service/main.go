package main

import (
	"fmt"
	"log"
	"net/http"

	"auth-service/internal/dal"
)

func main() {
	Run()
}

func Run() {
	dbURL := "postgres://tbaitleu:talgat9595@localhost:5432/sw-users"

	dal.InitializeDB(dbURL)
	defer dal.CloseDB()

	// mux := handler.SetupServer()

	fmt.Println("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
