package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func VideoStream(ctx context.Context, c *app.RequestContext) {
	var Stream VideoStreamParam
	var err error
	if err = c.Bind(&Stream); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	resp, err := rpc.VideoStream(ctx, &videos.StreamVideoRequest{
		Index: Stream.Index,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
