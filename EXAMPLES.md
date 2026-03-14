# 如何使用新架构示例

## 示例1: 添加新的中间件

在 `middleware/` 目录下创建 `ratelimit.go`:

```go
package middleware

import (
	"net/http"
	"time"
)

// RateLimitMiddleware 限流中间件
func RateLimitMiddleware(next http.HandlerFunc) http.HandlerFunc {
	lastRequest := time.Now()
	
	return func(w http.ResponseWriter, r *http.Request) {
		if time.Since(lastRequest) < time.Second {
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}
		lastRequest = time.Now()
		next(w, r)
	}
}
```

## 示例2: 添加新的服务

在 `service/` 目录下创建 `user.go`:

```go
package service

import (
	"encoding/json"
	"net/http"
)

// UserService 用户服务
type UserService struct{}

// NewUserService 创建用户服务实例
func NewUserService() *UserService {
	return &UserService{}
}

// GetUserInfo 获取用户信息
func (s *UserService) GetUserInfo(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"id":   1,
		"name": "John Doe",
		"email": "john@example.com",
	}
	
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateUser 创建用户
func (s *UserService) CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// 解析请求体
	var user map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	
	// 这里添加创建用户的逻辑
	response := map[string]interface{}{
		"message": "User created successfully",
		"user":    user,
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
```

## 示例3: 在路由中使用

在 `router/router.go` 的 `SetupRoutes()` 中添加：

```go
func SetupRoutes() {
	// ...existing code...
	
	// 初始化用户服务
	userService := service.NewUserService()
	
	// 用户信息路由 - 使用认证、日志和CORS中间件
	http.HandleFunc("/api/user", middleware.Chain(
		userService.GetUserInfo,
		middleware.AuthMiddleware,
		middleware.LoggerMiddleware,
		middleware.CORSMiddleware,
	))
	
	// 创建用户路由 - 使用限流、认证、日志和CORS中间件
	http.HandleFunc("/api/user/create", middleware.Chain(
		userService.CreateUser,
		middleware.RateLimitMiddleware,
		middleware.AuthMiddleware,
		middleware.LoggerMiddleware,
		middleware.CORSMiddleware,
	))
}
```

## 示例4: 自定义中间件组合

创建常用的中间件组合:

```go
// 在 middleware/middleware.go 中添加

// DefaultMiddlewares 默认中间件组合
var DefaultMiddlewares = []Middleware{
	LoggerMiddleware,
	CORSMiddleware,
}

// AuthMiddlewares 需要认证的中间件组合
var AuthMiddlewares = []Middleware{
	AuthMiddleware,
	LoggerMiddleware,
	CORSMiddleware,
}
```

使用:

```go
http.HandleFunc("/api/public", middleware.Chain(
	handler,
	middleware.DefaultMiddlewares...,
))

http.HandleFunc("/api/protected", middleware.Chain(
	handler,
	middleware.AuthMiddlewares...,
))
```

## 中间件执行顺序

使用 `Chain` 时，中间件从右到左执行（从后往前）:

```go
middleware.Chain(
	handler,
	A,  // 第三个执行
	B,  // 第二个执行
	C,  // 第一个执行
)
```

执行顺序: C -> B -> A -> handler

所以一般把最通用的中间件放在最后，最特定的放在最前面。

