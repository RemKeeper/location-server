# 项目结构说明

## 目录结构

```
location-server/
├── main.go                 # 主入口文件
├── router/                 # 路由层
│   └── router.go          # 路由配置
├── middleware/             # 中间件层
│   ├── middleware.go      # 中间件基础定义和Chain函数
│   ├── auth.go            # 认证中间件
│   ├── logger.go          # 日志中间件
│   └── cors.go            # CORS中间件
└── service/                # 服务层(业务逻辑)
    ├── location.go        # 位置服务
    └── admin.go           # 管理服务
```

## 架构说明

### 1. **main.go** - 主入口
- 负责初始化路由和启动服务
- 保持简洁，只调用路由设置

### 2. **router/** - 路由层
- 定义所有API路由
- 组合中间件和处理器
- 使用`Chain`函数链式组合中间件

### 3. **middleware/** - 中间件层
- `middleware.go`: 定义中间件类型和Chain函数
- `auth.go`: 认证中间件
- `logger.go`: 日志中间件
- `cors.go`: CORS跨域中间件

#### 中间件使用方式：
```go
type Middleware func(http.HandlerFunc) http.HandlerFunc

// 链式组合
http.HandleFunc("/api/admin", middleware.Chain(
    handler,
    middleware.AuthMiddleware,
    middleware.LoggerMiddleware,
))
```

### 4. **service/** - 服务层
- 包含具体的业务逻辑处理器
- 每个服务一个文件
- 使用结构体封装服务

## API路由

- `GET /` - 首页 (日志 + CORS)
- `GET /api/location` - 位置服务 (日志 + CORS)
- `GET /api/admin` - 管理面板 (认证 + 日志 + CORS)

## 如何添加新功能

### 添加新的中间件：
1. 在 `middleware/` 目录创建新文件
2. 定义函数签名为 `func(http.HandlerFunc) http.HandlerFunc`
3. 在路由中使用 `Chain` 组合

### 添加新的服务：
1. 在 `service/` 目录创建新文件
2. 定义服务结构体和方法
3. 在 `router/router.go` 中注册路由

### 添加新的路由：
在 `router/router.go` 的 `SetupRoutes()` 中添加：
```go
http.HandleFunc("/api/new", middleware.Chain(
    serviceHandler,
    middleware.SomeMiddleware,
))
```

## 运行项目

```bash
# 安装依赖
go mod download

# 运行项目
go run main.go
```

