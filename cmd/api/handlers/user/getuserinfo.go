package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func GetUserInfo(ctx context.Context, c *app.RequestContext) {
	var userId int64
	if v, err := jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	} else {
		userId = utils.Transfer(v)
	}
	resp, err := rpc.GetUserInfo(ctx, &users.GetUserInfoRequest{
		UserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), resp)
		return
	}
	SendResponse(c, errno.Success, resp)
}
