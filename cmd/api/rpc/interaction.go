package rpc

import (
	"context"
	"time"

	"HuaTug.com/config"
	"HuaTug.com/config/jaeger"
	"HuaTug.com/kitex_gen/interactions"
	"HuaTug.com/kitex_gen/interactions/interactionservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var InteractionClient interactionservice.Client

func InitInteractionRpc() {
	config.Init()
	//调用文件的位置则是main函数的起始位置
	r, err := etcd.NewEtcdResolver([]string{config.ConfigInfo.Etcd.Addr})
	if err != nil {
		hlog.Info(err)
	}
	suite, closer := jaeger.NewClientTracer().Init("Interaction")
	defer closer.Close()
	c, err := interactionservice.NewClient(
		"Interaction",
		/* 		client.WithMiddleware(middleware.CommonMiddleware),
		   		client.WithInstanceMW(middleware.ClientMiddleware), */
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Second),         // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
		client.WithSuite(suite),
	)
	if err != nil {
		hlog.Info(err)
	}
	InteractionClient = c
}

func CreateComment(ctx context.Context, req *interactions.CreateCommentRequest) (resp *interactions.CreateCommentResponse, err error) {
	resp, err = InteractionClient.CreateComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func ListComment(ctx context.Context, req *interactions.ListCommentRequest) (resp *interactions.ListCommentResponse, err error) {
	resp, err = InteractionClient.ListComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func DeleteComment(ctx context.Context, req *interactions.CommentDeleteRequest) (resp *interactions.CommentDeleteResponse, err error) {
	resp, err = InteractionClient.DeleteComment(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func LikeAction(ctx context.Context, req *interactions.LikeActionRequest) (resp *interactions.LikeActionResponse, err error) {
	resp, err = InteractionClient.LikeAction(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func LikeList(ctx context.Context, req *interactions.LikeListRequest) (resp *interactions.LikeListResponse, err error) {
	resp, err = InteractionClient.LikeList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
