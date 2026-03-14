package middleware

import (
	"github.com/syumai/workers/cloudflare"
	"log"
	"net/http"
)

const EnvDevKey = "DEV_KEY"

// AuthMiddleware 认证中间件示例
func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 这里可以添加认证逻辑
		devKey := r.Header.Get(EnvDevKey)

		if devKey == "" {
			log.Println("No authorization devKey provided")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized: No devKey provided"))
			return
		} else {
			if cloudflare.Getenv(EnvDevKey) != devKey {
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized: Invalid devKey"))
				return
			}
		}

		log.Printf("Auth check passed for: %s", r.URL.Path)
		next(w, r)
	}
}
