package service

import (
	"context"
	"strconv"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/config/cache"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type VideoPopularService struct {
	ctx context.Context
}

func NewVideoPopularService(ctx context.Context) *VideoPopularService {
	return &VideoPopularService{ctx: ctx}
}

func (v *VideoPopularService) VideoPopular(req *videos.VideoPopularRequest) (video []*base.Video, err error) {
	res, err := cache.RangeList("Rank")
	if err != nil {
		hlog.Info(err)
		return
	}
	var temp *base.Video
	for i := 0; i < len(res); i++ {
		s, err := strconv.Atoi(res[i])
		if err != nil {
			return video, errors.WithMessage(err, "Convert failed")
		}
		temp, err = db.FindVideo(v.ctx, int64(s))
		if err != nil {
			return nil, errors.WithMessage(err, "dao.FindVideo failed")
		}
		video = append(video, temp)
	}
	return video, nil
}
