package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func SendCode(ctx context.Context, c *app.RequestContext) {
	var Send SendCodeParam
	if err := c.BindAndValidate(&Send); err != nil {
		SendResponse(c, err, nil)
		return
	}
	resp, err := rpc.SendCode(ctx, &users.SendCodeRequest{
		Email: Send.Email,
	})
	if err != nil {
		SendResponse(c, err, nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
