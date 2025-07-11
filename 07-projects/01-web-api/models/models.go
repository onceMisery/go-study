package models

import (
	"time"

	"gorm.io/gorm"
)

// User 用户模型
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;not null;size:50" json:"username"`
	Email     string         `gorm:"uniqueIndex;not null;size:100" json:"email"`
	Password  string         `gorm:"not null;size:255" json:"-"` // 密码不返回给前端
	FirstName string         `gorm:"size:50" json:"first_name"`
	LastName  string         `gorm:"size:50" json:"last_name"`
	Avatar    string         `gorm:"size:255" json:"avatar"`
	Bio       string         `gorm:"type:text" json:"bio"`
	IsActive  bool           `gorm:"default:true" json:"is_active"`
	IsAdmin   bool           `gorm:"default:false" json:"is_admin"`
	LastLogin *time.Time     `json:"last_login"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Tasks []Task `gorm:"foreignKey:UserID" json:"tasks,omitempty"`
}

// Task 任务模型
type Task struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"not null;index" json:"user_id"`
	Title       string         `gorm:"not null;size:200" json:"title"`
	Description string         `gorm:"type:text" json:"description"`
	Status      TaskStatus     `gorm:"type:enum('pending','in_progress','completed','cancelled');default:'pending'" json:"status"`
	Priority    TaskPriority   `gorm:"type:enum('low','medium','high','urgent');default:'medium'" json:"priority"`
	DueDate     *time.Time     `json:"due_date"`
	CompletedAt *time.Time     `json:"completed_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	User User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Tags []Tag        `gorm:"many2many:task_tags;" json:"tags,omitempty"`
	Comments []Comment `gorm:"foreignKey:TaskID" json:"comments,omitempty"`
}

// TaskStatus 任务状态枚举
type TaskStatus string

const (
	TaskStatusPending    TaskStatus = "pending"
	TaskStatusInProgress TaskStatus = "in_progress"
	TaskStatusCompleted  TaskStatus = "completed"
	TaskStatusCancelled  TaskStatus = "cancelled"
)

// TaskPriority 任务优先级枚举
type TaskPriority string

const (
	TaskPriorityLow    TaskPriority = "low"
	TaskPriorityMedium TaskPriority = "medium"
	TaskPriorityHigh   TaskPriority = "high"
	TaskPriorityUrgent TaskPriority = "urgent"
)

// Tag 标签模型
type Tag struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"uniqueIndex;not null;size:50" json:"name"`
	Color     string    `gorm:"size:7;default:'#007bff'" json:"color"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	// 关联关系
	Tasks []Task `gorm:"many2many:task_tags;" json:"tasks,omitempty"`
}

// Comment 评论模型
type Comment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	TaskID    uint           `gorm:"not null;index" json:"task_id"`
	UserID    uint           `gorm:"not null;index" json:"user_id"`
	Content   string         `gorm:"not null;type:text" json:"content"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Task Task `gorm:"foreignKey:TaskID" json:"task,omitempty"`
	User User `gorm:"foreignKey:UserID" json:"user,omitempty"`
}

// Project 项目模型
type Project struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `gorm:"not null;size:100" json:"name"`
	Description string         `gorm:"type:text" json:"description"`
	Status      ProjectStatus  `gorm:"type:enum('planning','active','on_hold','completed','cancelled');default:'planning'" json:"status"`
	StartDate   *time.Time     `json:"start_date"`
	EndDate     *time.Time     `json:"end_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// 关联关系
	Members []User `gorm:"many2many:project_members;" json:"members,omitempty"`
}

// ProjectStatus 项目状态枚举
type ProjectStatus string

const (
	ProjectStatusPlanning  ProjectStatus = "planning"
	ProjectStatusActive    ProjectStatus = "active"
	ProjectStatusOnHold    ProjectStatus = "on_hold"
	ProjectStatusCompleted ProjectStatus = "completed"
	ProjectStatusCancelled ProjectStatus = "cancelled"
)

// TableName 自定义表名
func (User) TableName() string    { return "users" }
func (Task) TableName() string    { return "tasks" }
func (Tag) TableName() string     { return "tags" }
func (Comment) TableName() string { return "comments" }
func (Project) TableName() string { return "projects" }

// BeforeCreate 钩子函数
func (u *User) BeforeCreate(tx *gorm.DB) error {
	// 设置默认值或其他创建前的逻辑
	return nil
}

func (t *Task) BeforeCreate(tx *gorm.DB) error {
	// 任务创建前的逻辑
	return nil
}

func (t *Task) BeforeUpdate(tx *gorm.DB) error {
	// 如果状态改为已完成，设置完成时间
	if t.Status == TaskStatusCompleted && t.CompletedAt == nil {
		now := time.Now()
		t.CompletedAt = &now
	}
	return nil
} 