package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis"
)

func GetVideoDBKeys() ([]string, error) {
	keys, err := redisDBVideoUpload.Keys(`*`).Result()
	if err != nil {
		return nil, err
	}
	return keys, err
}
func GetChunkInfo(uid, uuid string) ([]string, error) {
	v, err := redisDBVideoUpload.LRange("l:"+uid+":"+uuid, 0, -1).Result()
	if err != nil {
		hlog.Info(err)
		return nil, err
	}
	return v, nil
}

// 在执行删除时 使用管道实现批量操作 减少网络开销
func DelVideoDBKeys(keys []string) error {
	pipe := redisDBVideoUpload.TxPipeline()
	for _, key := range keys {
		pipe.Del(key)
	}
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}

func NewVideoEvent(ctx context.Context, title, description, uid, chuckTotalNumber, lable_name, category string) (string, error) {
	uuid := fmt.Sprint(time.Now().Unix())
	exist, err := redisDBVideoUpload.Exists("l:" + uid + ":" + uuid).Result()
	if err != nil {
		return ``, err
	}
	if exist != 0 {
		return ``, errno.UserNotExistErr
	}
	if _, err := redisDBVideoUpload.RPush("l:"+uid+":"+uuid, chuckTotalNumber, title, description, lable_name, category).Result(); err != nil {
		return ``, err
	}
	return uuid, nil
}

// 这段代码利用了位图 用来标记上传的视频切片是否成功
func DoneChunkEvent(ctx context.Context, uuid, uid string, chunk int64) error {
	bitrecord, err := redisDBVideoUpload.GetBit("b:"+uid+":"+uuid, chunk).Result()
	if err != nil {
		hlog.Error("Failed to get bit from Redis:", err) // 记录错误
		return err
	}
	if bitrecord == 1 {
		return errors.New("Information already exists")
	}
	if _, err = redisDBVideoUpload.SetBit("b:"+uid+":"+uuid, chunk, 1).Result(); err != nil {
		hlog.Info("SetBit Error:", err)
		return err
	}
	hlog.Info("SetBit Success")
	return nil
}

// 这段代码是用来检查所有视频分片是否都已经被记录
func IsChunkAllRecorded(ctx context.Context, uuid, uid string) (bool, error) {
	// 由于这段代码使用的是Result,即其会返回一个字符串切片
	r, err := redisDBVideoUpload.LRange("l:"+uid+":"+uuid, 0, 0).Result()
	if err != nil {
		return false, err
	}
	chunkTotalNumber, _ := strconv.ParseInt(r[0], 10, 64)
	recordNumber, err := redisDBVideoUpload.BitCount("b:"+uid+":"+uuid, &redis.BitCount{
		Start: 0,
		End:   chunkTotalNumber - 1,
	}).Result()
	if err != nil {
		return false, err
	}
	return chunkTotalNumber == recordNumber, nil
}

func RecordM3U8Filename(ctx context.Context, uuid, uid, filename string) error {
	exist, err := redisDBVideoUpload.Exists("l:" + uid + ":" + uuid).Result()
	if err != nil {
		hlog.Info("First err")
		return err
	}
	if exist == 0 {
		hlog.Info("Second err")
		return errno.RequestErr
	}
	fLen, err := redisDBVideoUpload.LLen("l:" + uid + ":" + uuid).Result()
	if err != nil {
		hlog.Info("Thrid err")
		return err
	}
	if fLen == 4 {
		hlog.Info(errors.New("判断长度出错"))
		return errno.RequestErr
	}
	if _, err := redisDBVideoUpload.RPush("l:"+uid+":"+uuid, filename).Result(); err != nil {
		hlog.Info("Fifth err")
		return err
	}
	return nil
}

func GetM3U8Filename(ctx context.Context, uuid, uid string) (string, error) {
	if filename, err := redisDBVideoUpload.LRange("l:"+uid+":"+uuid, 3, 3).Result(); err != nil || filename[0] == `` {
		return ``, errno.RequestErr
	} else {
		return filename[0], nil
	}
}

func FinishVideoEvent(ctx context.Context, uuid, uid string) ([]string, error) {
	//这是表示在完成视频分片合成后 将视频的相关信息取出
	info, err := redisDBVideoUpload.LRange("l:"+uid+":"+uuid, 1, 2).Result()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func DeleteVideoEvent(ctx context.Context, uuid, uid string) error {
	hlog.Info("l:" + uid + ":" + uuid)
	pipe := redisDBVideoUpload.TxPipeline()
	pipe.Del("l:" + uid + ":" + uuid)
	pipe.Del("b:" + uid + ":" + uuid)
	if _, err := pipe.Exec(); err != nil {
		return err
	}
	return nil
}
