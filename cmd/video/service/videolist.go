package service

import (
	"context"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"github.com/pkg/errors"
)

type VideoListService struct {
	ctx context.Context
}

func NewVideoListService(ctx context.Context) *VideoListService {
	return &VideoListService{ctx: ctx}
}

func (v *VideoListService) VideoList(req *videos.VideoFeedListRequest) (video []*base.Video, count int64, err error) {
	if video, count, err = db.Videolist(v.ctx, req); err != nil {
		return video, count, errors.WithMessage(err, "dao.VideoList failed")
	}
	return video, count, err
}

func (v *VideoListService) VideoInfo(req *videos.VideoInfoRequest) (data *base.Video, err error) {
	data, err = db.GetVideoInfo(v.ctx, req.VideoId)
	if err != nil {
		return nil, err
	}
	return data, nil
}
