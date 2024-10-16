package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/relations"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"HuaTug.com/pkg/utils"
	"github.com/cloudwego/hertz/pkg/app"
)

func FriendList(ctx context.Context, c *app.RequestContext) {
	var relationservice RelationPageParam
	var userId int64
	if err := c.Bind(&relationservice); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return 
	}
	if v, err := jwt.ConvertJWTPayloadToString(ctx, c); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return 
	} else {
		userId = utils.Transfer(v)
	}
	resp := new(relations.FriendListResponse)
	var err error
	resp, err = rpc.FriendList(ctx, &relations.FriendListRequest{
		PageNum:  relationservice.PageNum,
		PageSize: relationservice.PageSize,
		UserId:   userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), resp)
		return
	}
	SendResponse(c, errno.Success, resp)
}
