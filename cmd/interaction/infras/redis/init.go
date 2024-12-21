package redis

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

var (
	redisDBCommentInfo *redis.Client
	redisDBVideoInfo   *redis.Client
)

func Load() {

	redisDBCommentInfo = redis.NewClient(&redis.Options{
		Addr: CommentInfo.Addr,
		DB:   CommentInfo.DB,
	})

	redisDBVideoInfo = redis.NewClient(&redis.Options{
		Addr: VideoInfo.Addr,
		DB:   VideoInfo.DB,
	})

	if _, err := redisDBCommentInfo.Ping().Result(); err != nil {
		hlog.Info("redisDBCommentInfo", err)
	}
	if _, err := redisDBVideoInfo.Ping().Result(); err != nil {
		hlog.Info("redisDBVideoInfo", err)
	}
}
