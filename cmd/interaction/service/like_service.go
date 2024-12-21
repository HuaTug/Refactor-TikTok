package service

import (
	"context"
	"strconv"
	"time"

	"HuaTug.com/cmd/interaction/dal/db"
	"HuaTug.com/cmd/interaction/infras/client"
	"HuaTug.com/cmd/interaction/infras/redis"
	"HuaTug.com/cmd/model"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/interactions"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type LikeActionService struct {
	ctx context.Context
}

func NewLikeActionService(ctx context.Context) *LikeActionService {
	return &LikeActionService{ctx: ctx}
}

func (service *LikeActionService) NewLikeActionEvent(ctx context.Context, req *interactions.LikeActionRequest) (err error) {
	//对视频的点赞和对评论的点赞分为两个情况进行处理
	if req.VideoId != 0 {
		switch req.ActionType {
		case "1":
			{
				like := &model.UserBehavior{
					UserId:       req.UserId,
					VideoId:      req.VideoId,
					BehaviorType: "like",
					BehaviorTime: time.Now().Format(constants.DataFormate),
				}
				if err := redis.AppendVideoLikeInfo(req.VideoId, req.UserId); err != nil {
					hlog.Info(err)
					return errno.RedisErr
				}
				go db.AddUserLikeBehavior(service.ctx, like)
			}
		case "2":
			{
				if err := redis.RemoveVideoLikeInfo(req.VideoId, req.UserId); err != nil {
					return errno.RedisErr
				}
				go db.DeleteUserLikeBehavior(service.ctx, req.UserId, req.VideoId, "like")
			}
		}
	} else if req.CommentId != 0 {
		switch req.ActionType {
		case "1":
			{
				if err := redis.AppendCommentLikeInfo(req.CommentId, req.UserId); err != nil {
					hlog.Info(err)
					return errno.RedisErr
				}
			}
		case "2":
			{
				if err := redis.RemoveCommentLikeInfo(req.CommentId, req.UserId); err != nil {
					return errno.RedisErr
				}
			}
		}
	} else {
		return errno.RedisErr
	}
	return nil
}

func (service *LikeActionService) NewLikeListEvent(ctx context.Context, req *interactions.LikeListRequest) (resp *interactions.LikeListResponse, err error) {
	resp = new(interactions.LikeListResponse)
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = constants.DefaultLimit
	}
	list, err := db.GetVideoLikeListByUserId(ctx, req.UserId, req.PageNum, req.PageSize)
	if err != nil {
		return nil, errno.ServiceErr
	}
	hlog.Info(list)
	res := make([]*base.Video, len(*list))
	temp := new(videos.VideoInfoResponse)
	for i, item := range *list {
		vid, _ := strconv.ParseInt(item, 10, 64)
		if temp, err = client.VideoInfo(service.ctx, &videos.VideoInfoRequest{VideoId: vid}); err != nil {
			return nil, errno.ElasticSearchErr
		}
		res[i] = temp.Items
	}
	resp.Items = res
	resp.Base = &base.Status{}
	return resp, nil
}

func (service *LikeActionService) NewCommentListEvent(ctx context.Context, req *interactions.ListCommentRequest) (resp *interactions.ListCommentResponse, err error) {
	if req.PageNum <= 0 {
		req.PageNum = 1
	}
	if req.PageSize <= 0 {
		req.PageSize = constants.DefaultLimit
	}

	var (
		res *[]*base.Comment
	)
	resp = new(interactions.ListCommentResponse)
	if req.VideoId != 0 {
		if res, err = NewCommentService(service.ctx).GetVideoComment(req); err != nil {
			return nil, err
		}
	} else if req.CommentId != 0 {
		if res, err = NewCommentService(service.ctx).GetCommentComment(req); err != nil {
			return nil, err
		}
	} else {
		return nil, errno.RequestErr
	}
	resp.Items = *res
	return resp, nil
}
