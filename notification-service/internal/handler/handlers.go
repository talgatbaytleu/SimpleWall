package handler

import (
	"fmt"
	"net/http"

	"notifier/internal/service"
	"notifier/pkg/apperrors"
	"notifier/pkg/utils"
)

type NotificationHandlerInterface interface{}

type notificationHandler struct {
	notificationService service.NotificationServiceInterface
}

func NewNotificationHandler(
	notificationService service.NotificationServiceInterface,
) *notificationHandler {
	return &notificationHandler{notificationService: notificationService}
}

func (h *notificationHandler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	utils.ResponseErrorJson(apperrors.ErrNotFound, w)
}

func (h *notificationHandler) GetNotification(w http.ResponseWriter, r *http.Request) {
}
