package service

import (
	"context"

	"HuaTug.com/cmd/user/infras/redis"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

type SendCodeService struct {
	ctx context.Context
}

func NewSendCodeService(ctx context.Context) *SendCodeService {
	return &SendCodeService{ctx: ctx}
}

func (s *SendCodeService) SendCode(req *users.SendCodeRequest) (string, error) {
  hlog.Info("from front req:" ,req.Email)
	if code, err := utils.SendEmail(req.Email); err != nil {
		hlog.Info(err)
		return "", err
	} else {
		redis.RecordCode(req.Email, code)
		return code, nil
	}
}
