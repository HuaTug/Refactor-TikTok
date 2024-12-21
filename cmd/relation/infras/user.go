package infras

import (
	"time"

	"HuaTug.com/config"
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
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Relation"}),
	)
	if err != nil {
		hlog.Info(err)
	}
	hlog.Info("Success init rpc")
	UserClient = c
}
