package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func VideoSearch(ctx context.Context, c *app.RequestContext) {
	var Serach VideoSearchParam
	var err error
	if err = c.Bind(&Serach); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	resp, err := rpc.VideoSearch(ctx, &videos.VideoSearchRequest{
		Keyword:  Serach.Keyword,
		PageNum:  Serach.PageNum,
		PageSize: Serach.PageSize,
		FromDate: Serach.FromDate,
		ToDate:   Serach.ToDate,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
