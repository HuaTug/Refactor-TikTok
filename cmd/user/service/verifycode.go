package service

import (
	"context"
	"errors"

	"HuaTug.com/cmd/user/infras/redis"
	"HuaTug.com/kitex_gen/users"
)

type VerifyCodeService struct {
	ctx context.Context
}

func NewVerifyCodeService(ctx context.Context) *VerifyCodeService {
	return &VerifyCodeService{ctx: ctx}
}

func (s *VerifyCodeService) VerifyCode(req *users.VerifyCodeRequest) error {
	rcv, err := redis.GetCode(req.Email)
	if err != nil {
		return err
	} else {
		if rcv != req.Code {
			return errors.New("验证码错误")
		}
		redis.DeleteKey(req.Email)
		return nil
	}
}
