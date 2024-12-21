package main

import (
	"context"
	"fmt"

	"HuaTug.com/cmd/api/rpc"
	"HuaTug.com/cmd/video/infras/redis"

	webs "HuaTug.com/cmd/api/router/websocket"
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
	//pprof.Load()
	r := server.New(
		server.WithHostPorts("0.0.0.0:8888"),
		server.WithHandleMethodNotAllowed(true),
		server.WithMaxRequestBodySize(16*1024*1024*1024),
	)
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
	register(r)

	ws := server.Default(
		server.WithHostPorts(`:10000`),
	)
	ws.NoHijackConnPool = true
	webs.WebsocketRegister(ws)

	go ws.Spin()
	r.Spin()

}
