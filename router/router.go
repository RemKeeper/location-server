package router

import (
	"net/http"

	"github.com/syumai/workers/_templates/cloudflare/deploy-go/middleware"
	"github.com/syumai/workers/_templates/cloudflare/deploy-go/service"
)

// SetupRoutes 设置所有路由
func SetupRoutes() {
	// 初始化服务
	locationService := service.NewLocationService()
	adminService := service.NewAdminService()

	// 首页路由 - 使用日志和CORS中间件
	http.HandleFunc("/", middleware.Chain(
		homeHandler,
		middleware.LoggerMiddleware,
		middleware.CORSMiddleware,
	))

	// 位置服务路由 - 使用日志和CORS中间件
	http.HandleFunc("/api/location", middleware.Chain(
		locationService.GetLocation,
		middleware.LoggerMiddleware,
		middleware.CORSMiddleware,
	))

	// 管理员路由 - 使用认证、日志和CORS中间件
	http.HandleFunc("/api/admin", middleware.Chain(
		adminService.HandleAdmin,
		middleware.AuthMiddleware,
		middleware.LoggerMiddleware,
		middleware.CORSMiddleware,
	))
}

// homeHandler 首页处理器
func homeHandler(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, Go! Location Server is running."
	if _, err := w.Write([]byte(msg)); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
