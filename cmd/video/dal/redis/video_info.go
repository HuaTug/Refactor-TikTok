package redis

import (
	"strconv"

	"github.com/go-redis/redis"
)

func PutVideoVisitInfo(vid, visitCount string) error {
	score, _ := strconv.ParseFloat(visitCount, 64)
	_, err := redisDBVideoInfo.ZAdd(`visit`, redis.Z{Score: score, Member: vid}).Result()
	if err != nil {
		return err
	}
	return nil
}

func IncrVideoVisitInfo(vid string) error {
	_, err := redisDBVideoInfo.ZIncrBy(`visit`, 1, vid).Result()
	if err != nil {
		return err
	}
	return nil
}

func GetVideoVisitCount(vid string) (int64, error) {
	_, err := redisDBVideoInfo.ZRank(`visit`, vid).Result()
	if err != nil {
		return -1, err
	}
	s, err := redisDBVideoInfo.ZScore(`visit`, vid).Result()
	if err != nil {
		return -1, err
	}
	return int64(s), nil
}

func IsVideoExist(vid string) bool {
	_, err := redisDBVideoInfo.ZScore(`visit`, vid).Result()
	return err == nil
}

