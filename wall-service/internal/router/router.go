package router

import (
	"net/http"

	psqladapter "wall/internal/adapters/driven-adapters/psql-adapter"
	redisadapter "wall/internal/adapters/driven-adapters/redis-adapter"
	gatewayadapter "wall/internal/adapters/driver-adapters/gateway-adapter"
	"wall/internal/core/service"
)

func InitServer() *http.ServeMux {
	mux := http.NewServeMux()

	psqlAdapter := psqladapter.NewPsqlAdapter(psqladapter.MainDB)
	redisAdapter := redisadapter.NewRedisAdapter(redisadapter.RedisClient)
	wallService := service.NewWallService(psqlAdapter, redisAdapter)
	wallHandler := gatewayadapter.NewWallHandler(wallService)
	//
	mux.HandleFunc("GET /wall", wallHandler.GetUserWall)
	mux.HandleFunc("/", wallHandler.NotFoundHandler)

	return mux
}
