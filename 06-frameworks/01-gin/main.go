package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ========== 数据模型 ==========

// User 用户模型
type User struct {
	ID       int       `json:"id"`
	Name     string    `json:"name" binding:"required"`
	Email    string    `json:"email" binding:"required,email"`
	Age      int       `json:"age" binding:"gte=0,lte=150"`
	CreateAt time.Time `json:"create_at"`
}

// Product 产品模型
type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Price       float64 `json:"price" binding:"required,gt=0"`
	CategoryID  int     `json:"category_id"`
}

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ========== 内存存储（模拟数据库） ==========

var (
	users    = make(map[int]*User)
	products = make(map[int]*Product)
	userID   = 1
	productID = 1
)

func init() {
	// 初始化一些测试数据
	users[1] = &User{
		ID:       1,
		Name:     "张三",
		Email:    "zhangsan@example.com",
		Age:      25,
		CreateAt: time.Now(),
	}
	
	products[1] = &Product{
		ID:          1,
		Name:        "Go语言编程",
		Description: "Go语言学习书籍",
		Price:       89.9,
		CategoryID:  1,
	}
	
	userID = 2
	productID = 2
}

// ========== 中间件 ==========

// Logger 自定义日志中间件
func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}

// CORS 跨域中间件
func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type,Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// AuthMiddleware 简单的认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		
		// 简单的token验证（实际项目中应该使用JWT等）
		if token == "" {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    401,
				Message: "缺少认证token",
			})
			c.Abort()
			return
		}
		
		if token != "Bearer valid-token" {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    401,
				Message: "无效的token",
			})
			c.Abort()
			return
		}
		
		// 设置用户信息到上下文
		c.Set("user_id", 1)
		c.Set("username", "admin")
		c.Next()
	}
}

// RateLimiter 简单的限流中间件
func RateLimiter() gin.HandlerFunc {
	// 简单的内存限流器（实际项目中应该使用Redis等）
	var requests = make(map[string][]time.Time)
	
	return func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()
		
		// 清理超过1分钟的记录
		if times, exists := requests[clientIP]; exists {
			var validTimes []time.Time
			for _, t := range times {
				if now.Sub(t) < time.Minute {
					validTimes = append(validTimes, t)
				}
			}
			requests[clientIP] = validTimes
		}
		
		// 检查请求频率（每分钟最多60次）
		if len(requests[clientIP]) >= 60 {
			c.JSON(http.StatusTooManyRequests, Response{
				Code:    429,
				Message: "请求太频繁，请稍后再试",
			})
			c.Abort()
			return
		}
		
		// 记录本次请求
		requests[clientIP] = append(requests[clientIP], now)
		c.Next()
	}
}

// ========== 处理器函数 ==========

// 首页处理器
func indexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "欢迎使用Go Gin API",
		Data: map[string]interface{}{
			"version": "1.0.0",
			"time":    time.Now(),
		},
	})
}

// 健康检查
func healthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "服务正常",
		Data: map[string]interface{}{
			"status": "healthy",
			"uptime": time.Since(startTime).String(),
		},
	})
}

var startTime = time.Now()

// ========== 用户相关处理器 ==========

// 获取所有用户
func getUsersHandler(c *gin.Context) {
	var userList []*User
	for _, user := range users {
		userList = append(userList, user)
	}
	
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取用户列表成功",
		Data:    userList,
	})
}

// 根据ID获取用户
func getUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}
	
	user, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}
	
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取用户成功",
		Data:    user,
	})
}

// 创建用户
func createUserHandler(c *gin.Context) {
	var user User
	
	// 绑定JSON数据并验证
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "数据验证失败: " + err.Error(),
		})
		return
	}
	
	// 设置用户信息
	user.ID = userID
	user.CreateAt = time.Now()
	users[userID] = &user
	userID++
	
	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "用户创建成功",
		Data:    user,
	})
}

// 更新用户
func updateUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}
	
	existingUser, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}
	
	var updateData User
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "数据验证失败: " + err.Error(),
		})
		return
	}
	
	// 更新用户信息
	existingUser.Name = updateData.Name
	existingUser.Email = updateData.Email
	existingUser.Age = updateData.Age
	
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "用户更新成功",
		Data:    existingUser,
	})
}

// 删除用户
func deleteUserHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}
	
	_, exists := users[id]
	if !exists {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}
	
	delete(users, id)
	
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "用户删除成功",
	})
}

// ========== 产品相关处理器 ==========

// 获取产品列表（带分页和搜索）
func getProductsHandler(c *gin.Context) {
	// 获取查询参数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	search := c.Query("search")
	
	var productList []*Product
	for _, product := range products {
		// 简单的搜索功能
		if search != "" && !contains(product.Name, search) && !contains(product.Description, search) {
			continue
		}
		productList = append(productList, product)
	}
	
	// 简单分页
	total := len(productList)
	start := (page - 1) * limit
	end := start + limit
	
	if start > total {
		productList = []*Product{}
	} else if end > total {
		productList = productList[start:]
	} else {
		productList = productList[start:end]
	}
	
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取产品列表成功",
		Data: map[string]interface{}{
			"products": productList,
			"total":    total,
			"page":     page,
			"limit":    limit,
		},
	})
}

// 简单的字符串包含检查
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 || 
		(len(s) > len(substr) && (s[:len(substr)] == substr || 
		s[len(s)-len(substr):] == substr || 
		strings.Contains(s, substr))))
}

// 创建产品
func createProductHandler(c *gin.Context) {
	var product Product
	
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "数据验证失败: " + err.Error(),
		})
		return
	}
	
	product.ID = productID
	products[productID] = &product
	productID++
	
	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "产品创建成功",
		Data:    product,
	})
}

// ========== 文件上传处理器 ==========

func uploadHandler(c *gin.Context) {
	// 限制文件大小为10MB
	c.Request.ParseMultipartForm(10 << 20)
	
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "文件上传失败: " + err.Error(),
		})
		return
	}
	defer file.Close()
	
	// 这里只是演示，实际项目中应该保存文件
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "文件上传成功",
		Data: map[string]interface{}{
			"filename": header.Filename,
			"size":     header.Size,
		},
	})
}

// ========== 主函数 ==========

func main() {
	// 设置Gin模式
	gin.SetMode(gin.ReleaseMode) // 生产环境使用
	
	// 创建Gin引擎
	r := gin.New()
	
	// 添加全局中间件
	r.Use(Logger())
	r.Use(gin.Recovery()) // 恢复中间件
	r.Use(CORS())
	r.Use(RateLimiter())
	
	// 基础路由
	r.GET("/", indexHandler)
	r.GET("/health", healthHandler)
	
	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 公开路由
		public := v1.Group("/public")
		{
			public.POST("/upload", uploadHandler)
		}
		
		// 需要认证的路由
		protected := v1.Group("/")
		protected.Use(AuthMiddleware())
		{
			// 用户相关路由
			users := protected.Group("/users")
			{
				users.GET("", getUsersHandler)
				users.GET("/:id", getUserHandler)
				users.POST("", createUserHandler)
				users.PUT("/:id", updateUserHandler)
				users.DELETE("/:id", deleteUserHandler)
			}
			
			// 产品相关路由
			products := protected.Group("/products")
			{
				products.GET("", getProductsHandler)
				products.POST("", createProductHandler)
			}
		}
	}
	
	// 启动服务器
	log.Println("服务器启动在 :8080")
	log.Fatal(r.Run(":8080"))
} 