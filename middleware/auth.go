package middleware

import (
	"log"
	"net/http"
)

// AuthMiddleware 认证中间件示例
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 这里可以添加认证逻辑
		token := r.Header.Get("Authorization")

		if token == "" {
			log.Println("No authorization token provided")
			// 可以选择是否继续执行
			// http.Error(w, "Unauthorized", http.StatusUnauthorized)
			// return
		}

		log.Printf("Auth check passed for: %s", r.URL.Path)
		next(w, r)
	}
}
