package main

import (
	"net"

	"HuaTug.com/cmd/interaction/dal"
	"HuaTug.com/cmd/interaction/infras/client"
	"HuaTug.com/cmd/interaction/infras/redis"
	"HuaTug.com/config/jaeger"

	"HuaTug.com/config"
	interaction "HuaTug.com/kitex_gen/interactions/interactionservice"
	"HuaTug.com/pkg/bound"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	//trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	//tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
	redis.Load()
	client.Init()
	// go common.NewCommentSync().Run()
	// go common.NewVideoSyncman().Run()
}

func main() {
	//pprof.Load()
	config.Init()
	Init()
	//cache.Init()

	suite, closer := jaeger.NewServerSuite().Init("Interaction")
	defer closer.Close()
	r, err := etcd.NewEtcdRegistry([]string{config.ConfigInfo.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":8893")
	if err != nil {
		panic(err)
	}

	//当出现了UserServiceImpl报错时 说明当前该接口的方法没有被完全实现

	svr := interaction.NewServer(new(InteractionServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Interaction"}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                                 // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithSuite(suite),                             // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()), // BoundHandler
		server.WithRegistry(r),                              // registry
	)
	err = svr.Run()
	if err != nil {
		hlog.Info(err)
	}
}
