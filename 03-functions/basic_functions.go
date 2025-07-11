package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// 1. 基本函数定义
// Go函数语法：func 函数名(参数) 返回类型 { }
// Java对比：public returnType methodName(parameters) { }

// 无参数无返回值函数
func sayHello() {
	fmt.Println("Hello, Go!")
}

// 有参数的函数
func greet(name string) {
	fmt.Printf("Hello, %s!\n", name)
}

// 有返回值的函数
func add(a, b int) int {
	return a + b
}

// 多个参数相同类型的简写
func multiply(x, y, z int) int {
	return x * y * z
}

// 2. 多返回值 - Go特有特性
// Java需要使用对象或数组来返回多个值，Go原生支持
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零")
	}
	return a / b, nil
}

// 命名返回值 - Go特有
func calculateStats(numbers []int) (sum, count int, average float64) {
	count = len(numbers)
	if count == 0 {
		return // 自动返回命名的返回值
	}

	for _, num := range numbers {
		sum += num
	}
	average = float64(sum) / float64(count)
	return // 等同于 return sum, count, average
}

// 3. 可变参数函数 - 类似Java的varargs
func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

// 混合参数：固定参数 + 可变参数
func printWithPrefix(prefix string, messages ...string) {
	for _, msg := range messages {
		fmt.Printf("%s: %s\n", prefix, msg)
	}
}

// 4. 函数作为参数 - 高阶函数
// 类似Java的函数式接口和Lambda表达式
type Calculator func(int, int) int

func applyOperation(a, b int, op Calculator) int {
	return op(a, b)
}

// 预定义的计算函数
func addFunc(a, b int) int      { return a + b }
func subFunc(a, b int) int      { return a - b }
func multiplyFunc(a, b int) int { return a * b }

// 5. 函数作为返回值
func getCalculator(operation string) Calculator {
	switch operation {
	case "add":
		return addFunc
	case "sub":
		return subFunc
	case "multiply":
		return multiplyFunc
	default:
		return func(a, b int) int { return 0 }
	}
}

// 6. 匿名函数和闭包
func demonstrateClosures() {
	fmt.Println("\n=== 匿名函数和闭包 ===")

	// 简单匿名函数
	square := func(x int) int {
		return x * x
	}
	fmt.Printf("5的平方: %d\n", square(5))

	// 闭包：函数捕获外部变量
	counter := 0
	increment := func() int {
		counter++
		return counter
	}

	fmt.Printf("计数器: %d\n", increment()) // 1
	fmt.Printf("计数器: %d\n", increment()) // 2
	fmt.Printf("计数器: %d\n", increment()) // 3

	// 返回闭包的函数
	createMultiplier := func(factor int) func(int) int {
		return func(x int) int {
			return x * factor
		}
	}

	double := createMultiplier(2)
	triple := createMultiplier(3)

	fmt.Printf("8 × 2 = %d\n", double(8))
	fmt.Printf("8 × 3 = %d\n", triple(8))
}

// 7. 递归函数
func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

// 斐波那契数列（递归版本）
func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

// 斐波那契数列（迭代版本，性能更好）
func fibonacciIterative(n int) int {
	if n <= 1 {
		return n
	}

	a, b := 0, 1
	for i := 2; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 8. defer语句 - Go特有的延迟执行
func demonstrateDefer() {
	fmt.Println("\n=== defer语句演示 ===")

	fmt.Println("开始执行")

	// defer语句会在函数返回前执行，类似Java的finally
	defer fmt.Println("这会最后执行") // defer语句按LIFO顺序执行
	defer fmt.Println("这会倒数第二执行")

	fmt.Println("中间执行")

	// defer常用于资源清理
	file := "模拟文件句柄"
	defer func() {
		fmt.Printf("清理资源: %s\n", file)
	}()

	fmt.Println("即将结束")
	// 函数结束时，defer语句逆序执行
}

// 9. 错误处理函数
func validateAge(age int) error {
	if age < 0 {
		return fmt.Errorf("年龄不能为负数: %d", age)
	}
	if age > 150 {
		return fmt.Errorf("年龄过大: %d", age)
	}
	return nil
}

func parseAndValidateAge(ageStr string) (int, error) {
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		return 0, fmt.Errorf("年龄格式错误: %v", err)
	}

	if err := validateAge(age); err != nil {
		return 0, err
	}

	return age, nil
}

// 10. 实际业务场景函数
// 用户注册验证
func validateUser(username, password, email string) []string {
	var errors []string

	if len(username) < 3 {
		errors = append(errors, "用户名至少3个字符")
	}

	if len(password) < 8 {
		errors = append(errors, "密码至少8个字符")
	}

	if !strings.Contains(email, "@") {
		errors = append(errors, "邮箱格式不正确")
	}

	return errors
}

// 计算订单总价（含税）
func calculateOrderTotal(items []float64, taxRate float64) (subtotal, tax, total float64) {
	for _, price := range items {
		subtotal += price
	}
	tax = subtotal * taxRate
	total = subtotal + tax
	return
}

// 字符串处理工具函数
func processText(text string, operations ...func(string) string) string {
	result := text
	for _, op := range operations {
		result = op(result)
	}
	return result
}

// 11. 性能相关的函数优化
func inefficientStringConcat(words []string) string {
	// 低效：每次都创建新字符串
	result := ""
	for _, word := range words {
		result += word + " "
	}
	return strings.TrimSpace(result)
}

func efficientStringConcat(words []string) string {
	// 高效：使用StringBuilder
	var builder strings.Builder
	for i, word := range words {
		if i > 0 {
			builder.WriteString(" ")
		}
		builder.WriteString(word)
	}
	return builder.String()
}

// 12. 函数的单元测试示例（测试驱动开发）
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func basicFunctionExamples() {
	fmt.Println("=== 基本函数示例 ===")

	// 调用各种函数
	sayHello()
	greet("张三")

	result := add(5, 3)
	fmt.Printf("5 + 3 = %d\n", result)

	product := multiply(2, 3, 4)
	fmt.Printf("2 × 3 × 4 = %d\n", product)

	// 多返回值
	quotient, err := divide(10, 3)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("10 ÷ 3 = %.2f\n", quotient)
	}

	// 错误处理
	_, err = divide(10, 0)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	}

	// 命名返回值
	numbers := []int{1, 2, 3, 4, 5}
	s, c, avg := calculateStats(numbers)
	fmt.Printf("数组%v: 和=%d, 个数=%d, 平均值=%.2f\n", numbers, s, c, avg)
}

func variableArgumentExamples() {
	fmt.Println("\n=== 可变参数示例 ===")

	// 可变参数
	fmt.Printf("1 + 2 + 3 = %d\n", sum(1, 2, 3))
	fmt.Printf("1 + 2 + 3 + 4 + 5 = %d\n", sum(1, 2, 3, 4, 5))

	// 传递切片给可变参数
	nums := []int{10, 20, 30}
	fmt.Printf("切片求和: %d\n", sum(nums...)) // 展开切片

	// 混合参数
	printWithPrefix("日志", "用户登录", "数据更新", "系统启动")
}

func higherOrderFunctionExamples() {
	fmt.Println("\n=== 高阶函数示例 ===")

	// 函数作为参数
	fmt.Printf("10 + 5 = %d\n", applyOperation(10, 5, addFunc))
	fmt.Printf("10 - 5 = %d\n", applyOperation(10, 5, subFunc))
	fmt.Printf("10 × 5 = %d\n", applyOperation(10, 5, multiplyFunc))

	// 使用匿名函数
	fmt.Printf("10 的平方 = %d\n", applyOperation(10, 10, func(a, b int) int {
		return a * a
	}))

	// 函数作为返回值
	calc := getCalculator("add")
	fmt.Printf("使用返回的函数: 7 + 3 = %d\n", calc(7, 3))
}

func recursionExamples() {
	fmt.Println("\n=== 递归函数示例 ===")

	// 阶乘
	fmt.Printf("5! = %d\n", factorial(5))

	// 斐波那契数列对比
	n := 10
	fmt.Printf("斐波那契数列第%d项 (递归): %d\n", n, fibonacci(n))
	fmt.Printf("斐波那契数列第%d项 (迭代): %d\n", n, fibonacciIterative(n))
}

func errorHandlingExamples() {
	fmt.Println("\n=== 错误处理示例 ===")

	testAges := []string{"25", "-5", "abc", "200", "30"}

	for _, ageStr := range testAges {
		age, err := parseAndValidateAge(ageStr)
		if err != nil {
			fmt.Printf("年龄 '%s' 验证失败: %v\n", ageStr, err)
		} else {
			fmt.Printf("年龄 '%s' 验证成功: %d岁\n", ageStr, age)
		}
	}
}

func businessExamples() {
	fmt.Println("\n=== 业务场景示例 ===")

	// 用户注册验证
	errors := validateUser("ab", "123", "invalid-email")
	if len(errors) > 0 {
		fmt.Println("用户注册验证失败:")
		for _, err := range errors {
			fmt.Printf("  - %s\n", err)
		}
	}

	// 订单计算
	items := []float64{99.99, 199.99, 49.99}
	subtotal, tax, total := calculateOrderTotal(items, 0.08)
	fmt.Printf("订单明细: 小计=%.2f, 税费=%.2f, 总计=%.2f\n", subtotal, tax, total)

	// 文本处理链
	toUpper := strings.ToUpper
	addPrefix := func(s string) string { return "处理后: " + s }

	result := processText("hello world", toUpper, addPrefix)
	fmt.Printf("文本处理结果: %s\n", result)
}

func performanceExamples() {
	fmt.Println("\n=== 性能优化示例 ===")

	words := []string{"Go", "语言", "函数", "编程", "实践"}

	result1 := inefficientStringConcat(words)
	result2 := efficientStringConcat(words)

	fmt.Printf("低效拼接结果: %s\n", result1)
	fmt.Printf("高效拼接结果: %s\n", result2)

	// 质数判断
	fmt.Println("前20个数字的质数判断:")
	for i := 1; i <= 20; i++ {
		if isPrime(i) {
			fmt.Printf("%d是质数 ", i)
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("Go语言函数 - 基础实践")
	fmt.Println("=======================")

	basicFunctionExamples()
	variableArgumentExamples()
	higherOrderFunctionExamples()
	demonstrateClosures()
	recursionExamples()
	demonstrateDefer()
	errorHandlingExamples()
	businessExamples()
	performanceExamples()

	fmt.Println("\n学习要点:")
	fmt.Println("1. Go函数支持多返回值，常用于错误处理")
	fmt.Println("2. 命名返回值可以简化代码")
	fmt.Println("3. 可变参数使用...语法")
	fmt.Println("4. 函数是一等公民，可以作为参数和返回值")
	fmt.Println("5. defer语句用于延迟执行，常用于资源清理")
	fmt.Println("6. 错误处理是显式的，不使用异常机制")
}
