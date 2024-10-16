package service

import (
	"context"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type VideoSearchService struct {
	ctx context.Context
}

func NewVideoSearchService(ctx context.Context) *VideoSearchService {
	return &VideoSearchService{ctx: ctx}
}

func (v *VideoSearchService) VideoSearch(req *videos.VideoSearchRequest) (video []*base.Video, count int64, err error) {
	if video, count, err = db.Videosearch(v.ctx, req); err != nil {
		hlog.Info(err)
		return video, count, errors.WithMessage(err, "dao.VideoSearch failed")
	}
	return video, count, err
}
