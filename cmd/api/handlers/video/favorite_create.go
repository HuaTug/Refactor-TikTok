package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/videos"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func CreateFavoriteVideo(ctx context.Context, c *app.RequestContext) {
	var CreateFavorite CreateFavoriteParam
	var err error
	var v interface{}
	var UserId int64

	if err = c.Bind(&CreateFavorite); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	if v, err = jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	} else {
		UserId = utils.Transfer(v)
	}

	resp, err := rpc.CreateFavorite(ctx, &videos.CreateFavoriteRequest{
		UserId:      UserId,
		Name:        CreateFavorite.Name,
		Description: CreateFavorite.Description,
		CoverUrl:    CreateFavorite.CoverUrl,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
