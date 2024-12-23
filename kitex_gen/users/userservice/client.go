// Code generated by Kitex v0.10.3. DO NOT EDIT.

package userservice

import (
	users "HuaTug.com/kitex_gen/users"
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	UpdateUser(ctx context.Context, req *users.UpdateUserRequest, callOptions ...callopt.Option) (r *users.UpdateUserResponse, err error)
	DeleteUser(ctx context.Context, req *users.DeleteUserRequest, callOptions ...callopt.Option) (r *users.DeleteUserResponse, err error)
	QueryUser(ctx context.Context, req *users.QueryUserRequest, callOptions ...callopt.Option) (r *users.QueryUserResponse, err error)
	CreateUser(ctx context.Context, req *users.CreateUserRequest, callOptions ...callopt.Option) (r *users.CreateUserResponse, err error)
	LoginUser(ctx context.Context, req *users.LoginUserResquest, callOptions ...callopt.Option) (r *users.LoginUserResponse, err error)
	GetUserInfo(ctx context.Context, req *users.GetUserInfoRequest, callOptions ...callopt.Option) (r *users.GetUserInfoResponse, err error)
	CheckUserExistsById(ctx context.Context, req *users.CheckUserExistsByIdRequst, callOptions ...callopt.Option) (r *users.CheckUserExistsByIdResponse, err error)
	VerifyCode(ctx context.Context, req *users.VerifyCodeRequest, callOptions ...callopt.Option) (r *users.VerifyCodeResponse, err error)
	SendCode(ctx context.Context, req *users.SendCodeRequest, callOptions ...callopt.Option) (r *users.SendCodeResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfoForClient(), options...)
	if err != nil {
		return nil, err
	}
	return &kUserServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kUserServiceClient struct {
	*kClient
}

func (p *kUserServiceClient) UpdateUser(ctx context.Context, req *users.UpdateUserRequest, callOptions ...callopt.Option) (r *users.UpdateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.UpdateUser(ctx, req)
}

func (p *kUserServiceClient) DeleteUser(ctx context.Context, req *users.DeleteUserRequest, callOptions ...callopt.Option) (r *users.DeleteUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.DeleteUser(ctx, req)
}

func (p *kUserServiceClient) QueryUser(ctx context.Context, req *users.QueryUserRequest, callOptions ...callopt.Option) (r *users.QueryUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.QueryUser(ctx, req)
}

func (p *kUserServiceClient) CreateUser(ctx context.Context, req *users.CreateUserRequest, callOptions ...callopt.Option) (r *users.CreateUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CreateUser(ctx, req)
}

func (p *kUserServiceClient) LoginUser(ctx context.Context, req *users.LoginUserResquest, callOptions ...callopt.Option) (r *users.LoginUserResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.LoginUser(ctx, req)
}

func (p *kUserServiceClient) GetUserInfo(ctx context.Context, req *users.GetUserInfoRequest, callOptions ...callopt.Option) (r *users.GetUserInfoResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.GetUserInfo(ctx, req)
}

func (p *kUserServiceClient) CheckUserExistsById(ctx context.Context, req *users.CheckUserExistsByIdRequst, callOptions ...callopt.Option) (r *users.CheckUserExistsByIdResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CheckUserExistsById(ctx, req)
}

func (p *kUserServiceClient) VerifyCode(ctx context.Context, req *users.VerifyCodeRequest, callOptions ...callopt.Option) (r *users.VerifyCodeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.VerifyCode(ctx, req)
}

func (p *kUserServiceClient) SendCode(ctx context.Context, req *users.SendCodeRequest, callOptions ...callopt.Option) (r *users.SendCodeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendCode(ctx, req)
}
