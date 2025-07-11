package main

import (
	"fmt"
	"time"
)

// 定义接口 - Go的接口是隐式实现的
type Worker interface {
	Work() string
	GetSalary() float64
}

type Manager interface {
	Worker // 接口嵌入
	ManageTeam(teamSize int) string
}

// Employee 员工结构体
type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float64
	HireDate   time.Time
	IsActive   bool
}

// NewEmployee 员工构造函数
func NewEmployee(id int, name, department string, salary float64) *Employee {
	return &Employee{
		ID:         id,
		Name:       name,
		Department: department,
		Salary:     salary,
		HireDate:   time.Now(),
		IsActive:   true,
	}
}

// Work 实现Worker接口
func (e *Employee) Work() string {
	return fmt.Sprintf("%s 正在 %s 部门工作", e.Name, e.Department)
}

// GetSalary 实现Worker接口
func (e *Employee) GetSalary() float64 {
	return e.Salary
}

// GiveFeedback 员工反馈方法
func (e *Employee) GiveFeedback(feedback string) {
	fmt.Printf("%s 收到反馈: %s\n", e.Name, feedback)
}

// TeamLead 团队负责人 - 组合Employee
type TeamLead struct {
	Employee   // 嵌入Employee结构体
	TeamSize   int
	Leadership string
}

// NewTeamLead 团队负责人构造函数
func NewTeamLead(employee *Employee, teamSize int, leadership string) *TeamLead {
	return &TeamLead{
		Employee:   *employee, // 嵌入Employee
		TeamSize:   teamSize,
		Leadership: leadership,
	}
}

// ManageTeam 实现Manager接口
func (t *TeamLead) ManageTeam(teamSize int) string {
	t.TeamSize = teamSize
	return fmt.Sprintf("%s 正在管理 %d 人的团队，领导风格: %s",
		t.Name, teamSize, t.Leadership)
}

// Work 重写父类方法
func (t *TeamLead) Work() string {
	baseWork := t.Employee.Work()
	return fmt.Sprintf("%s，同时负责团队管理", baseWork)
}

// Developer 开发者结构体
type Developer struct {
	Employee             // 嵌入Employee
	ProgrammingLanguages []string
	ProjectCount         int
}

// NewDeveloper 开发者构造函数
func NewDeveloper(employee *Employee, languages []string) *Developer {
	return &Developer{
		Employee:             *employee,
		ProgrammingLanguages: languages,
		ProjectCount:         0,
	}
}

// Work 开发者工作方法
func (d *Developer) Work() string {
	return fmt.Sprintf("%s 正在使用 %v 进行开发", d.Name, d.ProgrammingLanguages)
}

// Code 编程方法
func (d *Developer) Code(project string) string {
	d.ProjectCount++
	return fmt.Sprintf("%s 正在开发项目: %s (总项目数: %d)",
		d.Name, project, d.ProjectCount)
}

// Company 公司结构体 - 展示结构体数组的使用
type Company struct {
	Name      string
	Employees []Worker // 使用接口类型的切片
}

// NewCompany 公司构造函数
func NewCompany(name string) *Company {
	return &Company{
		Name:      name,
		Employees: make([]Worker, 0),
	}
}

// HireEmployee 雇佣员工
func (c *Company) HireEmployee(worker Worker) {
	c.Employees = append(c.Employees, worker)
	fmt.Printf("公司 %s 雇佣了新员工\n", c.Name)
}

// GetTotalSalary 计算总薪资
func (c *Company) GetTotalSalary() float64 {
	total := 0.0
	for _, emp := range c.Employees {
		total += emp.GetSalary()
	}
	return total
}

// PrintAllEmployees 打印所有员工信息
func (c *Company) PrintAllEmployees() {
	fmt.Printf("\n=== %s 员工名单 ===\n", c.Name)
	for i, emp := range c.Employees {
		fmt.Printf("%d. %s (薪资: %.2f)\n", i+1, emp.Work(), emp.GetSalary())

		// 类型断言 - 检查具体类型
		if manager, ok := emp.(Manager); ok {
			fmt.Printf("   管理信息: %s\n", manager.ManageTeam(5))
		}

		if dev, ok := emp.(*Developer); ok {
			fmt.Printf("   技能: %v\n", dev.ProgrammingLanguages)
		}
	}
}

func main() {
	fmt.Println("=== Go 结构体组合和接口示例 ===")

	// 创建基础员工
	emp1 := NewEmployee(1, "张三", "技术部", 8000.0)
	emp2 := NewEmployee(2, "李四", "产品部", 7000.0)

	// 创建开发者
	dev1 := NewDeveloper(
		NewEmployee(3, "王五", "技术部", 12000.0),
		[]string{"Go", "Java", "Python"},
	)

	// 创建团队负责人
	teamLead := NewTeamLead(
		NewEmployee(4, "赵六", "技术部", 15000.0),
		8,
		"敏捷开发",
	)

	// 创建公司并雇佣员工
	company := NewCompany("创新科技有限公司")
	company.HireEmployee(emp1)
	company.HireEmployee(emp2)
	company.HireEmployee(dev1)
	company.HireEmployee(teamLead)

	// 展示多态性
	fmt.Println("\n=== 多态性演示 ===")
	workers := []Worker{emp1, dev1, teamLead}
	for _, worker := range workers {
		fmt.Println(worker.Work())
	}

	// 开发者特有方法
	fmt.Println("\n=== 开发者工作 ===")
	fmt.Println(dev1.Code("电商系统"))
	fmt.Println(dev1.Code("支付网关"))

	// 公司统计
	fmt.Println("\n=== 公司统计 ===")
	company.PrintAllEmployees()
	fmt.Printf("总薪资支出: %.2f\n", company.GetTotalSalary())

	// 接口类型断言
	fmt.Println("\n=== 接口类型断言 ===")
	var worker Worker = teamLead
	if manager, ok := worker.(Manager); ok {
		fmt.Println("检测到管理者:", manager.ManageTeam(10))
	}
}
