package service

import "notifier/internal/dal"

type NotificationServiceInterface interface{}

type notificationService struct {
	notificationDal dal.NotificationDalInterface
}

func NewNotificationService(notificationDal dal.NotificationDalInterface) *notificationService {
	return &notificationService{notificationDal: notificationDal}
}
