package dal

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type NotificationDalInterface interface{}

type notificationDal struct {
	DB *pgxpool.Pool
}

func NewNotificationDal(db *pgxpool.Pool) *notificationDal {
	return &notificationDal{DB: db}
}
