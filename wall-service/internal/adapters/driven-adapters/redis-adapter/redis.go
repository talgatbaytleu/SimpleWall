package redisadapter

import (
	"time"

	"github.com/redis/go-redis/v9"
)

type redisAdapter struct {
	redisClient *redis.Client
}

func NewRedisAdapter(redisClient *redis.Client) *redisAdapter {
	return &redisAdapter{redisClient: redisClient}
}

func (a *redisAdapter) GetWall(user_id string) (string, error) {
	cacheKey := "wall:" + user_id

	jsonWall, err := a.redisClient.Get(ctx, cacheKey).Result()
	if err != nil {
		return "", err
	}

	return jsonWall, err
}

func (a *redisAdapter) SetWall(user_id string, jsonWall string) error {
	cacheKey := "wall:" + user_id

	err := a.redisClient.Set(ctx, cacheKey, jsonWall, 60*time.Second).Err()
	if err != nil {
		return err
	}

	return nil
}
