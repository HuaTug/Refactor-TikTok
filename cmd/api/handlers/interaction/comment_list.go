package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/interactions"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func ListComment(ctx context.Context, c *app.RequestContext) {
	var err error
	var Comment ListCommentParam
	if err = c.BindAndValidate(&Comment); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	if err := c.Bind(&Comment); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	resp, err := rpc.ListComment(ctx, &interactions.ListCommentRequest{
		VideoId:   Comment.VideoId,
		CommentId: Comment.CommentId,
		PageNum:   Comment.PageNum,
		PageSize:  Comment.PageSize,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
