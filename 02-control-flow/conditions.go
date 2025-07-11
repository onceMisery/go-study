package main

import (
	"fmt"
	"strconv"
	"time"
)

// User 用户结构体 - 用于演示条件判断
type User struct {
	Name     string
	Age      int
	IsActive bool
	Role     string
}

// 1. 基本条件语句 - if/else
// Java对比：语法类似，但Go的if可以包含初始化语句
func basicConditions() {
	fmt.Println("=== 基本条件语句 ===")

	age := 25

	// 基本if语句
	if age >= 18 {
		fmt.Printf("年龄%d岁，已成年\n", age)
	}

	// if-else语句
	if age < 18 {
		fmt.Println("未成年")
	} else {
		fmt.Println("已成年")
	}

	// if-else if-else语句
	if age < 13 {
		fmt.Println("儿童")
	} else if age < 18 {
		fmt.Println("青少年")
	} else if age < 60 {
		fmt.Println("成年人")
	} else {
		fmt.Println("老年人")
	}

	// Go特有：if语句中的初始化
	// Java中需要在if语句外声明变量
	if score := 85; score >= 90 {
		fmt.Println("优秀")
	} else if score >= 80 {
		fmt.Println("良好")
	} else {
		fmt.Println("一般")
	}
}

// 2. 复杂条件判断
// 演示逻辑运算符和复杂条件组合
func complexConditions() {
	fmt.Println("\n=== 复杂条件判断 ===")

	user := User{
		Name:     "张三",
		Age:      28,
		IsActive: true,
		Role:     "admin",
	}

	// 逻辑与 (&&)
	if user.Age >= 18 && user.IsActive {
		fmt.Printf("用户%s可以访问系统\n", user.Name)
	}

	// 逻辑或 (||)
	if user.Role == "admin" || user.Role == "manager" {
		fmt.Printf("用户%s具有管理权限\n", user.Name)
	}

	// 逻辑非 (!)
	if !user.IsActive {
		fmt.Println("用户账户已禁用")
	} else {
		fmt.Println("用户账户正常")
	}

	// 复合条件
	if (user.Age >= 18 && user.IsActive) && (user.Role == "admin" || user.Role == "manager") {
		fmt.Printf("用户%s具有完整管理权限\n", user.Name)
	}
}

// 3. 条件语句中的类型断言和接口判断
// Go特有的类型相关条件判断
func typeConditions() {
	fmt.Println("\n=== 类型相关条件判断 ===")

	var value interface{} = "Hello, Go!"

	// 类型断言 - Go特有，Java中使用instanceof
	if str, ok := value.(string); ok {
		fmt.Printf("值是字符串: %s\n", str)
	} else {
		fmt.Println("值不是字符串")
	}

	// 多种类型判断
	checkType := func(v interface{}) {
		switch v.(type) {
		case string:
			fmt.Printf("字符串值: %v\n", v)
		case int:
			fmt.Printf("整数值: %v\n", v)
		case bool:
			fmt.Printf("布尔值: %v\n", v)
		default:
			fmt.Printf("未知类型: %T\n", v)
		}
	}

	checkType("Hello")
	checkType(42)
	checkType(true)
	checkType(3.14)
}

// 4. 错误处理中的条件判断
// Go的错误处理模式 - 与Java的try-catch不同
func errorHandlingConditions() {
	fmt.Println("\n=== 错误处理条件判断 ===")

	// 模拟可能出错的操作
	parseNumber := func(s string) (int, error) {
		return strconv.Atoi(s)
	}

	// Go典型的错误处理模式
	if num, err := parseNumber("123"); err != nil {
		fmt.Printf("解析失败: %v\n", err)
	} else {
		fmt.Printf("解析成功: %d\n", num)
	}

	// 处理错误的另一种方式
	if num, err := parseNumber("abc"); err != nil {
		fmt.Printf("无法解析'abc': %v\n", err)
		// 可以继续处理错误或返回
	} else {
		fmt.Printf("解析成功: %d\n", num)
	}
}

// 5. 实际业务场景中的条件判断
// 用户权限验证示例
func businessLogicConditions() {
	fmt.Println("\n=== 业务逻辑条件判断 ===")

	users := []User{
		{"张三", 25, true, "admin"},
		{"李四", 17, true, "user"},
		{"王五", 30, false, "manager"},
		{"赵六", 22, true, "user"},
	}

	for _, user := range users {
		fmt.Printf("\n检查用户: %s\n", user.Name)

		// 年龄验证
		if user.Age < 18 {
			fmt.Println("  - 未成年用户，限制访问")
			continue
		}

		// 账户状态验证
		if !user.IsActive {
			fmt.Println("  - 账户已禁用")
			continue
		}

		// 权限级别判断
		switch user.Role {
		case "admin":
			fmt.Println("  - 管理员权限：完全访问")
		case "manager":
			fmt.Println("  - 经理权限：部分管理功能")
		case "user":
			fmt.Println("  - 普通用户权限：基本功能")
		default:
			fmt.Println("  - 未知角色，拒绝访问")
		}

		// 特殊条件：VIP用户判断（模拟）
		if user.Age >= 25 && user.Role != "user" {
			fmt.Println("  - VIP用户，享受特殊服务")
		}
	}
}

// 6. 时间相关的条件判断
// 实际项目中常见的时间判断场景
func timeConditions() {
	fmt.Println("\n=== 时间相关条件判断 ===")

	now := time.Now()

	// 工作时间判断
	hour := now.Hour()
	if hour >= 9 && hour < 18 {
		fmt.Println("当前是工作时间")
	} else {
		fmt.Println("当前是非工作时间")
	}

	// 周末判断
	weekday := now.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		fmt.Println("今天是周末")
	} else {
		fmt.Println("今天是工作日")
	}

	// 特定日期判断
	birthday := time.Date(2024, 1, 1, 0, 0, 0, 0, time.Local)
	if now.After(birthday) {
		fmt.Println("生日已过")
	} else if now.Before(birthday) {
		fmt.Println("生日未到")
	} else {
		fmt.Println("今天是生日！")
	}

	// 时间差判断
	duration := time.Since(birthday)
	if duration.Hours() > 24*365 {
		fmt.Println("距离生日超过一年")
	}
}

func main() {
	fmt.Println("Go语言控制流 - 条件语句实践")
	fmt.Println("===============================")

	// 执行各种条件语句示例
	basicConditions()
	complexConditions()
	typeConditions()
	errorHandlingConditions()
	businessLogicConditions()
	timeConditions()

	fmt.Println("\n学习要点:")
	fmt.Println("1. Go的if语句可以包含初始化语句")
	fmt.Println("2. 错误处理使用if err != nil模式，而不是try-catch")
	fmt.Println("3. 类型断言用于接口类型判断")
	fmt.Println("4. 逻辑运算符与Java相同：&&, ||, !")
	fmt.Println("5. 注意边界条件和空值处理")
}
