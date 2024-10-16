package main

import (
	"context"
	"fmt"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/cmd/video/dal/redis"
	jwt "HuaTug.com/pkg"
	"HuaTug.com/pkg/errno"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/middlewares/server/recovery"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Init() {
	rpc.InitRPC()
	redis.Load()
}

func main() {
	Init()
	r := server.New(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(16*1024*1024*1024),
	)
	/* 	authMiddleware, _ := jwt.New(&jwt.HertzJWTMiddleware{
		Key:        []byte(constants.SecretKey),
		Timeout:    time.Hour,
		MaxRefresh: time.Hour,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					constants.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch e.(type) {
			case errno.ErrNo:
				return e.(errno.ErrNo).ErrMsg
			default:
				return e.Error()
			}
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(consts.StatusOK, map[string]interface{}{
				"code":   errno.SuccessCode,
				"token":  token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(code, map[string]interface{}{
				"code":    errno.AuthorizationFailedErrCode,
				"message": message,
			})
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var loginVar handlers.UserParam
			if err := c.Bind(&loginVar); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			if len(loginVar.UserName) == 0 || len(loginVar.PassWord) == 0 {
				return "", jwt.ErrMissingLoginValues
			}

			//return rpc.CheckUser(context.Background(), &users.CheckUserRequest{UserName: loginVar.UserName, Password: loginVar.PassWord})
			return 0, nil
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	}) */
	jwt.AccessTokenJwtInit()
	jwt.RefreshTokenJwtInit()
	r.Use(recovery.Recovery(recovery.WithRecoveryHandler(
		func(ctx context.Context, c *app.RequestContext, err interface{}, stack []byte) {
			hlog.SystemLogger().CtxErrorf(ctx, "[Recovery] err=%v\nstack=%s", err, stack)
			c.JSON(consts.StatusInternalServerError, map[string]interface{}{
				"code":    errno.ServiceErrCode,
				"message": fmt.Sprintf("[Recovery] err=%v\nstack=%s", err, stack),
			})
		})))

	/* 	v1 := r.Group("/v1")
	   	user1 := v1.Group("/user")
	   	user1.POST("/login", handlers.Login)
	   	user1.POST("/register", handlers.Register)
	   	user1.POST("/query",handlers.QueryUserInfo)
	   	user1.DELETE("/delete",handlers.DeleteUser) */
	/* 	r.NoRoute(func(ctx context.Context, c *app.RequestContext) {
	   		c.String(consts.StatusOK, "no route")
	   	})
	   	r.NoMethod(func(ctx context.Context, c *app.RequestContext) {
	   		c.String(consts.StatusOK, "no method")
	   	}) */
	register(r)

	r.Spin()

}
