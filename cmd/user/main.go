// Copyright 2021 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"net"

	"HuaTug.com/cmd/user/dal"
	"HuaTug.com/cmd/user/infras/redis"
	"HuaTug.com/config"
	"HuaTug.com/config/jaeger"
	user "HuaTug.com/kitex_gen/users/userservice"
	"HuaTug.com/pkg/bound"
	"HuaTug.com/pkg/constants"
	"HuaTug.com/pkg/middleware"
	"HuaTug.com/pkg/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	etcd "github.com/kitex-contrib/registry-etcd"
	//trace "github.com/kitex-contrib/tracer-opentracing"
)

func Init() {
	//tracer2.InitJaeger(constants.UserServiceName)
	dal.Init()
}

func main() {
	config.Init()
	//pprof.Load()
	if err := oss.InitMinio(); err != nil {
		hlog.Info(err)
		return
	}
	suite, closer := jaeger.NewServerSuite().Init("User")
	defer closer.Close()
	r, err := etcd.NewEtcdRegistry([]string{config.ConfigInfo.Etcd.Addr})
	//r, err := etcd.NewEtcdRegistry([]string{"localhost:2379"})
	if err != nil {
		panic(err)
	}
	ip, err := constants.GetOutBoundIP()
	if err != nil {
		panic(err)
	}
	addr, err := net.ResolveTCPAddr("tcp", ip+":8889")
	if err != nil {
		panic(err)
	}
	Init()
	redis.Init()
	//cache.Init()
	//当出现了UserServiceImpl报错时 说明当前该接口的方法没有被完全实现

	svr := user.NewServer(new(UserServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: "User"}), // server name
		server.WithMiddleware(middleware.CommonMiddleware),                          // middleware
		server.WithMiddleware(middleware.ServerMiddleware),
		server.WithServiceAddr(addr),                                       // address
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}), // limit
		server.WithMuxTransport(),                                          // Multiplex
		//server.WithSuite(trace.NewDefaultServerSuite()),
		server.WithSuite(suite),                             // tracer
		server.WithBoundHandler(bound.NewCpuLimitHandler()), // BoundHandler
		server.WithRegistry(r),                           // registry

	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}
