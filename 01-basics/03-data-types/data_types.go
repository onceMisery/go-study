package main

import (
	"fmt"
	"strconv"
)

// Go语言数据类型学习示例
func main() {
	// 1. 基本数据类型演示
	demonstrateBasicTypes()

	// 2. 复合数据类型演示
	demonstrateCompositeTypes()

	// 3. 指针类型演示
	demonstratePointers()

	// 4. 类型转换演示
	demonstrateTypeConversions()
}

// demonstrateBasicTypes 演示基本数据类型
func demonstrateBasicTypes() {
	fmt.Println("=== 基本数据类型 ===")

	// 整数类型
	var age int = 25
	var smallNum int8 = 127
	var hugeNum int64 = 9223372036854775807
	var size uint = 100

	fmt.Printf("整数类型: age=%d, smallNum=%d, hugeNum=%d, size=%d\n", age, smallNum, hugeNum, size)

	// 浮点数类型
	var price float64 = 19.99
	var bigFloat float64 = 3.14159265359

	fmt.Printf("浮点数类型: price=%.2f, bigFloat=%.10f\n", price, bigFloat)

	// 布尔类型
	var isActive bool = true
	fmt.Printf("布尔类型: isActive=%t\n", isActive)

	// 字符和字符串
	var c rune = 'A'
	var s string = "Hello"
	var chinese string = "你好世界"

	fmt.Printf("字符和字符串: c=%c, s=%s, chinese=%s\n", c, s, chinese)
}

// demonstrateCompositeTypes 演示复合数据类型
func demonstrateCompositeTypes() {
	fmt.Println("\n=== 复合数据类型 ===")

	// 数组
	var numbers [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("数组: %v\n", numbers)

	// 切片
	slice := []int{1, 2, 3, 4, 5}
	slice = append(slice, 6, 7, 8)
	fmt.Printf("切片: %v\n", slice)

	// 映射
	userAges := map[string]int{
		"Alice": 25,
		"Bob":   30,
	}
	fmt.Printf("映射: %v\n", userAges)

	// 结构体
	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "Alice", Age: 25}
	fmt.Printf("结构体: %v\n", person)
}

// demonstratePointers 演示指针类型
func demonstratePointers() {
	fmt.Println("\n=== 指针类型 ===")

	var x int = 42
	// 定义指向 x 的指针变量 ptr
	var ptr *int = &x

	fmt.Printf("指针地址: %p\n", ptr)
	fmt.Printf("指针值: %d\n", *ptr)
	// 通过指针修改 x 的值为 100，并打印修改后的值
	*ptr = 100
	fmt.Printf("修改后的值: %d\n", x)
}

// demonstrateTypeConversions 演示类型转换
func demonstrateTypeConversions() {
	fmt.Println("\n=== 类型转换 ===")

	var f float64 = 3.14
	var i int = int(f)
	fmt.Printf("float到int转换: %v -> %d\n", f, i)

	// 字符串转换
	var num int = 42
	var str string = strconv.Itoa(num)
	fmt.Printf("整数转字符串: %d -> \"%s\"\n", num, str)

	numStr := "123"
	num2, err := strconv.Atoi(numStr)
	if err == nil {
		fmt.Printf("字符串转整数: \"%s\" -> %d\n", numStr, num2)
	}
}
