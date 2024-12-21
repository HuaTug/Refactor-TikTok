package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func VerifyCode(ctx context.Context, c *app.RequestContext) {
	var verify VerifyCodeParam
	if err := c.Bind(&verify); err != nil {
		SendResponse(c, errno.VerifyCodeErr, nil)
		return
	}
	resp, err := rpc.VerifyCode(ctx, &users.VerifyCodeRequest{
		Email: verify.Email,
		Code:  verify.Code,
	})
	if err != nil {
		SendResponse(c, errno.VerifyCodeErr, nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
