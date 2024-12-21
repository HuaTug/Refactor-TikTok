package redis

import (
	"time"

	"github.com/cloudwego/hertz/pkg/common/hlog"
)

// 记录一个临时的验证码
func RecordCode(email, code string) error {
	expiration := 5 * time.Minute
	if err := redisDB.Set(email, code, expiration).Err(); err != nil {
		hlog.Info("Redis set key failed : ", err)
		return err
	}
	return nil
}

// 验证一个临时的验证码
func GetCode(email string) (string, error) {
	code, err := redisDB.Get(email).Result()
	if err != nil {
		hlog.Info("Redis get key failed : ", err)
		return "", err
	}
	return code, nil
}

func DeleteKey(key string) error {
	if err := redisDB.Del(key).Err(); err != nil {
		hlog.Info("Redis delete key failed : ", err)
		return err
	}
	return nil
}
