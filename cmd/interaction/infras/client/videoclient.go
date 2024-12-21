package client

import (
	"context"
	"time"

	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/kitex_gen/videos/videoservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var VideoClient videoservice.Client

func InitVideoRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		hlog.Info(err)
	}

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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Interaction"}),
	)
	if err != nil {
		hlog.Info(err)
	}
	VideoClient = c
}

func VideoIdList(ctx context.Context, req *videos.VideoIdListRequest) (resp *videos.VideoIdListResponse, err error) {
	resp, err = VideoClient.VideoIdList(ctx, req)
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

func UpdateVisitCount(ctx context.Context, req *videos.UpdateVisitCountRequest) (resp *videos.UpdateVisitCountResponse, err error) {
	resp, err = VideoClient.UpdateVisitCount(ctx, req)
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

func VideoInfo(ctx context.Context, req *videos.VideoInfoRequest) (resp *videos.VideoInfoResponse, err error) {
	resp, err = VideoClient.VideoInfo(ctx, req)
	if err != nil {
		return resp, nil
	}
	return resp, err
}
