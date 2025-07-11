package main

import (
	"fmt"
	"runtime"
	"time"
)

// OrderStatus 订单状态枚举 - 使用Go的常量模拟Java枚举
type OrderStatus int

const (
	Pending OrderStatus = iota
	Processing
	Shipped
	Delivered
	Cancelled
)

// String 实现Stringer接口，类似Java的toString()
func (os OrderStatus) String() string {
	switch os {
	case Pending:
		return "待处理"
	case Processing:
		return "处理中"
	case Shipped:
		return "已发货"
	case Delivered:
		return "已送达"
	case Cancelled:
		return "已取消"
	default:
		return "未知状态"
	}
}

// 1. 基本switch语句
// Go的switch比Java更灵活，默认break，支持多种类型
func basicSwitch() {
	fmt.Println("=== 基本switch语句 ===")

	// 传统switch（类似Java，但默认break）
	day := 3
	fmt.Printf("今天是星期: ")
	switch day {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三") // Go中不需要break，默认会break
	case 4:
		fmt.Println("四")
	case 5:
		fmt.Println("五")
	case 6, 7: // 多个值匹配（Go特有）
		fmt.Println("周末")
	default:
		fmt.Println("无效日期")
	}

	// 字符串switch
	grade := "A"
	fmt.Printf("成绩等级%s: ", grade)
	switch grade {
	case "A":
		fmt.Println("优秀 (90-100分)")
	case "B":
		fmt.Println("良好 (80-89分)")
	case "C":
		fmt.Println("中等 (70-79分)")
	case "D":
		fmt.Println("及格 (60-69分)")
	case "F":
		fmt.Println("不及格 (<60分)")
	default:
		fmt.Println("无效等级")
	}
}

// 2. 表达式switch - Go特有
// 不需要匹配变量，可以在每个case中使用表达式
func expressionSwitch() {
	fmt.Println("\n=== 表达式switch ===")

	score := 85
	fmt.Printf("分数%d的等级: ", score)

	// 无匹配变量的switch（Go特有特性）
	switch {
	case score >= 90:
		fmt.Println("优秀")
	case score >= 80:
		fmt.Println("良好")
	case score >= 70:
		fmt.Println("中等")
	case score >= 60:
		fmt.Println("及格")
	default:
		fmt.Println("不及格")
	}

	// 时间相关的switch
	hour := time.Now().Hour()
	fmt.Printf("当前时间%d点: ", hour)
	switch {
	case hour >= 6 && hour < 12:
		fmt.Println("上午")
	case hour >= 12 && hour < 18:
		fmt.Println("下午")
	case hour >= 18 && hour < 22:
		fmt.Println("晚上")
	default:
		fmt.Println("深夜")
	}
}

// 3. 类型switch - Go特有
// 用于判断接口的具体类型
func typeSwitch() {
	fmt.Println("\n=== 类型switch ===")

	// 不同类型的值
	values := []interface{}{
		42,
		"Hello, Go!",
		3.14,
		true,
		[]int{1, 2, 3},
		map[string]int{"key": 100},
		nil,
	}

	for i, value := range values {
		fmt.Printf("值%d: ", i+1)

		// 类型switch - Go特有语法
		switch v := value.(type) {
		case nil:
			fmt.Println("空值")
		case bool:
			fmt.Printf("布尔值: %t\n", v)
		case int:
			fmt.Printf("整数: %d\n", v)
		case float64:
			fmt.Printf("浮点数: %.2f\n", v)
		case string:
			fmt.Printf("字符串: %s (长度: %d)\n", v, len(v))
		case []int:
			fmt.Printf("整数切片: %v (长度: %d)\n", v, len(v))
		case map[string]int:
			fmt.Printf("字符串到整数的映射: %v\n", v)
		default:
			fmt.Printf("未知类型: %T, 值: %v\n", v, v)
		}
	}
}

// 4. fallthrough - 穿透到下一个case
// 类似Java中不写break的效果
func fallthroughSwitch() {
	fmt.Println("\n=== fallthrough示例 ===")

	month := 2
	fmt.Printf("%d月份特点: ", month)

	switch month {
	case 12, 1, 2:
		fmt.Print("冬季")
		fallthrough // 继续执行下一个case
	case 6, 7, 8:
		if month >= 6 && month <= 8 {
			fmt.Print("夏季")
		}
		fmt.Print(" - 极端天气")
		fallthrough
	default:
		fmt.Println(" - 注意保暖/防暑")
	}
}

// 5. 业务场景中的switch应用
func businessSwitch() {
	fmt.Println("\n=== 业务场景switch应用 ===")

	// 订单状态处理
	orders := []struct {
		ID     string
		Status OrderStatus
	}{
		{"ORD001", Pending},
		{"ORD002", Processing},
		{"ORD003", Shipped},
		{"ORD004", Delivered},
		{"ORD005", Cancelled},
	}

	fmt.Println("1. 订单状态处理:")
	for _, order := range orders {
		fmt.Printf("订单%s (%s): ", order.ID, order.Status)

		switch order.Status {
		case Pending:
			fmt.Println("等待处理 - 可以修改订单")
		case Processing:
			fmt.Println("正在处理 - 无法修改订单")
		case Shipped:
			fmt.Println("已发货 - 可以跟踪物流")
		case Delivered:
			fmt.Println("已送达 - 可以评价")
		case Cancelled:
			fmt.Println("已取消 - 可以重新下单")
		default:
			fmt.Println("未知状态 - 请联系客服")
		}
	}

	// HTTP状态码处理
	fmt.Println("\n2. HTTP状态码处理:")
	statusCodes := []int{200, 201, 400, 401, 403, 404, 500, 502}

	for _, code := range statusCodes {
		fmt.Printf("状态码%d: ", code)

		switch {
		case code >= 200 && code < 300:
			fmt.Println("成功响应")
		case code >= 300 && code < 400:
			fmt.Println("重定向")
		case code >= 400 && code < 500:
			fmt.Println("客户端错误")
		case code >= 500:
			fmt.Println("服务器错误")
		default:
			fmt.Println("无效状态码")
		}
	}
}

// 6. 高级switch用法
func advancedSwitch() {
	fmt.Println("\n=== 高级switch用法 ===")

	// switch中的初始化语句
	fmt.Println("1. switch中的初始化:")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("运行在macOS系统")
	case "linux":
		fmt.Println("运行在Linux系统")
	case "windows":
		fmt.Println("运行在Windows系统")
	default:
		fmt.Printf("运行在%s系统\n", os)
	}

	// 复杂条件的switch
	fmt.Println("2. 复杂条件switch:")
	user := struct {
		Name string
		Age  int
		Role string
	}{"张三", 25, "admin"}

	switch {
	case user.Age < 18:
		fmt.Printf("用户%s未成年，限制访问\n", user.Name)
	case user.Age >= 18 && user.Role == "admin":
		fmt.Printf("管理员%s，完全访问权限\n", user.Name)
	case user.Age >= 18 && user.Role == "user":
		fmt.Printf("普通用户%s，基本访问权限\n", user.Name)
	default:
		fmt.Printf("用户%s，未知权限级别\n", user.Name)
	}

	// 函数返回值的switch
	fmt.Println("3. 函数返回值switch:")
	checkNumber := func(n int) string {
		switch {
		case n < 0:
			return "负数"
		case n == 0:
			return "零"
		case n > 0 && n%2 == 0:
			return "正偶数"
		case n > 0 && n%2 == 1:
			return "正奇数"
		default:
			return "未知"
		}
	}

	numbers := []int{-5, 0, 3, 8, 15}
	for _, num := range numbers {
		fmt.Printf("数字%d是: %s\n", num, checkNumber(num))
	}
}

// 7. switch vs if-else 性能对比
func switchVsIfElse() {
	fmt.Println("\n=== switch vs if-else 对比 ===")

	value := 5

	// 使用switch
	fmt.Println("使用switch:")
	switch value {
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三")
	case 4:
		fmt.Println("四")
	case 5:
		fmt.Println("五")
	default:
		fmt.Println("其他")
	}

	// 使用if-else（对比）
	fmt.Println("使用if-else:")
	if value == 1 {
		fmt.Println("一")
	} else if value == 2 {
		fmt.Println("二")
	} else if value == 3 {
		fmt.Println("三")
	} else if value == 4 {
		fmt.Println("四")
	} else if value == 5 {
		fmt.Println("五")
	} else {
		fmt.Println("其他")
	}

	fmt.Println("\n建议:")
	fmt.Println("- 多个固定值比较：使用switch（更清晰）")
	fmt.Println("- 复杂条件判断：使用if-else")
	fmt.Println("- 类型判断：使用type switch")
	fmt.Println("- 范围判断：使用表达式switch")
}

func main() {
	fmt.Println("Go语言控制流 - Switch语句实践")
	fmt.Println("================================")

	// 执行各种switch示例
	basicSwitch()
	expressionSwitch()
	typeSwitch()
	fallthroughSwitch()
	businessSwitch()
	advancedSwitch()
	switchVsIfElse()

	fmt.Println("\n学习要点:")
	fmt.Println("1. Go的switch默认break，不需要手动添加")
	fmt.Println("2. 支持多值匹配：case 1, 2, 3:")
	fmt.Println("3. 表达式switch：switch { case condition: }")
	fmt.Println("4. 类型switch：switch v := value.(type)")
	fmt.Println("5. fallthrough关键字可以穿透到下一个case")
	fmt.Println("6. switch比多个if-else更清晰和高效")
}
