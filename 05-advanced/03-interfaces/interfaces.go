package main

import (
	"fmt"
	"math"
	"sort"
)

// ========== 基础接口示例 ==========

// Shape 几何图形接口 - 类似Java中的interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Drawable 可绘制接口
type Drawable interface {
	Draw() string
}

// ColorfulShape 组合接口 - 接口嵌入
type ColorfulShape interface {
	Shape    // 嵌入Shape接口
	Drawable // 嵌入Drawable接口
	GetColor() string
}

// ========== 具体实现 ==========

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
	Color  string
}

// Area 实现Shape接口
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 实现Shape接口
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Draw 实现Drawable接口
func (r Rectangle) Draw() string {
	return fmt.Sprintf("绘制 %s 矩形: %.1f x %.1f", r.Color, r.Width, r.Height)
}

// GetColor 实现ColorfulShape接口
func (r Rectangle) GetColor() string {
	return r.Color
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
	Color  string
}

// Area 实现Shape接口
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter 实现Shape接口
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Draw 实现Drawable接口
func (c Circle) Draw() string {
	return fmt.Sprintf("绘制 %s 圆形: 半径 %.1f", c.Color, c.Radius)
}

// GetColor 实现ColorfulShape接口
func (c Circle) GetColor() string {
	return c.Color
}

// ========== 内置接口示例 ==========

// Person 实现fmt.Stringer接口
type Person struct {
	Name string
	Age  int
}

// String 实现fmt.Stringer接口 - 类似Java的toString()
func (p Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d}", p.Name, p.Age)
}

// ByAge 实现sort.Interface接口进行排序
type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ========== 自定义错误接口 ==========

// ValidationError 自定义错误类型
type ValidationError struct {
	Field   string
	Message string
}

// Error 实现error接口
func (e ValidationError) Error() string {
	return fmt.Sprintf("验证错误 [%s]: %s", e.Field, e.Message)
}

// Validator 验证器接口
type Validator interface {
	Validate() error
}

// User 用户结构体实现验证器接口
type User struct {
	Name  string
	Email string
	Age   int
}

// Validate 实现Validator接口
func (u User) Validate() error {
	if len(u.Name) == 0 {
		return ValidationError{Field: "Name", Message: "姓名不能为空"}
	}
	if len(u.Email) == 0 {
		return ValidationError{Field: "Email", Message: "邮箱不能为空"}
	}
	if u.Age < 0 || u.Age > 150 {
		return ValidationError{Field: "Age", Message: "年龄必须在0-150之间"}
	}
	return nil
}

// ========== 空接口和类型断言 ==========

// 空接口示例 - 类似Java的Object
func processAnyType(value interface{}) {
	fmt.Printf("类型: %T, 值: %v\n", value, value)

	// 类型断言
	switch v := value.(type) {
	case int:
		fmt.Printf("  这是一个整数: %d\n", v)
	case string:
		fmt.Printf("  这是一个字符串: %s\n", v)
	case Shape:
		fmt.Printf("  这是一个图形，面积: %.2f\n", v.Area())
	case Person:
		fmt.Printf("  这是一个人: %s\n", v.String())
	default:
		fmt.Printf("  未知类型\n")
	}
}

// ========== 接口作为函数参数 ==========

// PrintShapeInfo 接受Shape接口参数
func PrintShapeInfo(s Shape) {
	fmt.Printf("图形信息 - 面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())

	// 类型断言检查具体类型
	if drawable, ok := s.(Drawable); ok {
		fmt.Printf("绘制信息: %s\n", drawable.Draw())
	}

	if colorful, ok := s.(ColorfulShape); ok {
		fmt.Printf("颜色: %s\n", colorful.GetColor())
	}
}

// ========== 接口切片 ==========

// ShapeContainer 图形容器
type ShapeContainer struct {
	shapes []Shape
}

// Add 添加图形
func (sc *ShapeContainer) Add(shape Shape) {
	sc.shapes = append(sc.shapes, shape)
}

// TotalArea 计算总面积
func (sc *ShapeContainer) TotalArea() float64 {
	total := 0.0
	for _, shape := range sc.shapes {
		total += shape.Area()
	}
	return total
}

// DrawAll 绘制所有图形
func (sc *ShapeContainer) DrawAll() {
	fmt.Println("绘制所有图形:")
	for i, shape := range sc.shapes {
		fmt.Printf("%d. ", i+1)
		if drawable, ok := shape.(Drawable); ok {
			fmt.Println(drawable.Draw())
		} else {
			fmt.Printf("无法绘制的图形 (类型: %T)\n", shape)
		}
	}
}

// ========== 函数作为接口 ==========

// Handler 处理器函数类型
type Handler func(string) string

// ServeHTTP 让Handler实现某个接口
func (h Handler) ServeHTTP(path string) string {
	return h(path)
}

// HTTPHandler 网络处理器接口
type HTTPHandler interface {
	ServeHTTP(path string) string
}

// ========== 主函数演示 ==========

func main() {
	fmt.Println("=== Go 接口综合示例 ===")

	// 创建具体实例
	rect := Rectangle{Width: 10, Height: 5, Color: "红色"}
	circle := Circle{Radius: 3, Color: "蓝色"}

	fmt.Println("\n=== 基础接口使用 ===")
	PrintShapeInfo(rect)
	PrintShapeInfo(circle)

	// 接口切片
	fmt.Println("\n=== 接口切片 ===")
	shapes := []Shape{rect, circle, Rectangle{20, 10, "绿色"}}

	for i, shape := range shapes {
		fmt.Printf("图形 %d: 面积 %.2f\n", i+1, shape.Area())
	}

	// 图形容器
	fmt.Println("\n=== 图形容器 ===")
	container := &ShapeContainer{}
	container.Add(rect)
	container.Add(circle)
	container.Add(Rectangle{15, 8, "黄色"})

	fmt.Printf("总面积: %.2f\n", container.TotalArea())
	container.DrawAll()

	// 空接口和类型断言
	fmt.Println("\n=== 空接口和类型断言 ===")
	values := []interface{}{42, "Hello", rect, Person{"张三", 25}, 3.14}

	for _, value := range values {
		processAnyType(value)
	}

	// 内置接口 - Stringer
	fmt.Println("\n=== fmt.Stringer 接口 ===")
	person := Person{"李四", 30}
	fmt.Println(person) // 自动调用String()方法

	// 内置接口 - sort.Interface
	fmt.Println("\n=== sort.Interface 接口 ===")
	people := []Person{
		{"王五", 25},
		{"赵六", 30},
		{"孙七", 20},
	}

	fmt.Println("排序前:", people)
	sort.Sort(ByAge(people))
	fmt.Println("按年龄排序后:", people)

	// 自定义错误接口
	fmt.Println("\n=== 自定义错误接口 ===")
	users := []User{
		{"张三", "zhang@example.com", 25},
		{"", "li@example.com", 30},     // 无效：姓名为空
		{"王五", "", 35},                 // 无效：邮箱为空
		{"赵六", "zhao@example.com", -5}, // 无效：年龄负数
	}

	for i, user := range users {
		if err := user.Validate(); err != nil {
			fmt.Printf("用户 %d 验证失败: %v\n", i+1, err)
		} else {
			fmt.Printf("用户 %d 验证通过: %s\n", i+1, user.Name)
		}
	}

	// 函数作为接口
	fmt.Println("\n=== 函数作为接口 ===")
	var handler HTTPHandler = Handler(func(path string) string {
		return fmt.Sprintf("处理路径: %s", path)
	})

	result := handler.ServeHTTP("/api/users")
	fmt.Println(result)

	// 组合接口
	fmt.Println("\n=== 组合接口 ===")
	var colorfulShape ColorfulShape = rect
	fmt.Printf("彩色图形 - 面积: %.2f, 颜色: %s, 绘制: %s\n",
		colorfulShape.Area(), colorfulShape.GetColor(), colorfulShape.Draw())

	fmt.Println("\n=== 接口示例完成 ===")
}
