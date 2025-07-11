package services

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"go-demo/web-api/models"
)

// AuthService 认证服务
type AuthService struct {
	db        *gorm.DB
	jwtSecret []byte
}

// Claims JWT声明
type Claims struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	IsAdmin  bool   `json:"is_admin"`
	jwt.RegisteredClaims
}

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username  string `json:"username" binding:"required,min=3,max=50"`
	Email     string `json:"email" binding:"required,email"`
	Password  string `json:"password" binding:"required,min=6"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
}

// LoginResponse 登录响应
type LoginResponse struct {
	Token     string      `json:"token"`
	User      models.User `json:"user"`
	ExpiresAt time.Time   `json:"expires_at"`
}

// NewAuthService 创建认证服务
func NewAuthService(db *gorm.DB, jwtSecret string) *AuthService {
	return &AuthService{
		db:        db,
		jwtSecret: []byte(jwtSecret),
	}
}

// Register 用户注册
func (s *AuthService) Register(req RegisterRequest) (*models.User, error) {
	// 检查用户名是否已存在
	var existingUser models.User
	if err := s.db.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		return nil, errors.New("用户名或邮箱已存在")
	}

	// 加密密码
	hashedPassword, err := s.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := models.User{
		Username:  req.Username,
		Email:     req.Email,
		Password:  hashedPassword,
		FirstName: req.FirstName,
		LastName:  req.LastName,
		IsActive:  true,
		IsAdmin:   false,
	}

	if err := s.db.Create(&user).Error; err != nil {
		return nil, err
	}

	// 清除密码字段
	user.Password = ""
	return &user, nil
}

// Login 用户登录
func (s *AuthService) Login(req LoginRequest) (*LoginResponse, error) {
	// 查找用户
	var user models.User
	if err := s.db.Where("username = ? OR email = ?", req.Username, req.Username).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户名或密码错误")
		}
		return nil, err
	}

	// 检查用户是否激活
	if !user.IsActive {
		return nil, errors.New("账户已被禁用")
	}

	// 验证密码
	if !s.CheckPassword(user.Password, req.Password) {
		return nil, errors.New("用户名或密码错误")
	}

	// 更新最后登录时间
	now := time.Now()
	user.LastLogin = &now
	s.db.Model(&user).Update("last_login", now)

	// 生成JWT token
	token, expiresAt, err := s.GenerateToken(&user)
	if err != nil {
		return nil, err
	}

	// 清除密码字段
	user.Password = ""

	return &LoginResponse{
		Token:     token,
		User:      user,
		ExpiresAt: expiresAt,
	}, nil
}

// GenerateToken 生成JWT token
func (s *AuthService) GenerateToken(user *models.User) (string, time.Time, error) {
	expiresAt := time.Now().Add(24 * time.Hour) // 24小时过期

	claims := Claims{
		UserID:   user.ID,
		Username: user.Username,
		IsAdmin:  user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "go-demo-api",
			Subject:   user.Username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(s.jwtSecret)
	if err != nil {
		return "", time.Time{}, err
	}

	return tokenString, expiresAt, nil
}

// ValidateToken 验证JWT token
func (s *AuthService) ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名方法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("无效的签名方法")
		}
		return s.jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, errors.New("无效的token")
	}

	return claims, nil
}

// GetUserByID 根据ID获取用户
func (s *AuthService) GetUserByID(userID uint) (*models.User, error) {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return nil, err
	}

	// 清除密码字段
	user.Password = ""
	return &user, nil
}

// UpdatePassword 更新密码
func (s *AuthService) UpdatePassword(userID uint, oldPassword, newPassword string) error {
	// 获取用户
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return err
	}

	// 验证旧密码
	if !s.CheckPassword(user.Password, oldPassword) {
		return errors.New("原密码错误")
	}

	// 加密新密码
	hashedPassword, err := s.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新密码
	return s.db.Model(&user).Update("password", hashedPassword).Error
}

// ResetPassword 重置密码（管理员功能）
func (s *AuthService) ResetPassword(userID uint, newPassword string) error {
	// 加密新密码
	hashedPassword, err := s.HashPassword(newPassword)
	if err != nil {
		return err
	}

	// 更新密码
	return s.db.Model(&models.User{}).Where("id = ?", userID).Update("password", hashedPassword).Error
}

// ToggleUserStatus 切换用户状态（管理员功能）
func (s *AuthService) ToggleUserStatus(userID uint) error {
	var user models.User
	if err := s.db.First(&user, userID).Error; err != nil {
		return err
	}

	return s.db.Model(&user).Update("is_active", !user.IsActive).Error
}

// HashPassword 加密密码
func (s *AuthService) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

// CheckPassword 验证密码
func (s *AuthService) CheckPassword(hashedPassword, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

// RefreshToken 刷新token
func (s *AuthService) RefreshToken(tokenString string) (*LoginResponse, error) {
	// 验证当前token
	claims, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// 获取用户信息
	user, err := s.GetUserByID(claims.UserID)
	if err != nil {
		return nil, err
	}

	// 检查用户是否仍然激活
	if !user.IsActive {
		return nil, errors.New("账户已被禁用")
	}

	// 生成新token
	newToken, expiresAt, err := s.GenerateToken(user)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token:     newToken,
		User:      *user,
		ExpiresAt: expiresAt,
	}, nil
}

// GetUserProfile 获取用户资料
func (s *AuthService) GetUserProfile(userID uint) (*models.User, error) {
	var user models.User
	if err := s.db.Preload("Tasks", func(db *gorm.DB) *gorm.DB {
		return db.Where("status != ?", models.TaskStatusCompleted).Limit(5)
	}).First(&user, userID).Error; err != nil {
		return nil, err
	}

	// 清除密码字段
	user.Password = ""
	return &user, nil
}

// UpdateProfile 更新用户资料
func (s *AuthService) UpdateProfile(userID uint, updates map[string]interface{}) error {
	// 移除敏感字段
	delete(updates, "password")
	delete(updates, "is_admin")
	delete(updates, "id")

	return s.db.Model(&models.User{}).Where("id = ?", userID).Updates(updates).Error
} 