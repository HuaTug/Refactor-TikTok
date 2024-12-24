package rpc

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
		client.WithRPCTimeout(50*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Second),         // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		//client.WithSuite(trace.NewDefaultClientSuite()),   // tracer
		client.WithResolver(r), // resolver
		client.WithSuite(suite),
	)
	if err != nil {
		hlog.Info(err)
	}
	UserClient = c
}

func CreateUser(ctx context.Context, req *users.CreateUserRequest) error {
	resp, err := UserClient.CreateUser(ctx, req)
	if err != nil {
		return err
	}
	if resp.Base.Code != 0 {
		return err
	}
	return nil
}

func LoginUser(ctx context.Context, req *users.LoginUserResquest) (resp *users.LoginUserResponse, err error) {
	resp, err = UserClient.LoginUser(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func CheckUserExistsById(ctx context.Context, req *users.CheckUserExistsByIdRequst) (*users.CheckUserExistsByIdResponse, error) {
	resp, err := UserClient.CheckUserExistsById(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func QueryUser(ctx context.Context, req *users.QueryUserRequest) (resp *users.QueryUserResponse, err error) {
	resp, err = UserClient.QueryUser(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (resp *users.DeleteUserResponse, err error) {
	resp, err = UserClient.DeleteUser(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func GetUserInfo(ctx context.Context, req *users.GetUserInfoRequest) (resp *users.GetUserInfoResponse, err error) {
	resp = &users.GetUserInfoResponse{}
	resp, err = UserClient.GetUserInfo(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (resp *users.UpdateUserResponse, err error) {
	resp, err = UserClient.UpdateUser(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func VerifyCode(ctx context.Context, req *users.VerifyCodeRequest) (resp *users.VerifyCodeResponse, err error) {
	resp, err = UserClient.VerifyCode(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}

func SendCode(ctx context.Context, req *users.SendCodeRequest) (resp *users.SendCodeResponse, err error) {
	resp, err = UserClient.SendCode(ctx, req)
	if err != nil {
		return resp, err
	}
	return resp, nil
}
