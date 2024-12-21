package service

import (
	"context"
	"fmt"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/base"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/pkg/errors"
)

type GetUserInfoService struct {
	ctx context.Context
}

func NewGetUserInfoService(ctx context.Context) *GetUserInfoService {
	return &GetUserInfoService{ctx: ctx}
}

func (v *GetUserInfoService) GetUserInfo(userId int64) (user *base.User, err error) {
	//users, err := cache.CacheGetUser(userId)
	if user, err = db.GetUser(v.ctx, fmt.Sprint(userId)); err != nil {
		hlog.Info(err)
		return user, errors.WithMessage(err, "dao.GetUserInfo failed")
	} else {
		//go cache.CacheSetUser(user)
		return user, nil
	}
}
