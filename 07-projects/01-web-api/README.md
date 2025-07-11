# Go Web API 项目实战

## 项目概述
构建一个完整的RESTful API服务，实现用户管理和任务管理功能，帮助Java开发者理解Go Web开发的完整流程。

## 项目目标
- 掌握Go Web框架（Gin）的使用
- 理解Go项目结构和最佳实践
- 学会数据库操作（GORM）
- 掌握API设计和测试
- 对比Java Spring Boot开发模式

## 技术栈对比

| 功能 | Java (Spring Boot) | Go | 说明 |
|------|-------------------|----|----- |
| Web框架 | Spring MVC | Gin | Go更轻量 |
| ORM | JPA/Hibernate | GORM | 语法相似 |
| 依赖注入 | @Autowired | 手动管理 | Go更显式 |
| 配置管理 | application.yml | Viper | 类似功能 |
| 数据验证 | Bean Validation | validator | 标签驱动 |
| 中间件 | Filter/Interceptor | Gin Middleware | 概念相同 |

## 项目结构

```
web-api/
├── cmd/
│   └── server/
│       └── main.go              # 应用入口
├── internal/
│   ├── config/
│   │   └── config.go            # 配置管理
│   ├── database/
│   │   └── database.go          # 数据库连接
│   ├── handlers/
│   │   ├── user.go              # 用户处理器
│   │   └── task.go              # 任务处理器
│   ├── middleware/
│   │   ├── auth.go              # 认证中间件
│   │   ├── cors.go              # CORS中间件
│   │   └── logger.go            # 日志中间件
│   ├── models/
│   │   ├── user.go              # 用户模型
│   │   └── task.go              # 任务模型
│   ├── repository/
│   │   ├── user.go              # 用户数据访问
│   │   └── task.go              # 任务数据访问
│   ├── service/
│   │   ├── user.go              # 用户业务逻辑
│   │   └── task.go              # 任务业务逻辑
│   └── utils/
│       ├── response.go          # 统一响应
│       ├── jwt.go               # JWT工具
│       └── validator.go         # 验证工具
├── pkg/
│   └── logger/
│       └── logger.go            # 日志包
├── docs/
│   └── api.md                   # API文档
├── tests/
│   ├── handlers_test.go         # 处理器测试
│   └── integration_test.go      # 集成测试
├── configs/
│   ├── config.yaml              # 配置文件
│   └── config.example.yaml      # 配置示例
├── scripts/
│   ├── migrate.sql              # 数据库迁移
│   └── seed.sql                 # 测试数据
├── go.mod                       # Go模块文件
├── go.sum                       # 依赖校验
├── Dockerfile                   # Docker镜像
├── docker-compose.yml           # Docker编排
├── Makefile                     # 构建脚本
└── README.md                    # 项目说明
```

## 核心功能

### 1. 用户管理
- 用户注册
- 用户登录
- 用户信息查询
- 用户信息更新
- 用户列表（分页）

### 2. 任务管理
- 创建任务
- 查询任务列表
- 更新任务状态
- 删除任务
- 任务搜索和过滤

### 3. 认证授权
- JWT Token认证
- 角色权限控制
- 中间件拦截

## API设计

### 用户相关API
```
POST   /api/v1/users/register    # 用户注册
POST   /api/v1/users/login       # 用户登录
GET    /api/v1/users/profile     # 获取用户信息
PUT    /api/v1/users/profile     # 更新用户信息
GET    /api/v1/users             # 用户列表（管理员）
```

### 任务相关API
```
POST   /api/v1/tasks             # 创建任务
GET    /api/v1/tasks             # 任务列表
GET    /api/v1/tasks/:id         # 获取单个任务
PUT    /api/v1/tasks/:id         # 更新任务
DELETE /api/v1/tasks/:id         # 删除任务
```

## 数据模型

### 用户模型
```go
type User struct {
    ID        uint      `json:"id" gorm:"primaryKey"`
    Username  string    `json:"username" gorm:"uniqueIndex;not null" validate:"required,min=3,max=20"`
    Email     string    `json:"email" gorm:"uniqueIndex;not null" validate:"required,email"`
    Password  string    `json:"-" gorm:"not null" validate:"required,min=6"`
    Role      string    `json:"role" gorm:"default:user" validate:"oneof=admin user"`
    CreatedAt time.Time `json:"created_at"`
    UpdatedAt time.Time `json:"updated_at"`
    
    // 关联
    Tasks []Task `json:"tasks,omitempty" gorm:"foreignKey:UserID"`
}
```

### 任务模型
```go
type Task struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Title       string    `json:"title" gorm:"not null" validate:"required,max=100"`
    Description string    `json:"description" validate:"max=500"`
    Status      string    `json:"status" gorm:"default:pending" validate:"oneof=pending progress completed"`
    Priority    string    `json:"priority" gorm:"default:medium" validate:"oneof=low medium high"`
    DueDate     *time.Time `json:"due_date"`
    UserID      uint      `json:"user_id" gorm:"not null"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    
    // 关联
    User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
}
```

## 核心代码示例

### 主函数（类比Spring Boot Application）
```go
// Java: @SpringBootApplication
// Go: 手动配置和启动

package main

import (
    "log"
    "web-api/internal/config"
    "web-api/internal/database"
    "web-api/internal/handlers"
    "web-api/internal/middleware"
    
    "github.com/gin-gonic/gin"
)

func main() {
    // 加载配置
    cfg, err := config.Load()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }
    
    // 连接数据库
    db, err := database.Connect(cfg.Database)
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }
    
    // 创建Gin引擎
    r := gin.Default()
    
    // 注册中间件
    r.Use(middleware.Logger())
    r.Use(middleware.CORS())
    
    // 注册路由
    v1 := r.Group("/api/v1")
    {
        // 用户路由
        users := v1.Group("/users")
        {
            users.POST("/register", handlers.Register)
            users.POST("/login", handlers.Login)
            
            // 需要认证的路由
            auth := users.Use(middleware.Auth())
            {
                auth.GET("/profile", handlers.GetProfile)
                auth.PUT("/profile", handlers.UpdateProfile)
            }
        }
        
        // 任务路由（需要认证）
        tasks := v1.Group("/tasks").Use(middleware.Auth())
        {
            tasks.POST("", handlers.CreateTask)
            tasks.GET("", handlers.GetTasks)
            tasks.GET("/:id", handlers.GetTask)
            tasks.PUT("/:id", handlers.UpdateTask)
            tasks.DELETE("/:id", handlers.DeleteTask)
        }
    }
    
    // 启动服务器
    log.Printf("Server starting on port %s", cfg.Server.Port)
    if err := r.Run(":" + cfg.Server.Port); err != nil {
        log.Fatal("Failed to start server:", err)
    }
}
```

### 处理器（类比Spring Controller）
```go
// Java: @RestController
// Go: 普通函数

package handlers

import (
    "net/http"
    "strconv"
    "web-api/internal/models"
    "web-api/internal/service"
    "web-api/internal/utils"
    
    "github.com/gin-gonic/gin"
)

// 创建任务
func CreateTask(c *gin.Context) {
    // 获取当前用户
    userID, _ := c.Get("user_id")
    
    // 绑定请求数据
    var req models.CreateTaskRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid request data", err.Error())
        return
    }
    
    // 验证数据
    if err := utils.ValidateStruct(&req); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
        return
    }
    
    // 调用服务层
    task, err := service.CreateTask(userID.(uint), &req)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to create task", err.Error())
        return
    }
    
    utils.SuccessResponse(c, http.StatusCreated, "Task created successfully", task)
}

// 获取任务列表
func GetTasks(c *gin.Context) {
    userID, _ := c.Get("user_id")
    
    // 解析查询参数
    page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
    limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
    status := c.Query("status")
    priority := c.Query("priority")
    
    // 构建查询条件
    filter := &models.TaskFilter{
        Page:     page,
        Limit:    limit,
        Status:   status,
        Priority: priority,
        UserID:   userID.(uint),
    }
    
    // 调用服务层
    tasks, total, err := service.GetTasks(filter)
    if err != nil {
        utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to get tasks", err.Error())
        return
    }
    
    response := map[string]interface{}{
        "tasks": tasks,
        "pagination": map[string]interface{}{
            "page":  page,
            "limit": limit,
            "total": total,
        },
    }
    
    utils.SuccessResponse(c, http.StatusOK, "Tasks retrieved successfully", response)
}
```

### 服务层（类比Spring Service）
```go
// Java: @Service
// Go: 普通包

package service

import (
    "errors"
    "web-api/internal/models"
    "web-api/internal/repository"
)

func CreateTask(userID uint, req *models.CreateTaskRequest) (*models.Task, error) {
    // 业务逻辑验证
    if req.Title == "" {
        return nil, errors.New("task title is required")
    }
    
    // 创建任务对象
    task := &models.Task{
        Title:       req.Title,
        Description: req.Description,
        Status:      "pending",
        Priority:    req.Priority,
        DueDate:     req.DueDate,
        UserID:      userID,
    }
    
    // 调用数据访问层
    return repository.CreateTask(task)
}

func GetTasks(filter *models.TaskFilter) ([]models.Task, int64, error) {
    return repository.GetTasks(filter)
}
```

### 数据访问层（类比Spring Repository）
```go
// Java: @Repository
// Go: 普通包

package repository

import (
    "web-api/internal/database"
    "web-api/internal/models"
)

func CreateTask(task *models.Task) (*models.Task, error) {
    db := database.GetDB()
    
    if err := db.Create(task).Error; err != nil {
        return nil, err
    }
    
    return task, nil
}

func GetTasks(filter *models.TaskFilter) ([]models.Task, int64, error) {
    db := database.GetDB()
    
    var tasks []models.Task
    var total int64
    
    query := db.Model(&models.Task{}).Where("user_id = ?", filter.UserID)
    
    // 添加过滤条件
    if filter.Status != "" {
        query = query.Where("status = ?", filter.Status)
    }
    if filter.Priority != "" {
        query = query.Where("priority = ?", filter.Priority)
    }
    
    // 获取总数
    query.Count(&total)
    
    // 分页查询
    offset := (filter.Page - 1) * filter.Limit
    if err := query.Offset(offset).Limit(filter.Limit).
        Preload("User").Find(&tasks).Error; err != nil {
        return nil, 0, err
    }
    
    return tasks, total, nil
}
```

### 中间件（类比Spring Interceptor）
```go
// Java: @Component implements HandlerInterceptor
// Go: Gin中间件函数

package middleware

import (
    "net/http"
    "strings"
    "web-api/internal/utils"
    
    "github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        // 获取Authorization头
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            utils.ErrorResponse(c, http.StatusUnauthorized, "Authorization header required", nil)
            c.Abort()
            return
        }
        
        // 检查Bearer前缀
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == authHeader {
            utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid authorization format", nil)
            c.Abort()
            return
        }
        
        // 验证JWT
        claims, err := utils.ValidateJWT(tokenString)
        if err != nil {
            utils.ErrorResponse(c, http.StatusUnauthorized, "Invalid token", err.Error())
            c.Abort()
            return
        }
        
        // 设置用户信息到上下文
        c.Set("user_id", claims.UserID)
        c.Set("username", claims.Username)
        c.Set("role", claims.Role)
        
        c.Next()
    }
}
```

## 配置管理

### 配置文件（config.yaml）
```yaml
# 类比application.yml
server:
  port: "8080"
  mode: "debug"  # debug, release, test

database:
  host: "localhost"
  port: 3306
  username: "root"
  password: "password"
  dbname: "taskapi"
  charset: "utf8mb4"
  max_idle_conns: 10
  max_open_conns: 100

jwt:
  secret: "your-secret-key"
  expire_hours: 24

redis:
  host: "localhost"
  port: 6379
  password: ""
  db: 0

log:
  level: "info"
  format: "json"
  output: "stdout"
```

### 配置结构体
```go
type Config struct {
    Server   ServerConfig   `mapstructure:"server"`
    Database DatabaseConfig `mapstructure:"database"`
    JWT      JWTConfig      `mapstructure:"jwt"`
    Redis    RedisConfig    `mapstructure:"redis"`
    Log      LogConfig      `mapstructure:"log"`
}
```

## 数据库操作

### 数据库连接
```go
// Java: @Configuration + DataSource
// Go: 手动配置GORM

func Connect(cfg DatabaseConfig) (*gorm.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
        cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Charset)
    
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        return nil, err
    }
    
    // 配置连接池
    sqlDB, err := db.DB()
    if err != nil {
        return nil, err
    }
    
    sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
    sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
    
    // 自动迁移
    if err := db.AutoMigrate(&models.User{}, &models.Task{}); err != nil {
        return nil, err
    }
    
    return db, nil
}
```

## 测试

### 单元测试
```go
func TestCreateTask(t *testing.T) {
    // 设置测试数据库
    db := setupTestDB()
    defer cleanupTestDB(db)
    
    // 创建测试用户
    user := &models.User{
        Username: "testuser",
        Email:    "test@example.com",
        Password: "hashedpassword",
    }
    db.Create(user)
    
    // 测试创建任务
    req := &models.CreateTaskRequest{
        Title:       "Test Task",
        Description: "Test Description",
        Priority:    "high",
    }
    
    task, err := service.CreateTask(user.ID, req)
    
    assert.NoError(t, err)
    assert.NotNil(t, task)
    assert.Equal(t, "Test Task", task.Title)
    assert.Equal(t, user.ID, task.UserID)
}
```

### 集成测试
```go
func TestTaskAPI(t *testing.T) {
    // 设置测试服务器
    router := setupTestRouter()
    
    // 创建测试请求
    reqBody := `{"title":"Test Task","priority":"high"}`
    req, _ := http.NewRequest("POST", "/api/v1/tasks", strings.NewReader(reqBody))
    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("Authorization", "Bearer "+testToken)
    
    // 执行请求
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)
    
    // 验证响应
    assert.Equal(t, http.StatusCreated, w.Code)
    
    var response map[string]interface{}
    json.Unmarshal(w.Body.Bytes(), &response)
    assert.Equal(t, "success", response["status"])
}
```

## 部署

### Dockerfile
```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/configs ./configs

CMD ["./main"]
```

### docker-compose.yml
```yaml
version: '3.8'

services:
  api:
    build: .
    ports:
      - "8080:8080"
    environment:
      - GIN_MODE=release
    depends_on:
      - mysql
      - redis
    volumes:
      - ./configs:/root/configs

  mysql:
    image: mysql:8.0
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: taskapi
    ports:
      - "3306:3306"
    volumes:
      - mysql_data:/var/lib/mysql

  redis:
    image: redis:alpine
    ports:
      - "6379:6379"

volumes:
  mysql_data:
```

## 与Java Spring Boot对比

| 方面 | Java Spring Boot | Go Gin | 优劣对比 |
|------|------------------|---------|----------|
| 启动速度 | 较慢（~10-30s） | 快（~1-3s） | Go胜出 |
| 内存使用 | 较高（~200-500MB） | 低（~20-50MB） | Go胜出 |
| 开发效率 | 高（注解驱动） | 中等（手动配置） | Spring胜出 |
| 学习曲线 | 陡峭 | 平缓 | Go胜出 |
| 生态成熟度 | 非常成熟 | 快速发展 | Spring胜出 |
| 并发性能 | 好 | 优秀 | Go胜出 |

## 实践任务

1. **完成基础API**：实现所有用户和任务相关的API
2. **添加认证授权**：实现JWT认证和角色权限控制
3. **增加搜索功能**：支持任务的全文搜索
4. **实现文件上传**：支持任务附件上传
5. **添加缓存层**：使用Redis缓存热点数据
6. **性能测试**：使用压测工具测试API性能
7. **监控告警**：添加应用监控和日志系统

## 学习要点

1. **理解Go项目组织方式**：internal、pkg、cmd目录的作用
2. **掌握Gin框架核心概念**：路由、中间件、上下文
3. **学会GORM的使用**：模型定义、关联、查询
4. **理解Go的依赖管理**：go.mod、go.sum的作用
5. **掌握Go的错误处理**：error接口、错误包装

## 下一步

完成这个项目后，可以继续学习：
- 微服务架构（gRPC、服务发现）
- 消息队列（RabbitMQ、Kafka）
- 分布式系统（分布式锁、一致性）
- 云原生开发（Kubernetes、Docker）

这个项目提供了完整的Go Web开发体验，帮助Java开发者快速上手Go语言的Web开发。 