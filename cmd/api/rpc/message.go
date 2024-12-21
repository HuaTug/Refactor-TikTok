package rpc

import (
	"context"
	"time"

	"HuaTug.com/config/jaeger"
	"HuaTug.com/kitex_gen/messages"
	"HuaTug.com/kitex_gen/messages/messageservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var MessageClient messageservice.Client

func InitMessageRpc() {
	r, err := etcd.NewEtcdResolver([]string{"localhost:2379"})
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger.NewClientTracer().Init("Message")
	defer closer.Close()
	c, err := messageservice.NewClient(
		"Message",
		client.WithRPCTimeout(3*time.Second),
		client.WithConnectTimeout(50*time.Second),
		client.WithFailureRetry(retry.NewFailurePolicy()),
		client.WithResolver(r),
		client.WithSuite(suite),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Message"}),
	)
	if err != nil {
		panic(err)
	}
	MessageClient = c
}

func InsertMessage(ctx context.Context, req *messages.InsertMessageRequest) error {
	_, err := MessageClient.InsertMessage(ctx, req)
	if err != nil {
		return err
	}
	return nil
}

func PopMessage(ctx context.Context, req *messages.PopMessageRequest) (*messages.PopMessageResponseData, error) {
	resp, err := MessageClient.PopMessage(ctx, req)
	if err != nil {
		hlog.Info(err)
		return nil, err
	}
	return resp.Data, nil
}
