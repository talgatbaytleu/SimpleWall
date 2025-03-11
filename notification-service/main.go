package main

import (
	"context"
	"log"
	"net/http"

	"notifier/internal/dal"
	"notifier/internal/middleware"
	"notifier/internal/router"
	"notifier/internal/service"
	"notifier/pkg/logger"
)

func main() {
	dal.InitDB()
	defer dal.CloseDB()

	mux := router.InitServer()

	consumerKafka := service.NewConsumeKafkaService(dal.MainDB)
	go consumerKafka.ConsumeKafkaMessages(context.Background())

	logger.LogMessage("Server started on port: 8080")
	log.Fatal(http.ListenAndServe(":8080", middleware.RecoverMiddleware(mux)))
}
