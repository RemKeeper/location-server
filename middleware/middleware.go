package middleware

import "net/http"

// Middleware 中间件类型定义
type Middleware func(http.HandlerFunc) http.HandlerFunc

// Chain 链式组合中间件
func Chain(handler http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
