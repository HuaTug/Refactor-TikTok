package rpc

import (
	"context"
	"time"

	"HuaTug.com/config/jaeger"
	"HuaTug.com/kitex_gen/relations"
	"HuaTug.com/kitex_gen/relations/followservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var relationClient followservice.Client

func InitRealtionRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		hlog.Info(err)
	}
	suite, closer := jaeger.NewClientTracer().Init("Relation")
	defer closer.Close()
	c, err := followservice.NewClient(
		"Relation",
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
	relationClient = c
}

func Relation(ctx context.Context, req *relations.RelationServiceRequest) (resp *relations.RelationServiceResponse, err error) {
	resp = new(relations.RelationServiceResponse)
	resp, err = relationClient.RelationService(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func RelationService(ctx context.Context, req *relations.RelationServiceRequest) (resp *relations.RelationServiceResponse, err error) {
	resp = new(relations.RelationServiceResponse)
	resp, err = relationClient.RelationService(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func FollowingList(ctx context.Context, req *relations.FollowingListRequest) (resp *relations.FollowingListResponse, err error) {
	resp = new(relations.FollowingListResponse)
	resp, err = relationClient.FollowingList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func FollowerList(ctx context.Context, req *relations.FollowerListRequest) (resp *relations.FollowerListResponse, err error) {
	resp = new(relations.FollowerListResponse)
	resp, err = relationClient.FollowerList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func FriendList(ctx context.Context, req *relations.FriendListRequest) (resp *relations.FriendListResponse, err error) {
	resp = new(relations.FriendListResponse)
	resp, err = relationClient.FriendList(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
