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

// GetFavoriteVideo 获取用户收藏的视频列表
func GetFavoriteVideo(ctx context.Context, c *app.RequestContext) {
	var GetVideoFromFavorite GetVideoFromFavoriteParam
	var err error
	var v interface{}
	var UserId int64

	if err = c.Bind(&GetVideoFromFavorite); err != nil {
		hlog.Info(err)
		SendResponse(c, errno.ConvertErr(err), nil)
	}

	if v, err = jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	} else {
		UserId = utils.Transfer(v)
	}

	resp, err := rpc.GetVideoFromFavorite(ctx, &videos.GetVideoFromFavoriteRequest{
		UserId:     UserId,
		FavoriteId: GetVideoFromFavorite.FavoriteId,
		VideoId:    GetVideoFromFavorite.VideoId,
		PageNum:    GetVideoFromFavorite.PageNum,
		PageSize:   GetVideoFromFavorite.PageSize,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
