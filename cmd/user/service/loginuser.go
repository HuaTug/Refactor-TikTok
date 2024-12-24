package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/cmd/user/infras/redis"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type LoginuserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginuserService {
	return &LoginuserService{ctx: ctx}
}

func (v *LoginuserService) LoginUser(req *users.LoginUserResquest) (*base.User, error, bool) {
	var user base.User
	var err error
	var flag bool
	if value, _ := redis.GetCode(req.Email); value != "" {
		return nil, errors.WithMessage(nil, "please input verify_code again!"), false
	}
	hlog.Info("username:", req.UserName, ", password:", req.Password,".email:",req.Email)
	if user, err, flag = db.CheckUser(v.ctx, req.UserName, req.Password); err != nil || !flag {
		return nil, errors.WithMessage(err, "dao.CheckUser failed"), false
	}

	return &user, nil, true
}
