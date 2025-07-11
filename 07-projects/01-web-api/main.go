package main

import (
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"go-demo/web-api/models"
	"go-demo/web-api/services"
)

// Server 服务器结构体
type Server struct {
	db          *gorm.DB
	authService *services.AuthService
	router      *gin.Engine
}

// Response 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// PaginationQuery 分页查询参数
type PaginationQuery struct {
	Page  int `form:"page,default=1" binding:"min=1"`
	Limit int `form:"limit,default=10" binding:"min=1,max=100"`
}

// TaskQuery 任务查询参数
type TaskQuery struct {
	PaginationQuery
	Status   string `form:"status"`
	Priority string `form:"priority"`
	Search   string `form:"search"`
}

func main() {
	// 初始化服务器
	server := &Server{}
	
	// 初始化数据库
	server.initDB()
	
	// 初始化服务
	server.initServices()
	
	// 初始化路由
	server.initRoutes()
	
	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("服务器启动在端口 :%s", port)
	log.Fatal(server.router.Run(":" + port))
}

// initDB 初始化数据库
func (s *Server) initDB() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "root:password@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	}

	var err error
	s.db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	// 自动迁移
	err = s.db.AutoMigrate(
		&models.User{},
		&models.Task{},
		&models.Tag{},
		&models.Comment{},
		&models.Project{},
	)
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}

	log.Println("数据库连接成功")
}

// initServices 初始化服务
func (s *Server) initServices() {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "your-secret-key-here-change-in-production"
	}

	s.authService = services.NewAuthService(s.db, jwtSecret)
}

// initRoutes 初始化路由
func (s *Server) initRoutes() {
	// 设置Gin模式
	if os.Getenv("GIN_MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	s.router = gin.New()

	// 中间件
	s.router.Use(gin.Logger())
	s.router.Use(gin.Recovery())
	s.router.Use(s.corsMiddleware())

	// 健康检查
	s.router.GET("/health", s.healthCheck)

	// API路由组
	api := s.router.Group("/api/v1")
	{
		// 认证相关路由
		auth := api.Group("/auth")
		{
			auth.POST("/register", s.register)
			auth.POST("/login", s.login)
			auth.POST("/refresh", s.refreshToken)
		}

		// 需要认证的路由
		protected := api.Group("/")
		protected.Use(s.authMiddleware())
		{
			// 用户相关
			users := protected.Group("/users")
			{
				users.GET("/profile", s.getUserProfile)
				users.PUT("/profile", s.updateUserProfile)
				users.POST("/change-password", s.changePassword)
			}

			// 任务相关
			tasks := protected.Group("/tasks")
			{
				tasks.GET("", s.getTasks)
				tasks.POST("", s.createTask)
				tasks.GET("/:id", s.getTask)
				tasks.PUT("/:id", s.updateTask)
				tasks.DELETE("/:id", s.deleteTask)
				tasks.POST("/:id/comments", s.addTaskComment)
			}

			// 标签相关
			tags := protected.Group("/tags")
			{
				tags.GET("", s.getTags)
				tags.POST("", s.createTag)
				tags.PUT("/:id", s.updateTag)
				tags.DELETE("/:id", s.deleteTag)
			}

			// 管理员路由
			admin := protected.Group("/admin")
			admin.Use(s.adminMiddleware())
			{
				admin.GET("/users", s.getAllUsers)
				admin.PUT("/users/:id/status", s.toggleUserStatus)
				admin.POST("/users/:id/reset-password", s.resetUserPassword)
			}
		}
	}
}

// ========== 中间件 ==========

// corsMiddleware CORS中间件
func (s *Server) corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

// authMiddleware 认证中间件
func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    401,
				Message: "缺少认证token",
			})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    401,
				Message: "认证token格式错误",
			})
			c.Abort()
			return
		}

		claims, err := s.authService.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, Response{
				Code:    401,
				Message: "无效的token: " + err.Error(),
			})
			c.Abort()
			return
		}

		// 设置用户信息到上下文
		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("is_admin", claims.IsAdmin)
		c.Next()
	}
}

// adminMiddleware 管理员中间件
func (s *Server) adminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		isAdmin, exists := c.Get("is_admin")
		if !exists || !isAdmin.(bool) {
			c.JSON(http.StatusForbidden, Response{
				Code:    403,
				Message: "需要管理员权限",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}

// ========== 处理器函数 ==========

// healthCheck 健康检查
func (s *Server) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "服务正常",
		Data: map[string]interface{}{
			"status":    "healthy",
			"timestamp": time.Now(),
			"version":   "1.0.0",
		},
	})
}

// register 用户注册
func (s *Server) register(c *gin.Context) {
	var req services.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	user, err := s.authService.Register(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "注册成功",
		Data:    user,
	})
}

// login 用户登录
func (s *Server) login(c *gin.Context) {
	var req services.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	result, err := s.authService.Login(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "登录成功",
		Data:    result,
	})
}

// refreshToken 刷新token
func (s *Server) refreshToken(c *gin.Context) {
	var req struct {
		Token string `json:"token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	result, err := s.authService.RefreshToken(req.Token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, Response{
			Code:    401,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "Token刷新成功",
		Data:    result,
	})
}

// getUserProfile 获取用户资料
func (s *Server) getUserProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	user, err := s.authService.GetUserProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "用户不存在",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取用户资料成功",
		Data:    user,
	})
}

// updateUserProfile 更新用户资料
func (s *Server) updateUserProfile(c *gin.Context) {
	userID := c.GetUint("user_id")

	var updates map[string]interface{}
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if err := s.authService.UpdateProfile(userID, updates); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "更新失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "资料更新成功",
	})
}

// changePassword 修改密码
func (s *Server) changePassword(c *gin.Context) {
	userID := c.GetUint("user_id")

	var req struct {
		OldPassword string `json:"old_password" binding:"required"`
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if err := s.authService.UpdatePassword(userID, req.OldPassword, req.NewPassword); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "密码修改成功",
	})
}

// getTasks 获取任务列表
func (s *Server) getTasks(c *gin.Context) {
	userID := c.GetUint("user_id")

	var query TaskQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "查询参数错误: " + err.Error(),
		})
		return
	}

	db := s.db.Model(&models.Task{}).Where("user_id = ?", userID)

	// 添加过滤条件
	if query.Status != "" {
		db = db.Where("status = ?", query.Status)
	}
	if query.Priority != "" {
		db = db.Where("priority = ?", query.Priority)
	}
	if query.Search != "" {
		db = db.Where("title LIKE ? OR description LIKE ?", "%"+query.Search+"%", "%"+query.Search+"%")
	}

	// 计算总数
	var total int64
	db.Count(&total)

	// 分页查询
	var tasks []models.Task
	offset := (query.Page - 1) * query.Limit
	if err := db.Preload("Tags").Preload("User").
		Offset(offset).Limit(query.Limit).
		Order("created_at DESC").
		Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "查询失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取任务列表成功",
		Data: map[string]interface{}{
			"tasks": tasks,
			"pagination": map[string]interface{}{
				"page":  query.Page,
				"limit": query.Limit,
				"total": total,
			},
		},
	})
}

// createTask 创建任务
func (s *Server) createTask(c *gin.Context) {
	userID := c.GetUint("user_id")

	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	task.UserID = userID

	if err := s.db.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "任务创建成功",
		Data:    task,
	})
}

// getTask 获取单个任务
func (s *Server) getTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的任务ID",
		})
		return
	}

	var task models.Task
	if err := s.db.Preload("Tags").Preload("Comments.User").
		Where("id = ? AND user_id = ?", taskID, userID).
		First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "任务不存在",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取任务成功",
		Data:    task,
	})
}

// updateTask 更新任务
func (s *Server) updateTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的任务ID",
		})
		return
	}

	var updates models.Task
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	result := s.db.Model(&models.Task{}).
		Where("id = ? AND user_id = ?", taskID, userID).
		Updates(&updates)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "更新失败: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "任务不存在",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "任务更新成功",
	})
}

// deleteTask 删除任务
func (s *Server) deleteTask(c *gin.Context) {
	userID := c.GetUint("user_id")
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的任务ID",
		})
		return
	}

	result := s.db.Where("id = ? AND user_id = ?", taskID, userID).Delete(&models.Task{})
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "删除失败: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "任务不存在",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "任务删除成功",
	})
}

// addTaskComment 添加任务评论
func (s *Server) addTaskComment(c *gin.Context) {
	userID := c.GetUint("user_id")
	taskID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的任务ID",
		})
		return
	}

	var req struct {
		Content string `json:"content" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	// 验证任务是否存在且属于当前用户
	var task models.Task
	if err := s.db.Where("id = ? AND user_id = ?", taskID, userID).First(&task).Error; err != nil {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "任务不存在",
		})
		return
	}

	comment := models.Comment{
		TaskID:  uint(taskID),
		UserID:  userID,
		Content: req.Content,
	}

	if err := s.db.Create(&comment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "添加评论失败: " + err.Error(),
		})
		return
	}

	// 预加载用户信息
	s.db.Preload("User").First(&comment, comment.ID)

	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "评论添加成功",
		Data:    comment,
	})
}

// getTags 获取标签列表
func (s *Server) getTags(c *gin.Context) {
	var tags []models.Tag
	if err := s.db.Find(&tags).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "查询失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取标签列表成功",
		Data:    tags,
	})
}

// createTag 创建标签
func (s *Server) createTag(c *gin.Context) {
	var tag models.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if err := s.db.Create(&tag).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "创建失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, Response{
		Code:    201,
		Message: "标签创建成功",
		Data:    tag,
	})
}

// updateTag 更新标签
func (s *Server) updateTag(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的标签ID",
		})
		return
	}

	var updates models.Tag
	if err := c.ShouldBindJSON(&updates); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	result := s.db.Model(&models.Tag{}).Where("id = ?", tagID).Updates(&updates)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "更新失败: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "标签不存在",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "标签更新成功",
	})
}

// deleteTag 删除标签
func (s *Server) deleteTag(c *gin.Context) {
	tagID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的标签ID",
		})
		return
	}

	result := s.db.Delete(&models.Tag{}, tagID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "删除失败: " + result.Error.Error(),
		})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, Response{
			Code:    404,
			Message: "标签不存在",
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "标签删除成功",
	})
}

// getAllUsers 获取所有用户（管理员）
func (s *Server) getAllUsers(c *gin.Context) {
	var query PaginationQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "查询参数错误: " + err.Error(),
		})
		return
	}

	var users []models.User
	var total int64

	db := s.db.Model(&models.User{})
	db.Count(&total)

	offset := (query.Page - 1) * query.Limit
	if err := db.Offset(offset).Limit(query.Limit).Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "查询失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "获取用户列表成功",
		Data: map[string]interface{}{
			"users": users,
			"pagination": map[string]interface{}{
				"page":  query.Page,
				"limit": query.Limit,
				"total": total,
			},
		},
	})
}

// toggleUserStatus 切换用户状态（管理员）
func (s *Server) toggleUserStatus(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	if err := s.authService.ToggleUserStatus(uint(userID)); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "操作失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "用户状态切换成功",
	})
}

// resetUserPassword 重置用户密码（管理员）
func (s *Server) resetUserPassword(c *gin.Context) {
	userID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "无效的用户ID",
		})
		return
	}

	var req struct {
		NewPassword string `json:"new_password" binding:"required,min=6"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			Code:    400,
			Message: "请求参数错误: " + err.Error(),
		})
		return
	}

	if err := s.authService.ResetPassword(uint(userID), req.NewPassword); err != nil {
		c.JSON(http.StatusInternalServerError, Response{
			Code:    500,
			Message: "重置密码失败: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "密码重置成功",
	})
} 