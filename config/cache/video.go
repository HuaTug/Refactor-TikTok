package cache

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"HuaTug.com/cmd/model"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/go-redis/redis/v8"
	Redis "github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
)

var (
	redisDBVideoUpload *redis.Client
	redisDBVideoInfo   *redis.Client
)

func CacheSetAuthor(videoid, authorid int64) {
	key := strconv.FormatInt(videoid, 10)
	err := CacheHSet("video:"+key, key, authorid)
	if err != nil {
		logrus.Info("Set cache error: ", err)
	}
}

func CacheGetAuthor(videoid int64) (int64, error) {
	key := strconv.FormatInt(videoid, 10)
	data, err := CacheHGet("video:"+key, key)
	if err != nil {
		return 0, nil
	}
	var uid int64
	err = json.Unmarshal(data, &uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}

// ToDo :实现了批量插入操作 共有两种方法 我使用第二种方法
func Insert(videos []*model.Video) {
	/*	pipe := redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			DB:   0,
		}).Pipeline()
		for _, v := range videos {
			pipe.Set(fmt.Sprintf("VideoId:%d", v.VideoId), fmt.Sprintf("Content:%d", v.FavoriteCount), 24*3600*30)
		}
		cmds, err := pipe.Exec()
		if err != nil {
			fmt.Printf("Pipeline exec error: %s", err)
			return
		} else {
			for _, cmd := range cmds {
				if cmd.Err() != nil {
					logrus.Info("Err: ", cmd.Err())
					return
				}
			}
		}
		logrus.Info("Pipe执行完毕")*/
	client := Redis.NewClient(&Redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       1,  // use default DB
	})

	//不要使用for循环进行redis的插入操作，这样会降低效率，可以使用redis的批量插入优化
	data := make(map[string]interface{})
	datas := make(map[string]interface{})
	for _, v := range videos {
		data[fmt.Sprintf("key:%d", v.VideoId)] = 0
	}

	for _, v := range videos {
		datas[fmt.Sprintf("video:%d", v.VideoId)] = v.UserId
	}
	err := client.MSet(context.Background(), data).Err()
	if err != nil {
		fmt.Println(err)
	}
	//设置Video和Author的映射
	err = client.MSet(context.Background(), datas).Err()
	if err != nil {
		fmt.Println(err)
	}
}

func GetVideoDBKeys(ctx context.Context) ([]string, error) {
	keys, err := redisDBVideoUpload.Keys(ctx, `*`).Result()
	if err != nil {
		return nil, err
	}
	return keys, err
}

func DelVideoDBKeys(ctx context.Context, keys []string) error {
	pipe := redisDBVideoUpload.TxPipeline()
	for _, key := range keys {
		pipe.Del(ctx, key)
	}
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}
	return nil
}

func NewVideoEvent(ctx context.Context, title, description, uid, chunkTotalNumber string) (string, error) {
	if redisDBVideoUpload == nil {
		hlog.Info("这里出错了")
		return ``, errors.New("redisDBVideoUpload is not initialized")
	}

	uuid := fmt.Sprint(time.Now().Unix())
	key := "l:" + uid + ":" + uuid

	// Check if the key exists
	exist, err := redisDBVideoUpload.Exists(ctx, key).Result()
	if err != nil {
		log.Printf("Redis Exists error: %v", err) // Add logging
		return ``, err
	}

	// If the key exists, return an error
	if exist != 0 {
		return ``, fmt.Errorf("key %s already exists", key)
	}

	// Add the item to the list
	_, err = redisDBVideoUpload.RPush(ctx, key, chunkTotalNumber, title, description).Result()
	if err != nil {
		log.Printf("Redis RPush error: %v", err) // Add logging
		return ``, err
	}

	return uuid, nil
}

func DoneChunkEvent(ctx context.Context, uuid, uid string, chunk int64) error {
	bitrecord, err := redisDBVideoUpload.GetBit(ctx, "b:"+uuid, chunk).Result()
	if err != nil {
		return err
	}
	if bitrecord == 1 {
		return err
	}
	if _, err = redisDBVideoUpload.SetBit(ctx, "b:"+uid+":"+uuid, chunk, 1).Result(); err != nil {
		return err
	}
	return nil
}

func IsChunkAllRecorded(ctx context.Context, uuid, uid string) (bool, error) {
	r, err := redisDBVideoUpload.LRange(ctx, "l:"+uid+":"+uuid, 0, 0).Result()
	if err != nil {
		return false, nil
	}
	chunkTotalNumber, _ := strconv.ParseInt(r[0], 10, 64)
	recordNumber, err := redisDBVideoUpload.BitCount(ctx, "b:"+uid+":"+uuid, &redis.BitCount{
		Start: 0,
		End:   chunkTotalNumber - 1,
	}).Result()
	if err != nil {
		return false, err
	}
	return chunkTotalNumber == recordNumber, nil
}

func RecordM3U8Filename(ctx context.Context, uuid, uid, filename string) error {
	exist, err := redisDBVideoUpload.Exists(ctx, "l:"+uid+":"+uuid).Result()
	if err != nil {
		return err
	}
	if exist == 0 {
		return err
	}
	fLen, err := redisDBVideoUpload.LLen(ctx, "l:"+uid+":"+uuid).Result()
	if err != nil {
		return err
	}
	if fLen == 4 {
		return errno.RequestErr
	}
	if _, err := redisDBVideoUpload.RPush(ctx, "l:"+uid+":"+uuid, filename).Result(); err != nil {
		return err
	}
	return nil
}

func GetM3U8Filename(ctx context.Context, uuid, uid string) (string, error) {
	if filename, err := redisDBVideoUpload.LRange(ctx, "l:"+uid+":"+uuid, 3, 3).Result(); err != nil || filename[0] == `` {
		return ``, errno.RequestErr
	} else {
		return filename[0], nil
	}
}

func FinishVideoEvent(ctx context.Context, uuid, uid string) ([]string, error) {
	info, err := redisDBVideoUpload.LRange(ctx, "l:"+uid+":"+uuid, 1, 2).Result()
	if err != nil {
		return nil, err
	}
	return info, nil
}

func DeleteVideoEvent(ctx context.Context, uuid, uid string) error {
	hlog.Info("l:" + uid + ":" + uuid)
	pipe := redisDBVideoUpload.TxPipeline()
	pipe.Del(ctx, "l:"+uid+":"+uuid)
	pipe.Del(ctx, "b:"+uid+":"+uuid)
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}
	return nil
}
