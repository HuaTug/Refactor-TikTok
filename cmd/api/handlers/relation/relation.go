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

func RelationService(ctx context.Context, c *app.RequestContext) {
	var relationservice RelationParam
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
	resp := new(relations.RelationServiceResponse)
	var err error
	resp, err = rpc.Relation(ctx, &relations.RelationServiceRequest{
		ActionType: relationservice.ActionType,
		ToUserId:   relationservice.ToUserId,
		FromUserId: userId,
	})
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), resp)
		return
	}
	SendResponse(c, errno.Success, resp)
}
