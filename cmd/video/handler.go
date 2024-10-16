package main

import (
	"context"
	"fmt"

	"HuaTug.com/cmd/model"
	"HuaTug.com/cmd/video/service"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

type VideoServiceImpl struct{}

func (s *VideoServiceImpl) FeedService(ctx context.Context, req *videos.FeedServiceRequest) (resp *videos.FeedServiceResponse, err error) {
	resp = new(videos.FeedServiceResponse)
	resp.Base = &base.Status{}
	var video []*model.Video
	if video, err = service.NewFeedListService(ctx).FeedList(req); err != nil {
		hlog.CtxErrorf(ctx, "service.FeedService failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Get VideoFeed!"
		resp.VideoList = nil
		return resp, err
	}
	for _, videos := range video {
		baseVideo := &base.Video{
			VideoId:     videos.VideoId,
			UserId:      videos.UserId,
			VideoUrl:    videos.VideoUrl,
			CoverUrl:    videos.CoverUrl,
			Title:       videos.Title,
			Description: videos.Description,
			VisitCount:  videos.VisitCount,
			CreatedAt:   videos.CreatedAt,
			DeletedAt:   videos.DeletedAt,
			UpdatedAt:   videos.UpdatedAt,
		}
		resp.VideoList = append(resp.VideoList, baseVideo)
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Get VideoFeed Success"
	return resp, nil
}

func (s *VideoServiceImpl) VideoFeedList(ctx context.Context, req *videos.VideoFeedListRequest) (resp *videos.VideoFeedListResponse, err error) {
	resp = new(videos.VideoFeedListResponse)
	resp.Base = &base.Status{}
	var video []*base.Video
	var count int64
	if video, count, err = service.NewVideoListService(ctx).VideoList(req); err != nil {
		hlog.CtxErrorf(ctx, "service.VideoFeedList failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Get VideoList!"
		resp.VideoList = video
		return resp, err
	}
	//todo
	fmt.Print(count)

	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Get VideoList Success"
	resp.VideoList = video
	return resp, nil
}

func (s *VideoServiceImpl) VideoSearch(ctx context.Context, req *videos.VideoSearchRequest) (resp *videos.VideoSearchResponse, err error) {
	resp = new(videos.VideoSearchResponse)
	resp.Base = &base.Status{}
	var video []*base.Video
	var count int64

	if video, count, err = service.NewVideoSearchService(ctx).VideoSearch(req); err != nil {
		hlog.CtxErrorf(ctx, "service.VideoSearch failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Get VideoFeed!"
		resp.VideoSearch = video
		resp.Count = count

		return resp, err
	}

	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Get VideoFeed Success"
	resp.VideoSearch = video
	resp.Count = count

	return resp, nil
}

func (s *VideoServiceImpl) VideoPopular(ctx context.Context, req *videos.VideoPopularRequest) (resp *videos.VideoPopularResponse, err error) {
	resp = new(videos.VideoPopularResponse)
	resp.Base = &base.Status{}
	var video []*base.Video
	if video, err = service.NewVideoPopularService(ctx).VideoPopular(req); err != nil {
		hlog.CtxErrorf(ctx, "service.VideoPopular failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Get VideoFeed!"
		resp.Popular = video
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Get VideoFeed Success"
	resp.Popular = video
	return resp, nil
}

func (s *VideoServiceImpl) VideoPublishStart(ctx context.Context, req *videos.VideoPublishStartRequest) (resp *videos.VideoPublishStartResponse, err error) {
	resp = new(videos.VideoPublishStartResponse)
	resp.Base = &base.Status{}
	resp.Uuid = ``
	// TODO: Add your implementation logic here
	// Example:
	// err = service.NewVideoPublishStartService(ctx).StartPublishing(req)
	uuid, err := service.NewVideoUploadService(ctx).NewUploadEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.VideoPublishStart failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Start Video Publish!"
		resp.Uuid = ""
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Video Publish Started Successfully"
	resp.Uuid = uuid
	return resp, err
}

func (s *VideoServiceImpl) VideoPublishUploading(ctx context.Context, req *videos.VideoPublishUploadingRequest) (resp *videos.VideoPublishUploadingResponse, err error) {
	resp = new(videos.VideoPublishUploadingResponse)
	resp.Base = &base.Status{}
	// TODO: Add your implementation logic here
	// Example:
	// err = service.NewVideoPublishUploadingService(ctx).UploadVideo(req)
	err = service.NewVideoUploadService(ctx).NewUploadingEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.VideoPublishUploading failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Upload Video!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Video Uploaded Successfully"
	return resp, nil
}

func (s *VideoServiceImpl) VideoPublishComplete(ctx context.Context, req *videos.VideoPublishCompleteRequest) (resp *videos.VideoPublishCompleteResponse, err error) {
	resp = new(videos.VideoPublishCompleteResponse)
	resp.Base = &base.Status{}
	// TODO: Add your implementation logic here
	// Example:
	// err = service.NewVideoPublishCompleteService(ctx).CompletePublishing(req)
	err = service.NewVideoUploadService(ctx).NewUploadCompleteEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.VideoPublishComplete failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Complete Video Publish!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Video Publish Completed Successfully"
	return resp, nil
}

func (s *VideoServiceImpl) VideoPublishCancle(ctx context.Context, req *videos.VideoPublishCancleRequest) (resp *videos.VideoPublishCancleResponse, err error) {
	resp = new(videos.VideoPublishCancleResponse)
	resp.Base = &base.Status{}
	// TODO: Add your implementation logic here
	// Example:
	// err = service.NewVideoPublishCancelService(ctx).CancelPublishing(req)
	err = service.NewVideoUploadService(ctx).NewCancleUploadEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.VideoPublishCancle failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Cancel Video Publish!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Video Publish Canceled Successfully"
	return resp, nil
}

func (s *VideoServiceImpl) VideoVisit(ctx context.Context, req *videos.VideoVisitRequest) (resp *videos.VideoVisitResponse, err error) {
	resp = new(videos.VideoVisitResponse)
	resp.Base = &base.Status{}
	data := &base.Video{}
	// TODO: Add your implementation logic here
	// Example:
	// err = service.NewVideoVisitService(ctx).RecordVisit(req)
	data, err = service.NewVideoUploadService(ctx).NewVideoVisitEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.VideoVisit failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Record Video Visit!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Video Visit Recorded Successfully"
	resp.Item = data
	return resp, nil
}

func (s *VideoServiceImpl) GetVideoVisitCount(ctx context.Context, req *videos.GetVideoVisitCountRequest) (resp *videos.GetVideoVisitCountResponse, err error) {
	resp = new(videos.GetVideoVisitCountResponse)
	resp.Base = &base.Status{}
	resp.VisitCount, err = service.NewVideoUploadService(ctx).NewGetVisitCountEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.GetVideoVisitCount failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to GetVideoVisitCount!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "GetVideoVisitCount Successfully"
	return resp, nil
}

func (s *VideoServiceImpl) VideoDelete(ctx context.Context, req *videos.VideoDeleteRequest) (resp *videos.VideoDeleteResponse, err error) {
	resp = new(videos.VideoDeleteResponse)
	resp.Base = &base.Status{}
	// TODO: Add your implementation logic here
	// Example:
	err = service.NewVideoUploadService(ctx).NewDeleteEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.VideoDelete failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Delete Video!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Video Deleted Successfully"
	return resp, nil
}

func (s *VideoServiceImpl) VideoIdList(ctx context.Context, req *videos.VideoIdListRequest) (resp *videos.VideoIdListResponse, err error) {
	resp = new(videos.VideoIdListResponse)
	resp.Base = &base.Status{}
	isEnd, list, err := service.NewVideoUploadService(ctx).NewIdListEvent(req)
	if err != nil {
		resp.Base = &base.Status{
			Code: errno.ServiceErrCode,
			Msg:  "Failed get videolist by videoId",
		}
		hlog.Info(err)
		return resp, err
	}
	resp.Base = &base.Status{
		Code: 200,
		Msg:  "Success get videolist by videoId",
	}
	resp.IsEnd = isEnd
	resp.List = *list
	return resp, nil
}

func (s *VideoServiceImpl) VideoInfo(ctx context.Context, req *videos.VideoInfoRequest) (resp *videos.VideoInfoResponse, err error) {
	resp = new(videos.VideoInfoResponse)
	resp.Base = &base.Status{}
	data := new(base.Video)
	data, err = service.NewVideoListService(ctx).VideoInfo(req)
	if err != nil {
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = err.Error()
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Get Video Info Success"
	resp.Items = data
	return resp, nil
}

func (s *VideoServiceImpl) UpdateVisitCount(ctx context.Context, req *videos.UpdateVisitCountRequest) (resp *videos.UpdateVisitCountResponse, err error) {
	resp = new(videos.UpdateVisitCountResponse)
	resp.Base = &base.Status{}
	// TODO: Add your implementation logic here
	// Example:
	err = service.NewVideoUploadService(ctx).NewUpdateVideoVisitCountEvent(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.UpdateVisitCount failed, original error: %v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Update Visit Count!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Update Visit Count Success"
	return resp, nil
}

func (s *VideoServiceImpl) GetVideoVisitCountInRedis(ctx context.Context, req *videos.GetVideoVisitCountInRedisRequest) (resp *videos.GetVideoVisitCountInRedisResponse, err error) {
	resp = new(videos.GetVideoVisitCountInRedisResponse)
	resp.Base = &base.Status{}
	data, err := service.NewVideoUploadService(ctx).NewGetVisitCountInRedisEvent(req)
	if err != nil {
		resp.Base.Code = errno.ServiceErrCode
		resp.Base.Msg = "Failed to get Videovisit_Count"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Success to get Videovisit_Count from Redis"
	resp.VisitCount = data
	return resp, nil
}
