package authfunc

import (
	"context"

	handlers "HuaTug.com/cmd/api/handlers/interaction"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"

	"github.com/cloudwego/hertz/pkg/app"
)

func Auth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		DoubleTokenAuthFunc(),
	)
}

func DoubleTokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if !jwt.IsAccessTokenAvailable(ctx, c) {
			if !jwt.IsRefreshTokenAvailable(ctx, c) {
				handlers.SendResponse(c, errno.ConvertErr(errno.TokenInvailedErr), nil)
				return
			}
			//此时表示refresh-token并未过期 在生成一个新的access-token
			//resp:=new(Res)

			//ToDo
			jwt.GenerateAccessToken(ctx, c)

		}
		c.Next(ctx)
	}
}