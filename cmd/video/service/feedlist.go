package service

import (
	"context"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
)

type FeedListService struct {
	ctx context.Context
}

func NewFeedListService(ctx context.Context) *FeedListService {
	return &FeedListService{ctx: ctx}
}

// 这里的v指向方法，用于传递ctx上下文
// func (v *FeedListService) FeedList(req *videos.FeedServiceRequest) ([]*model.Video, error) {
// 	if video, err := db.Feedlist(v.ctx, req); err != nil {
// 		return video, errors.WithMessage(err, "dao.FeedList failed")
// 	} else {

// 		cache.Insert(video)
// 		for _, s := range video {
// 			err := cache.RangeAdd(0, s.UserId)
// 			if err != nil {
// 				hlog.Info(err)
// 			}
// 		}
// 		return video, nil
// 	}
// }

// FeedList 视频流接口
func (v *FeedListService) FeedList(req *videos.FeedServiceRequest) ([]*base.Video, error) {
	res, err := db.Feedlist(v.ctx, req)
	if err != nil {
		return nil, err
	}
	VideoFiles = res
	return res, nil
}
