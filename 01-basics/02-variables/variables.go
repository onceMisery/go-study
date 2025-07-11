package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

/*
Go语言变量和数据类型详解
展示Go语言的类型系统、变量声明方式和与Java的对比

学习重点：
1. 变量声明的多种方式
2. Go的类型系统
3. 常量和iota的使用
4. 类型转换
5. 作用域规则
*/

// 包级别变量（全局变量）
var GlobalCounter int = 0
var globalMessage = "这是全局变量"

// 常量定义
const (
	Pi        = 3.14159
	MaxUsers  = 1000
	AppName   = "Go学习程序"
	isEnabled = true
)

// 使用iota的枚举
const (
	Sunday    = iota // 0
	Monday           // 1
	Tuesday          // 2
	Wednesday        // 3
	Thursday         // 4
	Friday           // 5
	Saturday         // 6
)

// 复杂的iota示例
const (
	_  = iota             // 0 被忽略
	KB = 1 << (10 * iota) // 1024
	MB                    // 1048576
	GB                    // 1073741824
)

// HTTP状态码常量
const (
	StatusOK           = 200
	StatusBadRequest   = 400
	StatusUnauthorized = 401
	StatusNotFound     = 404
	StatusServerError  = 500
)

func main() {
	fmt.Println("=== Go语言变量和数据类型学习 ===\n")

	// 演示各种变量声明方式
	demonstrateVariableDeclarations()

	// 演示数据类型
	demonstrateDataTypes()

	// 演示常量和枚举
	demonstrateConstants()

	// 演示类型转换
	demonstrateTypeConversions()

	// 演示作用域
	demonstrateScope()

	// 演示零值
	demonstrateZeroValues()
}

// demonstrateVariableDeclarations 演示变量声明的四种方式
func demonstrateVariableDeclarations() {
	fmt.Println("=== 变量声明方式对比 ===")

	// 方式1：完整声明（类似Java但类型在后）
	var name string = "张三"
	var age int = 25
	var salary float64 = 8500.50

	fmt.Printf("完整声明 - 姓名: %s, 年龄: %d, 工资: %.2f\n", name, age, salary)

	// 方式2：类型推断
	var department = "技术部" // 推断为string
	var experience = 3     // 推断为int
	var bonus = 2000.0     // 推断为float64

	fmt.Printf("类型推断 - 部门: %s, 经验: %d年, 奖金: %.2f\n", department, experience, bonus)

	// 方式3：简短声明（Go特有的便利语法）
	position := "高级工程师"
	isManager := false
	workYears := 5

	fmt.Printf("简短声明 - 职位: %s, 是否管理者: %t, 工作年限: %d\n", position, isManager, workYears)

	// 方式4：批量声明
	var (
		company     string  = "阿里巴巴"
		location    string  = "杭州"
		teamSize    int     = 10
		performance float64 = 95.5
	)

	fmt.Printf("批量声明 - 公司: %s, 地点: %s, 团队规模: %d, 绩效: %.1f%%\n",
		company, location, teamSize, performance)

	// 多变量同时声明
	var x, y, z = 1, 2, 3
	a, b, c := "hello", true, 42

	fmt.Printf("多变量声明 - x=%d, y=%d, z=%d, a=%s, b=%t, c=%d\n", x, y, z, a, b, c)

	/*
		Java等价代码对比：

		// Java方式（类型在前）
		String name = "张三";
		int age = 25;
		double salary = 8500.50;

		// Java没有类型推断的简短声明
		String department = "技术部";
		int experience = 3;

		// Java没有:=这样的简短声明语法
		String position = "高级工程师";
		boolean isManager = false;
	*/

	fmt.Println()
}

// demonstrateDataTypes 演示Go的数据类型
func demonstrateDataTypes() {
	fmt.Println("=== 数据类型详解 ===")

	// 整数类型
	var i8 int8 = 127
	var i16 int16 = 32767
	var i32 int32 = 2147483647
	var i64 int64 = 9223372036854775807

	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615

	fmt.Printf("有符号整数: int8=%d, int16=%d, int32=%d, int64=%d\n", i8, i16, i32, i64)
	fmt.Printf("无符号整数: uint8=%d, uint16=%d, uint32=%d, uint64=%d\n", ui8, ui16, ui32, ui64)

	// 浮点类型
	var f32 float32 = 3.14159
	var f64 float64 = 3.141592653589793

	fmt.Printf("浮点数: float32=%.5f, float64=%.15f\n", f32, f64)

	// 字符串和字符
	var str string = "Hello, 世界"
	var r rune = '中' // rune是int32的别名，用于Unicode码点
	var b byte = 'A' // byte是uint8的别名

	fmt.Printf("字符串: %s, 字符(rune): %c(%d), 字节(byte): %c(%d)\n", str, r, r, b, b)

	// 多行字符串（反引号）
	var multiline string = `这是一个
多行字符串
可以包含"双引号"和'单引号'`

	fmt.Printf("多行字符串:\n%s\n", multiline)

	// 布尔类型
	var isActive bool = true
	var isCompleted bool = false

	fmt.Printf("布尔值: isActive=%t, isCompleted=%t\n", isActive, isCompleted)

	// 类型大小
	fmt.Printf("类型大小: int=%d字节, float64=%d字节, string=%d字节\n",
		unsafe.Sizeof(i32), unsafe.Sizeof(f64), unsafe.Sizeof(str))

	/*
		Java对比：

		// Java的基本类型
		byte b = 127;        // 1字节
		short s = 32767;     // 2字节
		int i = 2147483647;  // 4字节
		long l = 9223372036854775807L; // 8字节

		float f = 3.14f;     // 4字节
		double d = 3.14159;  // 8字节

		char c = 'A';        // 2字节，UTF-16
		boolean flag = true; // JVM实现相关

		String str = "Hello"; // 引用类型，对象在堆上
	*/

	fmt.Println()
}

// demonstrateConstants 演示常量和iota的使用
func demonstrateConstants() {
	fmt.Println("=== 常量和iota枚举 ===")

	// 使用预定义常量
	fmt.Printf("应用信息: %s, 最大用户数: %d, π: %.5f\n", AppName, MaxUsers, Pi)

	// 星期枚举
	fmt.Printf("星期枚举: Sunday=%d, Monday=%d, Saturday=%d\n", Sunday, Monday, Saturday)

	// 存储单位
	fmt.Printf("存储单位: KB=%d, MB=%d, GB=%d\n", KB, MB, GB)

	// HTTP状态码
	fmt.Printf("HTTP状态码: OK=%d, NotFound=%d, ServerError=%d\n",
		StatusOK, StatusNotFound, StatusServerError)

	// 动态使用枚举
	today := Wednesday
	switch today {
	case Monday, Tuesday, Wednesday, Thursday, Friday:
		fmt.Printf("今天是工作日: %d\n", today)
	case Saturday, Sunday:
		fmt.Printf("今天是周末: %d\n", today)
	}

	/*
		Java等价代码：

		// Java常量
		public static final double PI = 3.14159;
		public static final int MAX_USERS = 1000;
		public static final String APP_NAME = "Go学习程序";

		// Java枚举
		public enum Weekday {
		    SUNDAY(0), MONDAY(1), TUESDAY(2), WEDNESDAY(3),
		    THURSDAY(4), FRIDAY(5), SATURDAY(6);

		    private final int value;
		    Weekday(int value) { this.value = value; }
		    public int getValue() { return value; }
		}

		// 使用
		Weekday today = Weekday.WEDNESDAY;
	*/

	fmt.Println()
}

// demonstrateTypeConversions 演示类型转换
func demonstrateTypeConversions() {
	fmt.Println("=== 类型转换 ===")

	// 数值类型转换（必须显式转换）
	var i int = 42
	var f float64 = float64(i) // 显式转换
	var u uint = uint(f)

	fmt.Printf("类型转换: int(%d) -> float64(%.1f) -> uint(%d)\n", i, f, u)

	// 字符串和数值转换
	age := 25
	ageStr := strconv.Itoa(age)
	fmt.Printf("数字转字符串: %d -> \"%s\"\n", age, ageStr)

	numStr := "123"
	num, err := strconv.Atoi(numStr)
	if err != nil {
		fmt.Printf("字符串转数字失败: %v\n", err)
	} else {
		fmt.Printf("字符串转数字: \"%s\" -> %d\n", numStr, num)
	}

	// 其他字符串转换
	floatStr := "3.14159"
	floatVal, err := strconv.ParseFloat(floatStr, 64)
	if err == nil {
		fmt.Printf("字符串转浮点数: \"%s\" -> %.5f\n", floatStr, floatVal)
	}

	boolStr := "true"
	boolVal, err := strconv.ParseBool(boolStr)
	if err == nil {
		fmt.Printf("字符串转布尔值: \"%s\" -> %t\n", boolStr, boolVal)
	}

	// 错误处理示例
	invalidStr := "abc"
	_, err = strconv.Atoi(invalidStr)
	if err != nil {
		fmt.Printf("转换错误示例: \"%s\" -> 错误: %v\n", invalidStr, err)
	}

	/*
		Java对比：

		// Java的类型转换
		int i = 42;
		double f = i;              // 自动转换（向上转型）
		int j = (int) f;           // 强制转换（向下转型）

		// Java的字符串转换
		int age = 25;
		String ageStr = String.valueOf(age);  // 或 Integer.toString(age)

		try {
		    int num = Integer.parseInt("123");
		} catch (NumberFormatException e) {
		    // 异常处理
		}
	*/

	fmt.Println()
}

// demonstrateScope 演示作用域规则
func demonstrateScope() {
	fmt.Println("=== 作用域规则 ===")

	// 包级别变量
	fmt.Printf("全局变量: GlobalCounter=%d, globalMessage=%s\n", GlobalCounter, globalMessage)

	// 函数级别变量
	funcVar := "函数作用域变量"

	// 块级别作用域
	if true {
		blockVar := "块作用域变量"
		fmt.Printf("块内访问: funcVar=%s, blockVar=%s\n", funcVar, blockVar)

		// 变量遮蔽（shadowing）
		funcVar := "被遮蔽的变量"
		fmt.Printf("变量遮蔽: 块内funcVar=%s\n", funcVar)
	}
	// blockVar在这里不可访问

	fmt.Printf("块外访问: funcVar=%s\n", funcVar) // 原来的funcVar

	// 循环中的作用域
	for i := 0; i < 3; i++ {
		loopVar := fmt.Sprintf("循环变量_%d", i)
		fmt.Printf("循环内: i=%d, loopVar=%s\n", i, loopVar)
	}
	// i和loopVar在这里不可访问

	/*
		Java对比：

		public class ScopeExample {
		    // 类级别变量
		    private static String classVar = "类变量";

		    public static void main(String[] args) {
		        // 方法级别变量
		        String methodVar = "方法变量";

		        if (true) {
		            // 块级别变量（Java 10+支持var）
		            String blockVar = "块变量";
		            System.out.println(methodVar + ", " + blockVar);
		        }
		        // blockVar在这里不可访问

		        for (int i = 0; i < 3; i++) {
		            String loopVar = "循环变量_" + i;
		            System.out.println(loopVar);
		        }
		        // i和loopVar在这里不可访问
		    }
		}
	*/

	fmt.Println()
}

// demonstrateZeroValues 演示零值概念
func demonstrateZeroValues() {
	fmt.Println("=== 零值概念 ===")

	// 声明但不初始化的变量会被赋予零值
	var intZero int
	var floatZero float64
	var boolZero bool
	var stringZero string
	var sliceZero []int
	var mapZero map[string]int
	var pointerZero *int

	fmt.Printf("零值示例:\n")
	fmt.Printf("  int零值: %d\n", intZero)
	fmt.Printf("  float64零值: %.1f\n", floatZero)
	fmt.Printf("  bool零值: %t\n", boolZero)
	fmt.Printf("  string零值: \"%s\"\n", stringZero)
	fmt.Printf("  slice零值: %v (nil: %t)\n", sliceZero, sliceZero == nil)
	fmt.Printf("  map零值: %v (nil: %t)\n", mapZero, mapZero == nil)
	fmt.Printf("  pointer零值: %v (nil: %t)\n", pointerZero, pointerZero == nil)

	// 零值的实用性
	if stringZero == "" {
		fmt.Println("字符串零值检查: 字符串为空")
	}

	if sliceZero == nil {
		fmt.Println("切片零值检查: 切片为nil")
	}

	/*
		Java对比：

		// Java的默认值（成员变量）
		class Example {
		    int intDefault;        // 0
		    double doubleDefault;  // 0.0
		    boolean boolDefault;   // false
		    String stringDefault;  // null
		    List<String> listDefault; // null
		}

		// 局部变量必须初始化
		public void method() {
		    int local;  // 编译错误：变量未初始化
		    // System.out.println(local);
		}
	*/

	fmt.Println()
}

/*
Go变量和数据类型总结：

1. 变量声明方式：
   - var name type = value  (完整声明)
   - var name = value       (类型推断)
   - name := value          (简短声明)
   - var (批量声明)

2. 数据类型特点：
   - 类型严格，无自动转换
   - 支持无符号整数类型
   - rune代表Unicode码点
   - byte是uint8的别名

3. 常量和iota：
   - const定义常量
   - iota提供自增枚举
   - 编译时计算

4. 类型转换：
   - 必须显式转换
   - strconv包用于字符串转换
   - 返回error处理失败

5. 作用域：
   - 包级别（全局）
   - 函数级别
   - 块级别
   - 首字母大小写控制可见性

6. 零值：
   - 所有类型都有零值
   - 数值类型为0
   - 布尔类型为false
   - 字符串为空字符串
   - 引用类型为nil

与Java的主要差异：
- Go没有自动装箱拆箱
- Go没有自动类型转换
- Go通过首字母控制可见性
- Go的错误处理更显式
- Go的语法更简洁
*/
