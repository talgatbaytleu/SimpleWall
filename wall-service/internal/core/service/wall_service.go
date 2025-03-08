package service

import (
	"log"
	"strconv"

	"github.com/redis/go-redis/v9"

	"wall/internal/ports"
)

type wallService struct {
	psqlPort  ports.PsqlPort
	redisPort ports.RedisPort
}

func NewWallService(psqlPort ports.PsqlPort, redisPort ports.RedisPort) *wallService {
	return &wallService{psqlPort: psqlPort, redisPort: redisPort}
}

func (s *wallService) GetUserWall(user_idStr string) (string, error) {
	var jsonWall string

	jsonWall, err := s.redisPort.GetWall(user_idStr)
	if err != nil {
		if err == redis.Nil {
			user_id, err := strconv.Atoi(user_idStr)
			if err != nil {
				return "", err
			}

			jsonWall, err = s.psqlPort.SelectWall(user_id)
			if err != nil {
				return "", err
			}

			err = s.redisPort.SetWall(user_idStr, jsonWall)
			if err != nil {
				log.Println("redis-adapter: set: ", err)
			}

		} else {
			return "", err
		}
	}

	// if jsonWall == "" {
	// 	user_id, err := strconv.Atoi(user_idStr)
	// 	if err != nil {
	// 		return "", err
	// 	}
	//
	// 	jsonWall, err = s.psqlPort.SelectWall(user_id)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	fmt.Println(jsonWall)
	//
	// 	err = s.redisPort.SetWall(user_idStr, jsonWall)
	// 	if err != nil {
	// 		log.Println("redis-adapter: set: ", err)
	// 	}
	// 	fmt.Println(jsonWall)
	//
	// }

	return jsonWall, nil
}
