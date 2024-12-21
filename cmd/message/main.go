package main

import (
	"net"

	"HuaTug.com/cmd/message/dal"
	"HuaTug.com/config"
	"HuaTug.com/config/jaeger"
	message "HuaTug.com/kitex_gen/messages/messageservice"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	dal.Load()
}

func main() {
	config.Init()
	Init()
	r, err := etcd.NewEtcdRegistry([]string{config.ConfigInfo.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8893")
	if err != nil {
		panic(err)
	}

	suite, closer := jaeger.NewServerSuite().Init("Message")
	defer closer.Close()

	svr := message.NewServer(new(MessageServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithSuite(suite),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Message"}),
	)
	err = svr.Run()
	if err != nil {
		hlog.Fatal(err)
	}
}
