package service

import (
	"context"

	"HuaTug.com/cmd/user/dal/db"
	"HuaTug.com/kitex_gen/base"
	"HuaTug.com/kitex_gen/users"
	"github.com/pkg/errors"
)

type QueryUserService struct {
	ctx context.Context
}

func NewQueryUserService(ctx context.Context) *QueryUserService {
	return &QueryUserService{ctx: ctx}
}

func (v *QueryUserService) QueryUserInfo(req *users.QueryUserRequest) (user []*base.User, total int64, err error) {
	var count int64
	if user, count, err = db.QueryUser(v.ctx, req.Keyword, req.Page, req.PageSize); err != nil {
		return user, count, errors.WithMessage(err, "QueryUserInfo failed")
	}
	return user, count, nil
}
