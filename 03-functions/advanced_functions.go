package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// 1. 方法 (Methods) - 给类型添加行为
// 类似Java的实例方法，但Go的方法可以定义在任何类型上

// 自定义类型
type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

// 值接收者方法 - 类似Java的实例方法（但是传值）
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 指针接收者方法 - 可以修改接收者（推荐用法）
func (r *Rectangle) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

func (r *Rectangle) String() string {
	return fmt.Sprintf("Rectangle(%.2f x %.2f)", r.Width, r.Height)
}

// Circle的方法
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

func (c *Circle) Scale(factor float64) {
	c.Radius *= factor
}

func (c *Circle) String() string {
	return fmt.Sprintf("Circle(r=%.2f)", c.Radius)
}

// 2. 接口 - 定义行为契约
// 类似Java接口，但Go的接口是隐式实现的

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Scalable interface {
	Scale(factor float64)
}

// 组合接口
type ScalableShape interface {
	Shape
	Scalable
}

// 3. 函数类型和函数式编程
type Predicate func(int) bool
type Transformer func(int) int
type Reducer func(int, int) int

// 高阶函数：过滤
func Filter(slice []int, predicate Predicate) []int {
	var result []int
	for _, item := range slice {
		if predicate(item) {
			result = append(result, item)
		}
	}
	return result
}

// 高阶函数：映射
func Map(slice []int, transformer Transformer) []int {
	result := make([]int, len(slice))
	for i, item := range slice {
		result[i] = transformer(item)
	}
	return result
}

// 高阶函数：归约
func Reduce(slice []int, reducer Reducer, initial int) int {
	result := initial
	for _, item := range slice {
		result = reducer(result, item)
	}
	return result
}

// 4. 装饰器模式 - 使用函数包装函数
type Handler func(string) string

// 日志装饰器
func WithLogging(handler Handler) Handler {
	return func(input string) string {
		fmt.Printf("调用前: 输入=%s\n", input)
		result := handler(input)
		fmt.Printf("调用后: 输出=%s\n", result)
		return result
	}
}

// 计时装饰器
func WithTiming(handler Handler) Handler {
	return func(input string) string {
		start := time.Now()
		result := handler(input)
		duration := time.Since(start)
		fmt.Printf("执行耗时: %v\n", duration)
		return result
	}
}

// 5. 工厂模式和建造者模式
type Database struct {
	Host     string
	Port     int
	Username string
	Password string
	Name     string
}

// 工厂函数
func NewDatabase(host string, port int) *Database {
	return &Database{
		Host: host,
		Port: port,
	}
}

// 建造者模式的选项函数
type DatabaseOption func(*Database)

func WithCredentials(username, password string) DatabaseOption {
	return func(db *Database) {
		db.Username = username
		db.Password = password
	}
}

func WithDatabaseName(name string) DatabaseOption {
	return func(db *Database) {
		db.Name = name
	}
}

// 使用选项模式的构造函数
func NewDatabaseWithOptions(host string, port int, options ...DatabaseOption) *Database {
	db := &Database{
		Host: host,
		Port: port,
	}

	for _, option := range options {
		option(db)
	}

	return db
}

// 6. 管道模式 - 函数链式调用
type StringProcessor func(string) string

func (sp StringProcessor) Then(next StringProcessor) StringProcessor {
	return func(s string) string {
		return next(sp(s))
	}
}

// 字符串处理函数
func ToUpper(s string) string    { return strings.ToUpper(s) }
func AddPrefix(s string) string  { return "处理: " + s }
func AddSuffix(s string) string  { return s + " (完成)" }
func TrimSpaces(s string) string { return strings.TrimSpace(s) }

// 7. 泛型函数模拟 - 使用interface{}
func GenericFilter(slice interface{}, predicate func(interface{}) bool) []interface{} {
	// 注意：Go 1.18+有真正的泛型，这里是模拟实现
	var result []interface{}

	// 这里简化处理，实际需要反射
	switch s := slice.(type) {
	case []int:
		for _, item := range s {
			if predicate(item) {
				result = append(result, item)
			}
		}
	case []string:
		for _, item := range s {
			if predicate(item) {
				result = append(result, item)
			}
		}
	}

	return result
}

// 8. 回调和事件处理
type EventHandler func(string)

type EventManager struct {
	handlers map[string][]EventHandler
}

func NewEventManager() *EventManager {
	return &EventManager{
		handlers: make(map[string][]EventHandler),
	}
}

func (em *EventManager) Subscribe(event string, handler EventHandler) {
	em.handlers[event] = append(em.handlers[event], handler)
}

func (em *EventManager) Publish(event string, data string) {
	if handlers, exists := em.handlers[event]; exists {
		for _, handler := range handlers {
			handler(data)
		}
	}
}

// 9. 缓存和记忆化
type MemoizedFunc func(int) int

func Memoize(f func(int) int) MemoizedFunc {
	cache := make(map[int]int)

	return func(n int) int {
		if result, exists := cache[n]; exists {
			fmt.Printf("缓存命中: f(%d) = %d\n", n, result)
			return result
		}

		result := f(n)
		cache[n] = result
		fmt.Printf("计算结果: f(%d) = %d\n", n, result)
		return result
	}
}

// 10. 排序和比较函数
type Student struct {
	Name  string
	Age   int
	Score float64
}

func (s Student) String() string {
	return fmt.Sprintf("%s(年龄:%d, 成绩:%.1f)", s.Name, s.Age, s.Score)
}

// 实现sort.Interface
type ByScore []Student

func (s ByScore) Len() int           { return len(s) }
func (s ByScore) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ByScore) Less(i, j int) bool { return s[i].Score < s[j].Score }

// 使用函数进行排序
func SortStudents(students []Student, less func(Student, Student) bool) {
	sort.Slice(students, func(i, j int) bool {
		return less(students[i], students[j])
	})
}

// 示例函数
func methodExamples() {
	fmt.Println("=== 方法示例 ===")

	rect := Rectangle{Width: 5, Height: 3}
	fmt.Printf("矩形: %s\n", rect.String())
	fmt.Printf("面积: %.2f\n", rect.Area())
	fmt.Printf("周长: %.2f\n", rect.Perimeter())

	// 指针接收者方法
	rect.Scale(2)
	fmt.Printf("放大后: %s\n", rect.String())

	circle := Circle{Radius: 4}
	fmt.Printf("圆形: %s\n", circle.String())
	fmt.Printf("面积: %.2f\n", circle.Area())
}

func interfaceExamples() {
	fmt.Println("\n=== 接口示例 ===")

	shapes := []Shape{
		&Rectangle{Width: 4, Height: 3},
		&Circle{Radius: 2},
	}

	for _, shape := range shapes {
		fmt.Printf("形状: %s, 面积: %.2f, 周长: %.2f\n",
			shape, shape.Area(), shape.Perimeter())
	}

	// 类型断言
	for _, shape := range shapes {
		if scalable, ok := shape.(Scalable); ok {
			scalable.Scale(1.5)
			fmt.Printf("放大后: %s, 面积: %.2f\n", shape, shape.Area())
		}
	}
}

func functionalProgrammingExamples() {
	fmt.Println("\n=== 函数式编程示例 ===")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// 过滤偶数
	evens := Filter(numbers, func(n int) bool { return n%2 == 0 })
	fmt.Printf("偶数: %v\n", evens)

	// 映射平方
	squares := Map(numbers, func(n int) int { return n * n })
	fmt.Printf("平方: %v\n", squares)

	// 归约求和
	sum := Reduce(numbers, func(a, b int) int { return a + b }, 0)
	fmt.Printf("求和: %d\n", sum)

	// 链式操作
	result := Map(Filter(numbers, func(n int) bool { return n > 5 }),
		func(n int) int { return n * 2 })
	fmt.Printf("大于5的数乘以2: %v\n", result)
}

func decoratorExamples() {
	fmt.Println("\n=== 装饰器模式示例 ===")

	// 基础处理函数
	processor := func(input string) string {
		time.Sleep(10 * time.Millisecond) // 模拟处理时间
		return strings.ToUpper(input)
	}

	// 添加装饰器
	decoratedProcessor := WithLogging(WithTiming(processor))

	result := decoratedProcessor("hello world")
	fmt.Printf("最终结果: %s\n", result)
}

func builderPatternExamples() {
	fmt.Println("\n=== 建造者模式示例 ===")

	// 传统方式
	db1 := NewDatabase("localhost", 5432)
	db1.Username = "admin"
	db1.Password = "secret"
	db1.Name = "myapp"

	// 选项模式
	db2 := NewDatabaseWithOptions("localhost", 5432,
		WithCredentials("admin", "secret"),
		WithDatabaseName("myapp"))

	fmt.Printf("数据库1: %+v\n", db1)
	fmt.Printf("数据库2: %+v\n", db2)
}

func pipelineExamples() {
	fmt.Println("\n=== 管道模式示例 ===")

	// 构建处理管道
	pipeline := StringProcessor(TrimSpaces).
		Then(ToUpper).
		Then(AddPrefix).
		Then(AddSuffix)

	input := "  hello world  "
	result := pipeline(input)

	fmt.Printf("输入: '%s'\n", input)
	fmt.Printf("输出: '%s'\n", result)
}

func eventHandlingExamples() {
	fmt.Println("\n=== 事件处理示例 ===")

	em := NewEventManager()

	// 订阅事件
	em.Subscribe("user_login", func(data string) {
		fmt.Printf("日志记录: 用户登录 - %s\n", data)
	})

	em.Subscribe("user_login", func(data string) {
		fmt.Printf("发送通知: 欢迎 %s\n", data)
	})

	em.Subscribe("order_created", func(data string) {
		fmt.Printf("处理订单: %s\n", data)
	})

	// 发布事件
	em.Publish("user_login", "张三")
	em.Publish("order_created", "订单#12345")
}

func memoizationExamples() {
	fmt.Println("\n=== 记忆化示例 ===")

	// 斐波那契函数（低效版本）
	var fibonacci func(int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}
		return fibonacci(n-1) + fibonacci(n-2)
	}

	// 记忆化版本
	memoizedFib := Memoize(fibonacci)

	fmt.Println("计算斐波那契数列:")
	for i := 0; i <= 10; i++ {
		result := memoizedFib(i)
		fmt.Printf("fib(%d) = %d\n", i, result)
	}

	fmt.Println("\n重复计算（展示缓存效果）:")
	memoizedFib(8)
	memoizedFib(9)
	memoizedFib(10)
}

func sortingExamples() {
	fmt.Println("\n=== 排序示例 ===")

	students := []Student{
		{"张三", 20, 85.5},
		{"李四", 19, 92.0},
		{"王五", 21, 78.5},
		{"赵六", 20, 88.0},
	}

	fmt.Println("原始数据:")
	for _, s := range students {
		fmt.Printf("  %s\n", s)
	}

	// 按成绩排序
	sort.Sort(ByScore(students))
	fmt.Println("\n按成绩排序:")
	for _, s := range students {
		fmt.Printf("  %s\n", s)
	}

	// 使用函数排序：按年龄倒序
	SortStudents(students, func(a, b Student) bool {
		return a.Age > b.Age
	})
	fmt.Println("\n按年龄倒序:")
	for _, s := range students {
		fmt.Printf("  %s\n", s)
	}

	// 按姓名排序
	SortStudents(students, func(a, b Student) bool {
		return a.Name < b.Name
	})
	fmt.Println("\n按姓名排序:")
	for _, s := range students {
		fmt.Printf("  %s\n", s)
	}
}

func main() {
	fmt.Println("Go语言高级函数特性实践")
	fmt.Println("========================")

	methodExamples()
	interfaceExamples()
	functionalProgrammingExamples()
	decoratorExamples()
	builderPatternExamples()
	pipelineExamples()
	eventHandlingExamples()
	memoizationExamples()
	sortingExamples()

	fmt.Println("\n学习要点:")
	fmt.Println("1. 方法可以定义在任何类型上，使用接收者语法")
	fmt.Println("2. 接口是隐式实现的，提供了强大的多态性")
	fmt.Println("3. 函数是一等公民，支持函数式编程模式")
	fmt.Println("4. 装饰器模式可以优雅地扩展函数功能")
	fmt.Println("5. 选项模式是Go中常用的建造者模式实现")
	fmt.Println("6. 管道模式可以创建清晰的数据处理流程")
}
