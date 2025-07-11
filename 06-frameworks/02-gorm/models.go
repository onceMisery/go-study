package main

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ========== 数据模型定义 ==========

// User 用户模型 - 对比Java JPA实体
type User struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email     string    `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password  string    `gorm:"not null;size:255" json:"-"` // json:"-" 不序列化密码
	FullName  string    `gorm:"size:100" json:"full_name"`
	Age       int       `gorm:"check:age >= 0 AND age <= 150" json:"age"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"` // 软删除

	// 关联关系
	Posts    []Post    `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	Profile  Profile   `gorm:"foreignKey:UserID" json:"profile,omitempty"`
}

// TableName 自定义表名
func (User) TableName() string {
	return "users"
}

// BeforeCreate 创建前钩子 - 类似JPA的@PrePersist
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 可以在这里添加创建前的逻辑，如密码加密
	fmt.Printf("准备创建用户: %s\n", u.Username)
	return nil
}

// AfterCreate 创建后钩子 - 类似JPA的@PostPersist
func (u *User) AfterCreate(tx *gorm.DB) error {
	fmt.Printf("用户创建完成: ID=%d, Username=%s\n", u.ID, u.Username)
	return nil
}

// Profile 用户资料模型 - 一对一关系
type Profile struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"uniqueIndex;not null" json:"user_id"`
	Avatar      string    `gorm:"size:255" json:"avatar"`
	Bio         string    `gorm:"type:text" json:"bio"`
	Website     string    `gorm:"size:255" json:"website"`
	Location    string    `gorm:"size:100" json:"location"`
	Birthday    *time.Time `json:"birthday"` // 使用指针表示可空
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`

	// 关联关系
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Post 帖子模型 - 一对多关系
type Post struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint      `gorm:"not null;index" json:"user_id"`
	Title     string    `gorm:"not null;size:200" json:"title"`
	Content   string    `gorm:"type:longtext" json:"content"`
	Status    string    `gorm:"type:enum('draft','published','archived');default:'draft'" json:"status"`
	ViewCount int       `gorm:"default:0" json:"view_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	User User     `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Tags []Tag    `gorm:"many2many:post_tags;" json:"tags,omitempty"`
}

// Tag 标签模型 - 多对多关系
type Tag struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Color     string    `gorm:"size:7;default:'#007bff'" json:"color"` // 十六进制颜色
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联关系
	Posts []Post `gorm:"many2many:post_tags;" json:"posts,omitempty"`
}

// Category 分类模型 - 树形结构
type Category struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"not null;size:100" json:"name"`
	ParentID *uint  `gorm:"index" json:"parent_id"` // 使用指针表示可空
	Sort     int    `gorm:"default:0" json:"sort"`
	IsActive bool   `gorm:"default:true" json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 自关联
	Parent   *Category  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Children []Category `gorm:"foreignKey:ParentID" json:"children,omitempty"`
}

// ========== 数据库连接和初始化 ==========

var db *gorm.DB

func initDB() {
	var err error
	
	// 数据库连接字符串
	// 实际项目中应该从配置文件或环境变量读取
	dsn := "user:password@tcp(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"
	
	// 连接数据库
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// 禁用外键约束（可选）
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}
	
	// 自动迁移数据库表结构
	err = db.AutoMigrate(&User{}, &Profile{}, &Post{}, &Tag{}, &Category{})
	if err != nil {
		log.Fatal("数据库迁移失败:", err)
	}
	
	fmt.Println("数据库连接和迁移完成")
}

// ========== CRUD 操作示例 ==========

// 创建操作示例
func createExamples() {
	fmt.Println("\n=== 创建操作示例 ===")
	
	// 创建用户
	user := User{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "hashedpassword",
		FullName: "测试用户",
		Age:      25,
	}
	
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Printf("创建用户失败: %v\n", result.Error)
	} else {
		fmt.Printf("创建用户成功: ID=%d\n", user.ID)
	}
	
	// 创建用户资料
	profile := Profile{
		UserID:   user.ID,
		Bio:      "这是一个测试用户",
		Location: "北京",
	}
	
	db.Create(&profile)
	
	// 批量创建标签
	tags := []Tag{
		{Name: "Go语言", Color: "#00ADD8"},
		{Name: "Web开发", Color: "#007bff"},
		{Name: "后端", Color: "#28a745"},
	}
	
	db.Create(&tags)
	
	// 创建帖子并关联标签
	post := Post{
		UserID:  user.ID,
		Title:   "Go语言学习心得",
		Content: "Go语言是一门很棒的编程语言...",
		Status:  "published",
	}
	
	// 创建帖子后关联标签
	db.Create(&post)
	db.Model(&post).Association("Tags").Append(&tags[0], &tags[1])
	
	fmt.Printf("创建帖子成功: ID=%d\n", post.ID)
}

// 查询操作示例
func queryExamples() {
	fmt.Println("\n=== 查询操作示例 ===")
	
	// 1. 基础查询
	var user User
	db.First(&user, "username = ?", "testuser")
	fmt.Printf("查询用户: %+v\n", user)
	
	// 2. 预加载关联数据 - 类似JPA的fetch
	var userWithProfile User
	db.Preload("Profile").Preload("Posts").First(&userWithProfile, user.ID)
	fmt.Printf("用户及关联数据: %+v\n", userWithProfile)
	
	// 3. 条件查询
	var activeUsers []User
	db.Where("is_active = ? AND age > ?", true, 18).Find(&activeUsers)
	fmt.Printf("活跃成年用户数量: %d\n", len(activeUsers))
	
	// 4. 复杂查询 - Join
	var posts []Post
	db.Table("posts").
		Select("posts.*, users.username").
		Joins("JOIN users ON posts.user_id = users.id").
		Where("posts.status = ?", "published").
		Find(&posts)
	
	fmt.Printf("已发布帖子数量: %d\n", len(posts))
	
	// 5. 原生SQL查询
	var count int64
	db.Raw("SELECT COUNT(*) FROM users WHERE age BETWEEN ? AND ?", 20, 30).Scan(&count)
	fmt.Printf("20-30岁用户数量: %d\n", count)
	
	// 6. 聚合查询
	var result struct {
		AvgAge float64 `json:"avg_age"`
		MaxAge int     `json:"max_age"`
		MinAge int     `json:"min_age"`
	}
	
	db.Table("users").
		Select("AVG(age) as avg_age, MAX(age) as max_age, MIN(age) as min_age").
		Where("is_active = ?", true).
		Scan(&result)
	
	fmt.Printf("用户年龄统计: %+v\n", result)
	
	// 7. 分页查询
	var users []User
	var total int64
	
	page := 1
	pageSize := 10
	offset := (page - 1) * pageSize
	
	db.Model(&User{}).Where("is_active = ?", true).Count(&total)
	db.Where("is_active = ?", true).Limit(pageSize).Offset(offset).Find(&users)
	
	fmt.Printf("分页查询: 第%d页, 每页%d条, 总计%d条, 实际%d条\n", 
		page, pageSize, total, len(users))
}

// 更新操作示例
func updateExamples() {
	fmt.Println("\n=== 更新操作示例 ===")
	
	// 1. 更新单个字段
	db.Model(&User{}).Where("username = ?", "testuser").Update("age", 26)
	
	// 2. 更新多个字段
	db.Model(&User{}).Where("username = ?", "testuser").Updates(User{
		FullName: "更新的测试用户",
		Age:      27,
	})
	
	// 3. 使用map更新
	db.Model(&User{}).Where("username = ?", "testuser").Updates(map[string]interface{}{
		"full_name": "Map更新的用户",
		"age":       28,
	})
	
	// 4. 批量更新
	db.Model(&Post{}).Where("status = ?", "draft").Update("status", "published")
	
	// 5. 条件更新
	result := db.Model(&User{}).
		Where("age < ? AND is_active = ?", 30, true).
		Update("is_active", false)
	
	fmt.Printf("批量更新影响行数: %d\n", result.RowsAffected)
}

// 删除操作示例
func deleteExamples() {
	fmt.Println("\n=== 删除操作示例 ===")
	
	// 1. 软删除（因为模型中有DeletedAt字段）
	var user User
	db.Where("username = ?", "testuser").First(&user)
	db.Delete(&user) // 软删除，实际上是更新deleted_at字段
	
	// 2. 物理删除
	db.Unscoped().Delete(&user) // 真正从数据库删除
	
	// 3. 批量删除
	db.Where("is_active = ?", false).Delete(&User{})
	
	// 4. 删除关联
	var post Post
	db.First(&post)
	db.Model(&post).Association("Tags").Clear() // 清除帖子的所有标签关联
}

// ========== 事务示例 ==========

func transactionExample() {
	fmt.Println("\n=== 事务示例 ===")
	
	// 手动事务
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	
	// 在事务中创建用户
	user := User{
		Username: "txuser",
		Email:    "tx@example.com",
		Password: "password",
		FullName: "事务用户",
		Age:      30,
	}
	
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		fmt.Printf("创建用户失败: %v\n", err)
		return
	}
	
	// 在事务中创建资料
	profile := Profile{
		UserID:   user.ID,
		Bio:      "事务创建的用户资料",
		Location: "上海",
	}
	
	if err := tx.Create(&profile).Error; err != nil {
		tx.Rollback()
		fmt.Printf("创建资料失败: %v\n", err)
		return
	}
	
	// 提交事务
	if err := tx.Commit().Error; err != nil {
		fmt.Printf("提交事务失败: %v\n", err)
		return
	}
	
	fmt.Println("事务执行成功")
	
	// 使用DB.Transaction的自动事务
	err := db.Transaction(func(tx *gorm.DB) error {
		// 在这个函数中的所有数据库操作都在同一个事务中
		if err := tx.Create(&User{
			Username: "autotx",
			Email:    "autotx@example.com",
			Password: "password",
			FullName: "自动事务用户",
		}).Error; err != nil {
			return err // 返回错误会自动回滚
		}
		
		// 其他数据库操作...
		return nil // 返回nil会自动提交
	})
	
	if err != nil {
		fmt.Printf("自动事务失败: %v\n", err)
	} else {
		fmt.Println("自动事务执行成功")
	}
}

// ========== 高级查询示例 ==========

func advancedQueryExamples() {
	fmt.Println("\n=== 高级查询示例 ===")
	
	// 1. 子查询
	subQuery := db.Table("posts").Select("user_id").Where("status = ?", "published")
	var users []User
	db.Where("id IN (?)", subQuery).Find(&users)
	fmt.Printf("有已发布帖子的用户数量: %d\n", len(users))
	
	// 2. 联合查询 (Union)
	// GORM中需要使用原生SQL
	var results []map[string]interface{}
	db.Raw(`
		SELECT 'user' as type, id, username as name, created_at FROM users 
		UNION ALL 
		SELECT 'post' as type, id, title as name, created_at FROM posts
		ORDER BY created_at DESC 
		LIMIT 10
	`).Scan(&results)
	
	fmt.Printf("最新的用户和帖子: %d条\n", len(results))
	
	// 3. 窗口函数（MySQL 8.0+）
	var rankedPosts []struct {
		Post
		Rank int `json:"rank"`
	}
	
	db.Raw(`
		SELECT *, 
		       ROW_NUMBER() OVER (PARTITION BY user_id ORDER BY created_at DESC) as rank
		FROM posts 
		WHERE status = 'published'
	`).Scan(&rankedPosts)
	
	fmt.Printf("用户帖子排名: %d条\n", len(rankedPosts))
	
	// 4. 复杂的聚合查询
	var userStats []struct {
		UserID    uint `json:"user_id"`
		Username  string `json:"username"`
		PostCount int `json:"post_count"`
		AvgViews  float64 `json:"avg_views"`
	}
	
	db.Table("users").
		Select(`users.id as user_id, 
			    users.username, 
			    COUNT(posts.id) as post_count,
			    COALESCE(AVG(posts.view_count), 0) as avg_views`).
		Joins("LEFT JOIN posts ON users.id = posts.user_id").
		Group("users.id, users.username").
		Having("COUNT(posts.id) > 0").
		Scan(&userStats)
	
	fmt.Printf("用户统计数据: %d条\n", len(userStats))
}

// ========== 主函数 ==========

func main() {
	fmt.Println("=== GORM 数据库操作示例 ===")
	
	// 注意：这个示例需要MySQL数据库
	// 请先创建数据库: CREATE DATABASE go_demo;
	// 并修改上面的数据库连接字符串
	
	// 初始化数据库
	initDB()
	
	// 运行示例
	createExamples()
	queryExamples()
	updateExamples()
	transactionExample()
	advancedQueryExamples()
	
	// 注意：deleteExamples() 会删除数据，谨慎运行
	// deleteExamples()
	
	fmt.Println("\n=== GORM 示例完成 ===")
	fmt.Println("提示: 请确保MySQL数据库正在运行并且连接信息正确")
} 