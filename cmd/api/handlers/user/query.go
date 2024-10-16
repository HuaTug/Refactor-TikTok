package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func QueryUser(ctx context.Context, c *app.RequestContext) {
	var query QueryParam
	if err := c.Bind(&query); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}

	resp, err := rpc.QueryUser(ctx, &users.QueryUserRequest{
		PageSize: query.PageSize,
		Page:     query.PageNum,
		Keyword:  &query.Keyword,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
