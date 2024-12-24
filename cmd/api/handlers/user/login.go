package handlers

import (
	"context"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/kitex_gen/users"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func LoginUser(ctx context.Context, c *app.RequestContext) {
	var loginVar LoginParam
	var err error
	if err := c.Bind(&loginVar); err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	resp := new(users.LoginUserResponse)
  hlog.Info(loginVar.UserName,",",loginVar.PassWord,",",loginVar.Email)
	resp, err = rpc.LoginUser(ctx, &users.LoginUserResquest{
		UserName: loginVar.UserName,
		Password: loginVar.PassWord,
		Email:    loginVar.Email,
	})
	jwt.AccessTokenJwtMiddleware.LoginHandler(ctx, c)
	jwt.RefreshTokenJwtMiddleware.LoginHandler(ctx, c)

	AccessToken := c.GetString("Access-Token")
	RefreshToken := c.GetString("Refresh-Token")
	// AccessToken := ""
	// RefreshToken := ""
	if resp.User != nil {
		resp.Base.Code = consts.StatusOK
		resp.Base.Msg = "Login Success"
		resp.Token = AccessToken
		resp.RefreshToken = RefreshToken
	}
	if err != nil {
		SendResponse(c, errno.ConvertErr(err), nil)
		return
	}
	SendResponse(c, errno.Success, resp)
}
