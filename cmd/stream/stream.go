package main

import (
	"context"
	"net/http"

	"HuaTug.com/cmd/video/dal/db"
	"HuaTug.com/cmd/video/service"
	"HuaTug.com/pkg/oss"
	"github.com/cloudwego/hertz/pkg/common/hlog"
)

func init() {
	db.Init()
	oss.InitMinio()
}
func main() {
	ctx := context.Background()
	services := service.NewStreamVideoService(ctx)        // 创建服务实例
	http.HandleFunc("/video/stream", services.ServeVideo) // 使用服务的 ServeVideo 方法
	if err := http.ListenAndServe(":7000", nil); err != nil {
		hlog.Fatal("Failed to start server:", err)
	}
}
