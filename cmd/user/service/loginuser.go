package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"github.com/pkg/errors"
)

type LoginuserService struct {
	ctx context.Context
}

func NewLoginUserService(ctx context.Context) *LoginuserService {
	return &LoginuserService{ctx: ctx}
}

func (v *LoginuserService) LoginUser(req *users.LoginUserResquest) (*base.User, error,bool) {
	var user base.User
	var err error
	var flag bool
	if user, err, flag = db.CheckUser(v.ctx, req.UserName, req.Password); err != nil || !flag {
		return &user, errors.WithMessage(err, "dao.CheckUser failed"),false
	}
	return &user, nil,true
}
