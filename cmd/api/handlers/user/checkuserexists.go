package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func CheckUserExistsById(ctx context.Context, c *app.RequestContext) {
	var userId int64
	if v, err := jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return 
	} else {
		userId = utils.Transfer(v)
	}
	hlog.Info(userId)
	resp := new(users.CheckUserExistsByIdResponse)
	var err error
	resp, err = rpc.CheckUserExistsById(ctx, &users.CheckUserExistsByIdRequst{
		UserId: userId,
	})
	if err != nil || !resp.Exists {
		SendResponse(c, errno.ConvertErr(err), resp)
		return
	}
	SendResponse(c, errno.Success, resp)
}
