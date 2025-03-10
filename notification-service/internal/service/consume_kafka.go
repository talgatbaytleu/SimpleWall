package service

import (
	"context"

	"github.com/segmentio/kafka-go"

	"notifier/internal/dal"
	"notifier/pkg/logger"
)

type ConsumeKafkaService struct {
	notificationDal dal.NotificationDalInterface
}

func NewConsumeKafkaService(notificationDal dal.NotificationDalInterface) *ConsumeKafkaService {
	return &ConsumeKafkaService{notificationDal: notificationDal}
}

func (s *ConsumeKafkaService) ConsumeKafkaMessages(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "comments-notifications",
		GroupID: "notification-service",
	})

	for {
		m, err := r.ReadMessage(ctx)
		if err != nil {
			logger.LogError(err)
			break
		}

		logger.LogMessage(string(m.Key) + ": " + string(m.Value))

		// Save notification to DB
	}
}
