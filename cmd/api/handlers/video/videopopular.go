package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/videos"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
)

func VideoPopular(ctx context.Context,c *app.RequestContext){
	resp,err:=rpc.VideoPopular(ctx,&videos.VideoPopularRequest{})
	if err!=nil{
		SendResponse(c,errno.ConvertErr(err),nil)
		return
	}
	SendResponse(c,errno.Success,resp)
}