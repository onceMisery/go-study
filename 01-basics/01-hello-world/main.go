package main

import (
	"fmt"
	"os"
)

/*
Go语言Hello World程序
这是你的第一个Go程序，展示了Go的基本语法结构

与Java对比：
1. 包声明更简单：package main vs package com.example
2. 导入语句直接：import "fmt" vs import java.util.*
3. 主函数简洁：func main() vs public static void main(String[] args)
4. 无需类包装：直接在包级别定义函数
*/

func main() {
	// 基础输出 - 类似Java的System.out.println()
	fmt.Println("Hello, World!")
	fmt.Println("欢迎来到Go语言世界！")

	// 变量声明和使用
	demonstrateVariables()

	// 处理命令行参数
	handleCommandLineArgs()

	// 格式化输出
	demonstrateFormatting()
}

// demonstrateVariables 演示Go的变量声明方式
func demonstrateVariables() {
	fmt.Println("\n=== 变量声明演示 ===")

	// 方式1：完整声明（类似Java但类型在后）
	var name string = "Go语言"
	var year int = 2024
	var isAwesome bool = true

	fmt.Printf("语言: %s, 年份: %d, 很棒: %t\n", name, year, isAwesome)

	// 方式2：类型推断
	var version = "1.21"  // 自动推断为string
	var downloads = 1000000  // 自动推断为int

	fmt.Printf("版本: %s, 下载量: %d\n", version, downloads)

	// 方式3：简短声明（Go特有，Java没有）
	creator := "Google"
	popularity := 95.5

	fmt.Printf("创建者: %s, 流行度: %.1f%%\n", creator, popularity)

	/*
	Java等价代码：
	String name = "Go语言";
	int year = 2024;
	boolean isAwesome = true;
	
	String creator = "Google";  // Java没有类型推断的简短声明
	double popularity = 95.5;
	*/
}

// handleCommandLineArgs 处理命令行参数
func handleCommandLineArgs() {
	fmt.Println("\n=== 命令行参数演示 ===")

	// os.Args 类似 Java的 String[] args
	args := os.Args

	fmt.Printf("程序名: %s\n", args[0])
	fmt.Printf("参数个数: %d\n", len(args)-1)

	if len(args) > 1 {
		fmt.Println("传入的参数:")
		for i, arg := range args[1:] {
			fmt.Printf("  参数%d: %s\n", i+1, arg)
		}
	} else {
		fmt.Println("没有传入参数")
		fmt.Println("尝试运行: go run main.go 参数1 参数2")
	}

	/*
	Java等价代码：
	public static void main(String[] args) {
	    System.out.println("参数个数: " + args.length);
	    for (int i = 0; i < args.length; i++) {
	        System.out.println("参数" + (i+1) + ": " + args[i]);
	    }
	}
	*/
}

// demonstrateFormatting 演示各种输出格式
func demonstrateFormatting() {
	fmt.Println("\n=== 格式化输出演示 ===")

	name := "张三"
	age := 25
	salary := 8500.50

	// 不同的格式化方式
	fmt.Printf("姓名: %s, 年龄: %d, 工资: %.2f\n", name, age, salary)
	fmt.Printf("十六进制年龄: %x, 科学计数法工资: %e\n", age, salary)

	// 使用占位符宽度
	fmt.Printf("姓名: %-10s 年龄: %3d 工资: %8.2f\n", name, age, salary)

	// Sprintf 格式化到字符串（类似Java的String.format）
	formatted := fmt.Sprintf("员工信息: %s (%d岁)", name, age)
	fmt.Println(formatted)

	/*
	Java等价代码：
	String name = "张三";
	int age = 25;
	double salary = 8500.50;
	
	System.out.printf("姓名: %s, 年龄: %d, 工资: %.2f%n", name, age, salary);
	String formatted = String.format("员工信息: %s (%d岁)", name, age);
	System.out.println(formatted);
	*/
}

/*
关键学习要点：

1. 包结构简化：
   - Go: package main (可执行程序)
   - Java: package com.company.project

2. 导入机制：
   - Go: import "fmt" (直接包名)
   - Java: import java.util.* (完整包路径)

3. 主函数：
   - Go: func main() (简洁)
   - Java: public static void main(String[] args)

4. 变量声明：
   - Go: var name string 或 name := "value"
   - Java: String name = "value"

5. 无需类包装：
   - Go: 函数可以直接在包级别定义
   - Java: 所有方法必须在类中

6. 内置工具：
   - go fmt: 自动格式化代码
   - go build: 编译程序
   - go run: 直接运行源代码

运行方式：
1. go run main.go
2. go build && ./hello-world
3. go run main.go arg1 arg2 (带参数)
*/ 