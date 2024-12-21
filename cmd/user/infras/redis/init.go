package redis

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var redisDB *redis.Client

func Init() {
	redisDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		DB:   1,
	})

	pong, err := redisDB.Ping().Result()
	if err != nil {
		hlog.Info("Could not connect to redis : ", err)
	}
	hlog.Info("Connected to redis : ", pong)
}
