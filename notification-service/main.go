package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"notifier/internal/dal"
	"notifier/internal/middleware"
	"notifier/internal/router"
	"notifier/internal/service"
)

func main() {
	dbURL := "postgres://tbaitleu:talgat9595@localhost:5432/sw_posts_db"

	dal.InitDB(dbURL)
	defer dal.CloseDB()

	mux := router.InitServer()

	consumerKafka := service.NewConsumeKafkaService(dal.MainDB)
	go consumerKafka.ConsumeKafkaMessages(context.Background())

	fmt.Println("Server started on port: 8087")
	log.Fatal(http.ListenAndServe(":8087", middleware.RecoverMiddleware(mux)))
}
