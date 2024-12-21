package rpc

import (
	"context"
	"fmt"
	"time"

	"HuaTug.com/config/jaeger"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/kitex_gen/videos/videoservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var VideoClient videoservice.Client

func InitVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		panic(err)
	}
	suite, closer := jaeger.NewClientTracer().Init("Video")
	defer closer.Close()
	c, err := videoservice.NewClient(
		"Video",
		/* 		client.WithMiddleware(middleware.CommonMiddleware),
		   		client.WithInstanceMW(middleware.ClientMiddleware), */
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(50*time.Second),         // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
		client.WithSuite(suite),
	)
	if err != nil {
		hlog.Info(err)
	}
	VideoClient = c
}

func FeedList(ctx context.Context, req *videos.FeedServiceRequest) (resp *videos.FeedServiceResponse, err error) {
	resp, err = VideoClient.FeedService(ctx, req)
	resp.Base = &base.Status{}
	if err != nil {
		if resp == nil {
			return
		}
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to provide FeedService!"
		return resp, err
	}
	return resp, nil
}

func VideoFeedList(ctx context.Context, req *videos.VideoFeedListRequest) (resp *videos.VideoFeedListResponse, err error) {
	resp, err = VideoClient.VideoFeedList(ctx, req)
	resp.Base = &base.Status{}
	if err != nil {
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to provide VideoFeedList Service!"
		return resp, err
	}
	return resp, err
}

func VideoSearch(ctx context.Context, req *videos.VideoSearchRequest) (resp *videos.VideoSearchResponse, err error) {
	resp, err = VideoClient.VideoSearch(ctx, req)
	resp.Base = &base.Status{}
	if err != nil {
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to provide VideoSearch Service!"
		return resp, err
	}
	return resp, err
}

func VideoPublishStart(ctx context.Context, req *videos.VideoPublishStartRequest) (resp *videos.VideoPublishStartResponse, err error) {
	resp = new(videos.VideoPublishStartResponse)
	resp, err = VideoClient.VideoPublishStart(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("failed to start video publish: %v", err)
	}

	if resp == nil {
		return nil, fmt.Errorf("received nil response from VideoPublishStart")
	}
	return resp, err
}

func VideoPublishUploading(ctx context.Context, req *videos.VideoPublishUploadingRequest) (resp *videos.VideoPublishUploadingResponse, err error) {
	resp = new(videos.VideoPublishUploadingResponse)
	resp, err = VideoClient.VideoPublishUploading(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func VideoPublishCancle(ctx context.Context, req *videos.VideoPublishCancleRequest) (resp *videos.VideoPublishCancleResponse, err error) {
	resp, err = VideoClient.VideoPublishCancle(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func VideoPublishComplete(ctx context.Context, req *videos.VideoPublishCompleteRequest) (resp *videos.VideoPublishCompleteResponse, err error) {
	resp, err = VideoClient.VideoPublishComplete(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func VideoVisit(ctx context.Context, req *videos.VideoVisitRequest) (resp *videos.VideoVisitResponse, err error) {
	resp, err = VideoClient.VideoVisit(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func VideoPopular(ctx context.Context, req *videos.VideoPopularRequest) (resp *videos.VideoPopularResponse, err error) {
	resp, err = VideoClient.VideoPopular(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func VideoDelete(ctx context.Context, req *videos.VideoDeleteRequest) (resp *videos.VideoDeleteResponse, err error) {
	resp, err = VideoClient.VideoDelete(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func VideoInfo(ctx context.Context, req *videos.VideoInfoRequest) (resp *videos.VideoInfoResponse, err error) {
	resp, err = VideoClient.VideoInfo(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}

func VideoIdList(ctx context.Context, req *videos.VideoIdListRequest) (resp *videos.VideoIdListResponse, err error) {
	resp, err = VideoClient.VideoIdList(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}

func UpdateVideoVisitCount(ctx context.Context, req *videos.UpdateVisitCountRequest) (resp *videos.UpdateVisitCountResponse, err error) {
	resp, err = VideoClient.UpdateVisitCount(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}

func UpdateVideoCommentCount(ctx context.Context, req *videos.UpdateVideoCommentCountRequest) (resp *videos.UpdateVideoCommentCountResponse, err error) {
	resp, err = VideoClient.UpdateVideoCommentCount(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}

func UpdateVideoLikeCount(ctx context.Context, req *videos.UpdateLikeCountRequest) (resp *videos.UpdateLikeCountResponse, err error) {
	resp, err = VideoClient.UpdateVideoLikeCount(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}

func UpdateVideoHisLikeCount(ctx context.Context, req *videos.UpdateVideoHisLikeCountRequest) (resp *videos.UpdateVideoHisLikeCountResponse, err error) {
	resp, err = VideoClient.UpdateVideoHisLikeCount(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}
func GetVideoVisitCount(ctx context.Context, req *videos.GetVideoVisitCountRequest) (resp *videos.GetVideoVisitCountResponse, err error) {
	resp, err = VideoClient.GetVideoVisitCount(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}

func GetVideoVisitCountInRedis(ctx context.Context, req *videos.GetVideoVisitCountInRedisRequest) (resp *videos.GetVideoVisitCountInRedisResponse, err error) {
	resp, err = VideoClient.GetVideoVisitCountInRedis(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}

func VideoStream(ctx context.Context, req *videos.StreamVideoRequest) (resp *videos.StreamVideoResponse, err error) {
	resp, err = VideoClient.StreamVideo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func CreateFavorite(ctx context.Context, req *videos.CreateFavoriteRequest) (resp *videos.CreateFavoriteResponse, err error) {
	resp, err = VideoClient.CreateFavorite(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func GetFavoroteList(ctx context.Context, req *videos.GetFavoriteListRequest) (resp *videos.GetFavoriteListResponse, err error) {
	resp, err = VideoClient.GetFavoriteList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func GetFavoriteVideoList(ctx context.Context, req *videos.GetFavoriteVideoListRequest) (resp *videos.GetFavoriteVideoListResponse, err error) {
	resp, err = VideoClient.GetFavoriteVideoList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func GetVideoFromFavorite(ctx context.Context, req *videos.GetVideoFromFavoriteRequest) (resp *videos.GetVideoFromFavoriteResponse, err error) {
	resp, err = VideoClient.GetVideoFromFavorite(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
func AddFavoriteVideo(ctx context.Context, req *videos.AddFavoriteVideoRequest) (resp *videos.AddFavoriteVideoResponse, err error) {
	resp, err = VideoClient.AddFavoriteVideo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func SharedVideo(ctx context.Context, req *videos.SharedVideoRequest) (resp *videos.SharedVideoResponse, err error) {
	resp, err = VideoClient.SharedVideo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func DeleteFavorite(ctx context.Context, req *videos.DeleteFavoriteRequest) (resp *videos.DeleteFavoriteResponse, err error) {
	resp, err = VideoClient.DeleteFavorite(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func DeleteVideoFromFavortie(ctx context.Context, req *videos.DeleteVideoFromFavoriteRequest) (resp *videos.DeleteVideoFromFavoriteResponse, err error) {
	resp, err = VideoClient.DeleteVideoFromFavorite(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}

func RecommendVideo(ctx context.Context, req *videos.RecommendVideoRequest) (resp *videos.RecommendVideoResponse, err error) {
	resp, err = VideoClient.RecommendVideo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, err
}
