package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func VideoFeedList(ctx context.Context, c *app.RequestContext) {
	var VideoList VideoFeedListParam
	var err error
	if err = c.Bind(&VideoList); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	hlog.Info(VideoList.AuthorId)
	resp, err := rpc.VideoFeedList(ctx, &videos.VideoFeedListRequest{
		UserId:   VideoList.AuthorId,
		PageNum:  VideoList.PageNum,
		PageSize: VideoList.PageSize,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
