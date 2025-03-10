package router

import (
	"net/http"

	"notifier/internal/dal"
	"notifier/internal/handler"
	"notifier/internal/service"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	notificationDal := dal.NewNotificationDal(dal.MainDB)
	notificationService := service.NewNotificationService(notificationDal)
	notificationHandler := handler.NewNotificationHandler(notificationService)

	mux.HandleFunc("GET /notification", notificationHandler.GetNotification)
	mux.HandleFunc("/", notificationHandler.NotFoundHandler)

	return mux
}
