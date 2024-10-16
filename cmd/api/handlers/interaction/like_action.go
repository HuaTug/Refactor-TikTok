package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/interactions"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func LikeAction(ctx context.Context, c *app.RequestContext) {
	var err error
	var v interface{}
	var UserId int64
	var Like LikeParam
	if err = c.BindAndValidate(&Like); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	if v, err = jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	} else {
		UserId = utils.Transfer(v)
	}
	if err := c.Bind(&Like); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	resp, err := rpc.LikeAction(ctx, &interactions.LikeActionRequest{
		UserId:     UserId,
		VideoId:    Like.VideoId,
		CommentId:  Like.CommentId,
		ActionType: Like.ActionType,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
