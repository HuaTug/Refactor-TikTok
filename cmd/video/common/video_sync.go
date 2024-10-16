package common

import (
	"context"
	"fmt"
	"strconv"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/cmd/video/dal/redis"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type VideoSync struct {
	ctx    context.Context
	cancle context.CancelFunc
}

type VideoSyncData struct {
	VideoId    int64
	VisitCount int64
}

func NewVideoSync() *VideoSync {
	ctx, cancle := context.WithCancel(context.Background())
	return &VideoSync{
		ctx:    ctx,
		cancle: cancle,
	}
}

func (service *VideoSync) Run() {
	if err := VideoSyncInit(); err != nil {
		hlog.Info(err)
		panic(err)
	}
}

func (service *VideoSync) Stop() {
	service.cancle()
}

func VideoSyncInit() error {
	var (
		list *[]string
		err  error
	)

	//ToDo 等到数据量大的时候考虑处理进行性能优化
	for i := 0; ; i++ {
		list, err = db.GetVideoIdList(context.Background(), int64(i), 10)
		if err != nil {
			return err
		}
		if len(*list) == 0 {
			break
		}
		var (
			syncList = make([]VideoSyncData, 0)
			data     VideoSyncData
		)
		for _, v := range *list {
			videoId, _ := strconv.ParseInt(v, 10, 64)
			data.VideoId = videoId
			if data.VisitCount, err = db.GetVideoVisitCount(context.Background(), v); err != nil {
				return err
			}
			syncList = append(syncList, data)
		}
		if err := videoSyncDB2Redis(&syncList); err != nil {
			return err
		}
	}
	return nil
}

func videoSyncDB2Redis(syncList *[]VideoSyncData) error {
	for _, item := range *syncList {
		if err := redis.PutVideoVisitInfo(fmt.Sprint(item.VideoId), fmt.Sprint(item.VisitCount)); err != nil {
			return err
		}
	}
	return nil
}
