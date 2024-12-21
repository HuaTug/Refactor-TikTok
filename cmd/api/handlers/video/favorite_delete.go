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

func DeleteFavorite(ctx context.Context, c *app.RequestContext) {
	var DeleteFav DeleteFavoriteParam
	var err error
	var v interface{}
	var UserId int64

	if err = c.Bind(&DeleteFav); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
	}
	if v, err = jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	} else {
		UserId = utils.Transfer(v)
	}

	resp, err := rpc.DeleteFavorite(ctx, &videos.DeleteFavoriteRequest{
		UserId:     UserId,
		FavoriteId: DeleteFav.FavoriteId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
