package common

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/cmd/video/infras/redis"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type SyncService struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func NewSyncSerivce() *SyncService {
	ctx, cancle := context.WithCancel(context.Background())
	return &SyncService{ctx: ctx, cancel: cancle}
}

type SyncData struct {
	vid        string
	list       *[]string
	VisitCount int64
	ShareCount int64
	err        error
}

func (s *SyncService) Run() {
	if err := SyncInit(); err != nil {
		hlog.Info(err)
		panic(err)
	}

	go func() {
		for {
			if err := s.SyncVideoData(); err != nil {
				hlog.Info(err)
			}
			time.Sleep(10 * time.Second)
			select {
			case <-s.ctx.Done():
				hlog.Info("Ok,停止同步[video]")
				return
			default:
			}
		}
	}()
}

func (s *SyncService) SyncVideoData() error {
	var (
		wg      sync.WaitGroup
		errChan = make(chan error, 1)
		err     error
		sync    SyncData
	)
	for i := 0; ; i++ {
		var vidList *[]string
		if vidList, err = db.GetVideoIdList(s.ctx, 1, 1000); err != nil {
			return err
		}

		for _, vid := range *vidList {
			videoId, _ := strconv.ParseInt(vid, 10, 64)
			wg.Add(2)
			go func(vid string) {
				defer wg.Done()
				sync.VisitCount, err = redis.GetVideoVisitCount(vid)
				if err != nil {
					hlog.Info(err)
					errChan <- err
				}
			}(vid)

			go func(vid string) {
				defer wg.Done()
				sync.ShareCount, err = redis.GetVideoShareCount(vid)
				if err != nil {
					hlog.Info(err)
					errChan <- err
				}
			}(vid)
			wg.Wait()

			select {
			case err = <-errChan:
				hlog.Info(err)
				return err
			default:
			}
			wg.Add(2)
			go func() {
				defer wg.Done()
				if err := db.UpdateVideoVisit(s.ctx, videoId, sync.VisitCount); err != nil {
					errChan <- err
				}
			}()

			go func() {
				defer wg.Done()
				if err := db.UpdateVideoShareCount(s.ctx, videoId, sync.ShareCount); err != nil {
					errChan <- err
				}
			}()
			wg.Wait()

			select {
			case err = <-errChan:
				hlog.Info(err)
				return err
			default:
			}

		}
	}
	return nil
}

func (s *SyncService) Stop() {
	s.cancel()
}

func SyncInit() error {
	var (
		list       *[]string
		visitCount int64
		shareCount int64
		err        error
	)
	if list, err = db.GetVideoIdList(context.Background(), 1, 100); err != nil {
		return err
	}

	for _, vid := range *list {
		if visitCount, err = db.GetVideoVisitCount(context.Background(), vid); err != nil {
			return err
		}
		if shareCount, err = db.GetVideoShareCount(context.Background(), vid); err != nil {
			return err
		}
		if err = redis.PutVideoInfo(vid, fmt.Sprint(visitCount), fmt.Sprint(shareCount)); err != nil {
			return err
		}
	}
	return nil
}
