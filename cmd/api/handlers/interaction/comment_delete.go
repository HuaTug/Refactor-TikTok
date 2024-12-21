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

func DeleteComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var v interface{}
	var UserId int64
	var Comment DeleteCommentParam
	if err = c.BindAndValidate(&Comment); err != nil {
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
	if err := c.Bind(&Comment); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	resp, err := rpc.DeleteComment(ctx, &interactions.CommentDeleteRequest{
		VideoId:    Comment.VideoId,
		CommentId:  Comment.CommentId,
		FromUserId: UserId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
