package main

import (
	"github.com/syumai/workers"
	"github.com/syumai/workers/_templates/cloudflare/deploy-go/router"
)

func main() {
	// 设置路由
	router.SetupRoutes()

	// 启动服务
	workers.Serve(nil) // use http.DefaultServeMux
}
