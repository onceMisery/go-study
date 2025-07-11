# Variables and Data Types - 变量和数据类型

## 📋 学习目标

- 掌握Go语言的变量声明方式
- 理解Go的基本数据类型
- 学会常量定义和iota的使用
- 掌握类型转换和类型断言
- 理解作用域规则

## 🔍 Java vs Go 变量声明对比

### Java变量声明

```java
// 基本类型
int age = 25;
double salary = 8500.50;
boolean isActive = true;
char grade = 'A';

// 引用类型
String name = "张三";
List<String> hobbies = new ArrayList<>();

// 常量
final int MAX_SIZE = 100;
final String COMPANY = "阿里巴巴";
```

### Go变量声明

```go
// 完整声明
var age int = 25
var salary float64 = 8500.50
var isActive bool = true

// 类型推断
var name = "张三"
var hobbies = []string{}

// 简短声明（Go特有）
age := 25
name := "张三"

// 常量
const MaxSize = 100
const Company = "阿里巴巴"
```

## 📊 数据类型详解

### 1. 数值类型

| Go类型    | Java类型 | 大小   | 范围               |
|---------|--------|------|------------------|
| int8    | byte   | 1字节  | -128 ~ 127       |
| int16   | short  | 2字节  | -32,768 ~ 32,767 |
| int32   | int    | 4字节  | -2³¹ ~ 2³¹-1     |
| int64   | long   | 8字节  | -2⁶³ ~ 2⁶³-1     |
| int     | int    | 平台相关 | 32位或64位          |
| uint8   | -      | 1字节  | 0 ~ 255          |
| uint16  | -      | 2字节  | 0 ~ 65,535       |
| uint32  | -      | 4字节  | 0 ~ 2³²-1        |
| uint64  | -      | 8字节  | 0 ~ 2⁶⁴-1        |
| float32 | float  | 4字节  | IEEE-754         |
| float64 | double | 8字节  | IEEE-754         |

### 2. 字符串和字符

**Java:**

```java
String str = "Hello";
char c = 'A';
String multiline = """
    多行字符串
    第二行
    """;
```

**Go:**

```go
var str string = "Hello"
var c rune = 'A' // rune是int32的别名，用于Unicode码点
var multiline string = `
多行字符串
第二行
`
```

### 3. 布尔类型

**Java:**

```java
boolean flag = true;
boolean result = (age > 18) && isActive;
```

**Go:**

```go
var flag bool = true
var result bool = (age > 18) && isActive
```

## 🎯 变量声明的四种方式

### 1. 完整声明

```go
var name string = "Go语言"
var age int = 5
var version float64 = 1.21
```

### 2. 类型推断

```go
var name = "Go语言" // 推断为string
var age = 5         // 推断为int
var version = 1.21 // 推断为float64
```

### 3. 简短声明

```go
name := "Go语言"
age := 5
version := 1.21
```

### 4. 批量声明

```go
var (
name    string = "Go语言"
age     int = 5
version float64 = 1.21
)

// 或者
var name, age, version = "Go语言", 5, 1.21
```

## 🔒 常量定义

### 基本常量

```go
const Pi = 3.14159
const Company = "Google"
const MaxRetries = 3

// 批量定义
const (
StatusOK = 200
StatusNotFound = 404
StatusError = 500
)
```

### iota 枚举器

```go
const (
Sunday = iota // 0
Monday        // 1
Tuesday       // 2
Wednesday        // 3
Thursday         // 4
Friday           // 5
Saturday         // 6
)

// 复杂的iota使用
const (
_ = iota // 0, 被忽略
KB = 1 << (10 * iota) // 1024
MB                    // 1048576
GB                    // 1073741824
)
```

**Java等价代码:**

```java
public enum Weekday {
    SUNDAY(0),
    MONDAY(1),
    TUESDAY(2),
    WEDNESDAY(3),
    THURSDAY(4),
    FRIDAY(5),
    SATURDAY(6);
    
    private final int value;
    Weekday(int value) { this.value = value; }
}
```

## 🔄 类型转换

### 基本类型转换

```go
var i int = 42
var f float64 = float64(i) // 显式转换
var u uint = uint(f)

// Java中的自动装箱拆箱在Go中不存在
// 必须显式转换
```

**Java对比:**

```java
int i = 42;
double f = i;        // 自动转换
float ff = (float)f; // 强制转换
```

### 字符串转换

```go
import "strconv"

// 数字转字符串
age := 25
ageStr := strconv.Itoa(age)

// 字符串转数字
str := "123"
num, err := strconv.Atoi(str)
if err != nil {
// 处理错误
}
```

**Java对比:**

```java
int age = 25;
String ageStr = String.valueOf(age);

String str = "123";
int num = Integer.parseInt(str); // 可能抛出异常
```

## 🎯 作用域规则

### 包级别作用域

```go
package main

var globalVar = "全局变量" // 包级别

func main() {
	var localVar = "局部变量" // 函数级别

	if true {
		var blockVar = "块级别" // 块级别
		fmt.Println(globalVar, localVar, blockVar)
	}
	// blockVar在这里不可访问
}
```

### 可见性规则

```go
var PublicVar = "公开的"  // 首字母大写，包外可见
var privateVar = "私有的" // 首字母小写，包内可见

func PublicFunction() {}   // 公开函数
func privateFunction() {} // 私有函数
```

**Java对比:**

```java
public class Example {
    public static String publicVar = "公开的";
    private static String privateVar = "私有的";
    
    public static void publicMethod() {}
    private static void privateMethod() {}
}
```

## 💡 零值概念

Go中的所有类型都有零值，声明但未初始化的变量会被赋予零值：

```go
var i int     // 0
var f float64 // 0.0
var b bool       // false
var s string     // ""
var p *int       // nil
var slice []int // nil
var m map[string]int // nil
```

**Java对比:**

```java
// Java中基本类型有默认值，引用类型为null
int i;           // 0
boolean b;       // false
String s;        // null
```

## 📝 实践任务

### 任务1: 变量声明练习

1. 使用四种不同方式声明变量
2. 练习类型推断
3. 对比Java的声明方式

### 任务2: 常量和iota

1. 定义业务常量
2. 使用iota创建枚举
3. 对比Java的枚举实现

### 任务3: 类型转换

1. 实现数字类型转换
2. 练习字符串转换
3. 处理转换错误

## 🎯 学习要点

### Go的优势

1. **类型推断**: 减少冗余代码
2. **简短声明**: `:=` 语法简洁
3. **零值概念**: 所有类型都有合理默认值
4. **iota**: 方便的枚举生成

### 需要注意的差异

1. **无自动类型转换**: 必须显式转换
2. **无装箱拆箱**: 基本类型就是基本类型
3. **可见性规则**: 通过首字母大小写控制
4. **错误处理**: 类型转换可能返回错误

## 🎯 下一步

- 学习运算符和表达式
- 理解控制流语句
- 掌握函数定义和调用

## 📚 参考文档

- [Go语言规范 - 变量](https://golang.org/ref/spec#Variables)
- [Go语言规范 - 常量](https://golang.org/ref/spec#Constants)
- [Effective Go - 变量](https://golang.org/doc/effective_go.html#variables) 