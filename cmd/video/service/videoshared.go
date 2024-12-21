package service

import (
	"context"
	"fmt"
	"time"

	"HuaTug.com/cmd/model"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/cmd/video/infras/client"
	"HuaTug.com/cmd/video/infras/redis"

	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/constants"
	"github.com/pkg/errors"
)

type SharedVideoService struct {
	ctx context.Context
}

func NewSharedVideoService(ctx context.Context) *SharedVideoService {
	return &SharedVideoService{ctx: ctx}
}

func (s *SharedVideoService) SharedVideo(req *videos.SharedVideoRequest) error {

	userExistsCh := make(chan bool)
	videoExistsCh := make(chan bool)

	go func() {
		res := new(users.CheckUserExistsByIdResponse)
		res, _ = client.CheckUserExistsById(s.ctx, &users.CheckUserExistsByIdRequst{UserId: req.ToUserId})
		// 检查用户是否存在
		if !res.Exists {
			userExistsCh <- false
			return
		}
		userExistsCh <- true
	}()

	go func() {
		// 检查视频是否存在
		temp, err := db.GetVideo(s.ctx, req.VideoId)
		if err != nil {
			videoExistsCh <- false
			return
		}
		videoExistsCh <- (temp != nil)
	}()
	// 等待两个 goroutine 完成  用于阻塞操作
	userExists := <-userExistsCh
	videoExists := <-videoExistsCh

	if !userExists {
		return errors.New("user not exist")
	}

	if !videoExists {
		return errors.New("video not exist")
	}

	if err := db.SharedVideo(s.ctx, &model.VideoShare{
		UserId:    req.UserId,
		VideoId:   req.VideoId,
		ToUserId:  req.ToUserId,
		CreatedAt: time.Now().Format(constants.DataFormate),
	}); err != nil {
		return errors.WithMessage(err, "Failed to shared video")
	}
	share := &model.UserBehavior{
		UserId:       req.UserId,
		VideoId:      req.VideoId,
		BehaviorType: "share",
		BehaviorTime: time.Now().Format(constants.DataFormate),
	}
	go redis.IncrVideoShareInfo(fmt.Sprint(req.VideoId))
	go db.AddUserShareBehavior(s.ctx, share)
	return nil
}
