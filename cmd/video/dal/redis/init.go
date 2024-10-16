package redis

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var (
	redisDBVideoUpload *redis.Client
	redisDBVideoInfo   *redis.Client
)

func Load() {

	redisDBVideoUpload = redis.NewClient(&redis.Options{
		Addr: VideoUpload.Addr,
		DB:   VideoUpload.DB,
	})

	redisDBVideoInfo = redis.NewClient(&redis.Options{
		Addr: VideoInfo.Addr,
		DB:   VideoInfo.DB,
	})
 
	if _, err := redisDBVideoUpload.Ping().Result(); err != nil {
		hlog.Info("redisDBVideoUpload", err)
	}
	if _, err := redisDBVideoInfo.Ping().Result(); err != nil {
		hlog.Info("redisDBVideoInfo", err)
	}
}
