package client

import (
	"context"
	"time"

	"HuaTug.com/config"
	"HuaTug.com/config/jaeger"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/kitex_gen/users/userservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var UserClient userservice.Client

func InitUserRpc() {
	config.Init()
	//调用文件的位置则是main函数的起始位置
	r, err := etcd.NewEtcdResolver([]string{config.ConfigInfo.Etcd.Addr})
	if err != nil {
		hlog.Info(err)
	}
	suite, closer := jaeger.NewClientTracer().Init("User")
	defer closer.Close()
	c, err := userservice.NewClient(
		"User",
		/* 		client.WithMiddleware(middleware.CommonMiddleware),
		   		client.WithInstanceMW(middleware.ClientMiddleware), */
		client.WithMuxConnection(3),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Second),         // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
		client.WithSuite(suite),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Video"}),
	)
	if err != nil {
		hlog.Info(err)
	}
	UserClient = c
}

func CheckUserExistsById(ctx context.Context, req *users.CheckUserExistsByIdRequst) (*users.CheckUserExistsByIdResponse, error) {
	resp, err := UserClient.CheckUserExistsById(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *users.GetUserInfoRequest) (*users.GetUserInfoResponse, error) {
	resp, err := UserClient.GetUserInfo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
