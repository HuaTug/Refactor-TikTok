package websocket

import (
	"context"

	jwt "HuaTug.com/pkg"
	"github.com/cloudwego/hertz/pkg/app"
)

func _wsAuth() []app.HandlerFunc {
	return append(make([]app.HandlerFunc, 0),
		tokenAuthFunc(),
	)
}

func tokenAuthFunc() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		if !jwt.IsAccessTokenAvailable(ctx, c) {
			c.AbortWithStatus(401)
			return
		}
		c.Next(ctx)
	}
}
