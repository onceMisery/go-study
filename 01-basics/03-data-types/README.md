# Go语言数据类型详解

## 学习目标
- 掌握Go语言的基本数据类型
- 理解Go与Java数据类型的差异
- 学会复合数据类型的使用
- 掌握类型转换和类型断言

## Java vs Go 数据类型对比

| 类型分类 | Java | Go | 主要差异 |
|---------|------|----|---------| 
| 整数 | byte, short, int, long | int8, int16, int32, int64, int, uint | Go有无符号类型，int大小依赖平台 |
| 浮点数 | float, double | float32, float64 | Go没有默认浮点类型 |
| 字符 | char (UTF-16) | rune (UTF-8) | Go原生支持UTF-8 |
| 布尔 | boolean | bool | 类似 |
| 字符串 | String (对象) | string (值类型) | Go字符串不可变但是值类型 |
| 数组 | 引用类型 | 值类型 | 完全不同的内存模型 |

## 1. 基本数据类型

### 整数类型
```go
// Java: int age = 25;
var age int = 25

// Go支持不同大小的整数
var smallNum int8 = 127      // -128 到 127
var mediumNum int16 = 32767  // -32768 到 32767
var bigNum int32 = 2147483647
var hugeNum int64 = 9223372036854775807

// 无符号整数 (Java没有)
var positiveOnly uint8 = 255  // 0 到 255
var size uint = 100          // 平台相关
```

### 浮点数类型
```go
// Java: double price = 19.99;
var price float64 = 19.99

// Go需要明确指定精度
var smallFloat float32 = 3.14  // 单精度
var bigFloat float64 = 3.14159265359  // 双精度
```

### 布尔类型
```go
// Java: boolean isActive = true;
var isActive bool = true

// Go的bool只有true和false，不能用数字
// 这与Java相同，但Go更严格
```

### 字符和字符串
```go
// Java: char c = 'A'; String s = "Hello";
var c rune = 'A'          // rune是int32的别名，表示Unicode码点
var s string = "Hello"    // 字符串是不可变的

// Go字符串是UTF-8编码的字节序列
var chinese string = "你好世界"
```

## 2. 复合数据类型

### 数组 (固定长度)
```go
// Java: int[] numbers = {1, 2, 3, 4, 5};
var numbers [5]int = [5]int{1, 2, 3, 4, 5}

// 简化写法
numbers2 := [5]int{1, 2, 3, 4, 5}
numbers3 := [...]int{1, 2, 3, 4, 5}  // 自动推断长度
```

### 切片 (动态数组)
```go
// 类似Java的ArrayList
var slice []int = []int{1, 2, 3, 4, 5}
slice2 := []int{1, 2, 3, 4, 5}

// 动态添加元素
slice = append(slice, 6, 7, 8)
```

### 映射 (Map)
```go
// Java: Map<String, Integer> map = new HashMap<>();
var userAges map[string]int = make(map[string]int)
userAges["Alice"] = 25
userAges["Bob"] = 30

// 简化写法
userAges2 := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
```

### 结构体
```go
// 替代Java的class
type Person struct {
    Name string
    Age  int
}

// Java: Person person = new Person("Alice", 25);
person := Person{Name: "Alice", Age: 25}
```

## 3. 指针类型

```go
// Java没有显式指针，Go有
var x int = 42
var ptr *int = &x  // ptr指向x的地址

fmt.Println(*ptr)  // 42，解引用
*ptr = 100         // 修改x的值
fmt.Println(x)     // 100
```

## 4. 接口类型

```go
// 类似Java接口，但实现是隐式的
type Writer interface {
    Write([]byte) (int, error)
}

// 任何有Write方法的类型都实现了Writer接口
```

## 5. 函数类型

```go
// Go中函数是一等公民
type Calculator func(int, int) int

func add(a, b int) int {
    return a + b
}

var calc Calculator = add
result := calc(5, 3)  // 8
```

## 6. 类型转换

### 显式转换
```go
// Java: int i = (int) 3.14;
var f float64 = 3.14
var i int = int(f)  // Go要求显式转换

// 字符串转换
var num int = 42
var str string = strconv.Itoa(num)  // 整数转字符串
var back int, err = strconv.Atoi(str)  // 字符串转整数
```

### 类型断言
```go
// 类似Java的instanceof和强制转换
var i interface{} = "hello"

// 类型断言
str, ok := i.(string)
if ok {
    fmt.Println("i is a string:", str)
}

// 类型切换
switch v := i.(type) {
case string:
    fmt.Println("i is string:", v)
case int:
    fmt.Println("i is int:", v)
default:
    fmt.Println("unknown type")
}
```

## 7. 零值概念

```go
// Java: Object obj = null;
// Go的每种类型都有零值，不是null/nil

var i int        // 0
var f float64    // 0.0
var b bool       // false
var s string     // ""
var ptr *int     // nil
var slice []int  // nil
var m map[string]int  // nil
```

## 与Java的主要差异

| 特性 | Java | Go | 说明 |
|------|------|----|------|
| 包装类型 | Integer, Double等 | 无 | Go只有基本类型 |
| null处理 | 可能NullPointerException | 零值设计 | Go更安全 |
| 数组 | 引用类型 | 值类型 | 内存模型不同 |
| 字符串 | 对象 | 值类型 | Go字符串更高效 |
| 泛型 | 完整支持 | 1.18+支持 | Go泛型相对简单 |
| 自动装箱 | 支持 | 不支持 | Go类型系统更严格 |

## 实践任务

1. 创建一个程序，演示所有基本数据类型的使用
2. 实现一个简单的学生管理系统，使用结构体和映射
3. 练习类型转换和类型断言
4. 对比数组和切片的区别

## 常见陷阱

1. **数组是值类型**：传递数组会复制整个数组
2. **切片和数组不同**：切片是引用类型
3. **字符串不可变**：修改字符串会创建新字符串
4. **类型转换必须显式**：Go不会自动转换类型

## 下一步

掌握了数据类型后，建议学习：
- 运算符和表达式
- 控制流语句
- 函数定义和调用

## 参考资源

- [Go语言规范 - 类型](https://golang.org/ref/spec#Types)
- [Effective Go - 数据](https://golang.org/doc/effective_go#data)
- [Go by Example - 基本类型](https://gobyexample.com/) 