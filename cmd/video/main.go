package main

import (
	"net"
	"time"

	video "HuaTug.com/kitex_gen/videos/videoservice"

	"HuaTug.com/cmd/video/common"
	"HuaTug.com/cmd/video/dal"
	"HuaTug.com/cmd/video/dal/redis"
	"HuaTug.com/pkg/bound"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/middleware"
	"HuaTug.com/pkg/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func Init() {
	//tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
	redis.Load()
	oss.InitMinio()
	common.NewVideoSync().Run()
}

func main() {
	Init()
	//r, err := etcd.NewEtcdRegistry([]string{config.ConfigInfo.Etcd.Addr})
	r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":8891")
	if err != nil {
		panic(err)
	}
	//

	//当出现了UserServiceImpl报错时 说明当前该接口的方法没有被完全实现

	//注意 这里的video等等方法在进行服务注册发现时 video此时是kitex生成下的一个service
	svr := video.NewServer(new(VideoServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "Video"}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                           // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),                    // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()), // BoundHandler
		server.WithRegistry(r),                              // registry
		server.WithMaxConnIdleTime(30*time.Second),
	)
	err = svr.Run()
	if err != nil {
		hlog.Info(err)
	}
}
