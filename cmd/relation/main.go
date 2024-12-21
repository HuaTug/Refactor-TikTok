package main

import (
	"net"

	"HuaTug.com/cmd/relation/dal"
	"HuaTug.com/cmd/relation/infras"
	"HuaTug.com/config"
	"HuaTug.com/config/cache"
	"HuaTug.com/config/jaeger"
	"HuaTug.com/config/pprof"
	relation "HuaTug.com/kitex_gen/relations/followservice"
	"HuaTug.com/pkg/bound"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/middleware"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	//tracer2.InitJaeger(constants.UserServiceName)
	infras.Init()
	dal.Init()
}

func main() {
	config.Init()
	pprof.Load()
	suite, closer := jaeger.NewServerSuite().Init("Relation")
	defer closer.Close()
	r, err := etcd.NewEtcdRegistry([]string{config.ConfigInfo.Etcd.Addr})
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":8892")
	if err != nil {
		panic(err)
	}
	Init()
	cache.Init()
	//当出现了UserServiceImpl报错时 说明当前该接口的方法没有被完全实现

	svr := relation.NewServer(new(RelationServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Relation"}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                              // middleware
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
