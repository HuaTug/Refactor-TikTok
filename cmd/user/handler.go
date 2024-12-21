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
	"context"
	"fmt"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/cmd/user/service"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/pkg/errors"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// CreateUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) CreateUser(ctx context.Context, req *users.CreateUserRequest) (resp *users.CreateUserResponse, err error) {
	resp = new(users.CreateUserResponse)

	if len(req.UserName) == 0 || len(req.Password) == 0 {
		return resp, nil
	}

	err = service.NewCreateUserService(ctx).CreateUser(req)
	if err != nil {
		hlog.CtxErrorf(ctx, "service.CreateUser failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		return resp, err
	}

	return resp, nil
}

func (s *UserServiceImpl) UpdateUser(ctx context.Context, req *users.UpdateUserRequest) (resp *users.UpdateUserResponse, err error) {
	var user *base.User
	resp = new(users.UpdateUserResponse)
	resp.Base = &base.Status{}
	if err := service.NewUpdateUserService(ctx).UpdateUser(req); err != nil {
		hlog.CtxErrorf(ctx, "service.UpdateUser failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail To Update User"
		resp.Data = nil
		return resp, err
	}
	if user, err = db.GetUser(ctx, fmt.Sprint(req.UserId)); err != nil {
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Get user info failed"
	}

	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Update User Success"
	resp.Data = user
	return resp, nil
}

func (s *UserServiceImpl) LoginUser(ctx context.Context, req *users.LoginUserResquest) (resp *users.LoginUserResponse, err error) {
	resp = new(users.LoginUserResponse)
	var user *base.User
	var flag bool
	user, err, flag = service.NewLoginUserService(ctx).LoginUser(req)
	if err != nil || !flag || user == nil {
		hlog.CtxErrorf(ctx, "service.LoginUser failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		return resp, err
	}
	resp.User = user
	return resp, nil
}

func (s *UserServiceImpl) CheckUserExistsById(ctx context.Context, req *users.CheckUserExistsByIdRequst) (resp *users.CheckUserExistsByIdResponse, err error) {
	resp = new(users.CheckUserExistsByIdResponse)
	resp.Base = &base.Status{}
	resp.Exists, err = service.NewCheckUserExistsService(ctx).CheckUserExists(req)
	if err != nil || !resp.Exists {
		hlog.CtxErrorf(ctx, "service.CheckUserExists failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "User not exists!"
		return resp, err
	}
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "User exists!"
	return resp, nil
}

func (s *UserServiceImpl) QueryUser(ctx context.Context, req *users.QueryUserRequest) (resp *users.QueryUserResponse, err error) {
	resp = new(users.QueryUserResponse)
	resp.Users, resp.Totoal, err = service.NewQueryUserService(ctx).QueryUserInfo(req)
	if err != nil || resp.Totoal == 0 {
		hlog.CtxErrorf(ctx, "service.QueryUser failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		return resp, err
	}
	for i := 0; i < int(resp.Totoal); i++ {
		resp.Users[i].Password = "xxxxxx"
	}
	//ToDo 需要先对其进行初始化 分配指针内存
	resp.Base = &base.Status{}
	resp.Base.Code = 200
	resp.Base.Msg = "Query Success"
	return resp, nil
}

func (s *UserServiceImpl) GetUserInfo(ctx context.Context, req *users.GetUserInfoRequest) (resp *users.GetUserInfoResponse, err error) {
	resp = new(users.GetUserInfoResponse)
	resp.Base = &base.Status{}
	var user *base.User
	if user, err = service.NewGetUserInfoService(ctx).GetUserInfo(req.UserId); err != nil {
		hlog.CtxErrorf(ctx, "service.GetUserInfo failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Code = consts.StatusBadRequest
		resp.Base.Msg = "Fail to Get User Info!"
		return resp, err
	}
	resp.User = user
	resp.Base.Code = consts.StatusOK
	resp.Base.Msg = "Get User Info Success!"
	return resp, nil
}

func (s *UserServiceImpl) DeleteUser(ctx context.Context, req *users.DeleteUserRequest) (resp *users.DeleteUserResponse, err error) {
	resp = new(users.DeleteUserResponse)
	resp.Base = &base.Status{}
	if err := service.NewDeleteUSerService(ctx).DeleteUser(req.UserId); err != nil {
		hlog.CtxErrorf(ctx, "service.DeleteUser failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Msg = "Fail to Delete User!"
		resp.Base.Code = 500
		return resp, err
	}
	resp.Base = &base.Status{}
	resp.Base.Code = 200
	resp.Base.Msg = "Delete User Success!"
	return resp, nil
}

func (s *UserServiceImpl) VerifyCode(ctx context.Context, req *users.VerifyCodeRequest) (resp *users.VerifyCodeResponse, err error) {
	resp = new(users.VerifyCodeResponse)
	resp.Base = &base.Status{}
	if err := service.NewVerifyCodeService(ctx).VerifyCode(req); err != nil {
		hlog.CtxErrorf(ctx, "service.VerifyCode failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Msg = "Fail to Verify Code!"
		resp.Base.Code = 500
		return resp, err
	}
	resp.Base = &base.Status{}
	resp.Base.Code = 200
	resp.Base.Msg = "Verify Code Success!"
	return resp, nil
}

func (s *UserServiceImpl) SendCode(ctx context.Context, req *users.SendCodeRequest) (resp *users.SendCodeResponse, err error) {
	resp = new(users.SendCodeResponse)
	resp.Base = &base.Status{}
	var code string
	if code, err = service.NewSendCodeService(ctx).SendCode(req); err != nil {
		hlog.CtxErrorf(ctx, "service.SendCode failed,original error:%v", errors.Cause(err))
		hlog.CtxErrorf(ctx, "stack trace: \n%+v\n", err)
		resp.Base.Msg = "Fail to Send Code!"
		resp.Base.Code = consts.StatusBadRequest
		return resp, err
	}
	resp.Base = &base.Status{}
	resp.Base.Code = 200
	resp.Base.Msg = code
	return resp, nil
}
