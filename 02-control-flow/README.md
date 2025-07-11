# Go语言控制流详解

## 学习目标
- 掌握Go语言的所有控制流语句
- 理解Go与Java控制流的差异
- 学会Go特有的控制流特性
- 掌握错误处理模式

## Java vs Go 控制流对比

| 控制结构 | Java | Go | 主要差异 |
|----------|------|----|----------|
| if语句 | if (condition) | if condition | Go不需要括号 |
| for循环 | for, while, do-while | 只有for | Go用for实现所有循环 |
| switch语句 | 需要break | 自动break | Go不会fallthrough |
| 异常处理 | try-catch-finally | 返回error | 完全不同的错误处理 |

## 1. if 语句

### 基本语法
```go
// Java: if (age >= 18) { ... }
age := 20
if age >= 18 {
    fmt.Println("成年人")
}

// if-else
if age >= 18 {
    fmt.Println("成年人")
} else {
    fmt.Println("未成年人")
}

// if-else if-else
if age < 13 {
    fmt.Println("儿童")
} else if age < 18 {
    fmt.Println("青少年")
} else {
    fmt.Println("成年人")
}
```

### Go特有：初始化语句
```go
// 在if中声明变量，作用域仅限于if块
if score := getScore(); score >= 90 {
    fmt.Println("优秀")
} else if score >= 80 {
    fmt.Println("良好") 
} else {
    fmt.Println("需要改进")
}
// score在这里不可访问
```

## 2. for 循环

Go只有for循环，但可以实现Java的所有循环形式。

### 传统for循环
```go
// Java: for (int i = 0; i < 10; i++) { ... }
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// 多个变量
for i, j := 0, 10; i < j; i, j = i+1, j-1 {
    fmt.Printf("i=%d, j=%d\n", i, j)
}
```

### while形式的for
```go
// Java: while (condition) { ... }
i := 0
for i < 10 {
    fmt.Println(i)
    i++
}

// Java: while (true) { ... } (无限循环)
for {
    // 无限循环
    break // 需要break退出
}
```

### range循环
```go
// 遍历数组/切片
numbers := []int{1, 2, 3, 4, 5}

// Java: for (int num : numbers) { ... }
for index, value := range numbers {
    fmt.Printf("索引：%d，值：%d\n", index, value)
}

// 只要值
for _, value := range numbers {
    fmt.Println(value)
}

// 只要索引
for index := range numbers {
    fmt.Println(index)
}

// 遍历字符串
for index, char := range "Hello" {
    fmt.Printf("索引：%d，字符：%c\n", index, char)
}

// 遍历map
userAges := map[string]int{"Alice": 25, "Bob": 30}
for name, age := range userAges {
    fmt.Printf("%s的年龄是%d\n", name, age)
}
```

## 3. switch 语句

### 基本switch
```go
// Java需要break，Go自动break
day := "Monday"
switch day {
case "Monday":
    fmt.Println("周一")
case "Tuesday":
    fmt.Println("周二")
case "Wednesday", "Thursday", "Friday":
    fmt.Println("工作日")
default:
    fmt.Println("其他")
}
```

### 表达式switch
```go
score := 85
switch {
case score >= 90:
    fmt.Println("A")
case score >= 80:
    fmt.Println("B")
case score >= 70:
    fmt.Println("C")
default:
    fmt.Println("D")
}
```

### 类型switch
```go
var x interface{} = "hello"
switch v := x.(type) {
case string:
    fmt.Printf("字符串：%s\n", v)
case int:
    fmt.Printf("整数：%d\n", v)
case bool:
    fmt.Printf("布尔：%t\n", v)
default:
    fmt.Printf("未知类型：%T\n", v)
}
```

### fallthrough
```go
// 如果需要继续执行下一个case
switch day {
case "Saturday":
    fmt.Println("周六")
    fallthrough
case "Sunday":
    fmt.Println("周末")
}
```

## 4. 跳转语句

### break 和 continue
```go
// break：跳出循环
for i := 0; i < 10; i++ {
    if i == 5 {
        break // 跳出循环
    }
    fmt.Println(i)
}

// continue：跳过当前迭代
for i := 0; i < 10; i++ {
    if i%2 == 0 {
        continue // 跳过偶数
    }
    fmt.Println(i) // 只打印奇数
}
```

### 标签和goto
```go
// 带标签的break/continue
outer:
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        if i == 1 && j == 1 {
            break outer // 跳出外层循环
        }
        fmt.Printf("i=%d, j=%d\n", i, j)
    }
}

// goto (不推荐使用)
func gotoExample() {
    i := 0
loop:
    if i < 3 {
        fmt.Println(i)
        i++
        goto loop
    }
}
```

## 5. defer 语句

defer是Go特有的控制流，类似Java的finally但更灵活。

```go
func fileOperation() {
    file, err := os.Open("data.txt")
    if err != nil {
        return
    }
    defer file.Close() // 函数返回前执行，确保文件关闭
    
    // 使用文件...
}

// 多个defer，LIFO顺序执行
func deferExample() {
    defer fmt.Println("第三个执行")
    defer fmt.Println("第二个执行") 
    defer fmt.Println("第一个执行")
    fmt.Println("正常执行")
}
```

## 6. 错误处理

Go不使用异常，而是通过返回值处理错误。

```go
// Java: try-catch
// try {
//     int result = divide(10, 0);
// } catch (ArithmeticException e) {
//     System.out.println("除零错误");
// }

// Go: 错误返回值
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("不能除零")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println("错误：", err)
        return
    }
    fmt.Println("结果：", result)
}
```

### 错误处理模式

```go
// 1. 立即处理
if err := doSomething(); err != nil {
    log.Fatal(err)
}

// 2. 传播错误
func processData() error {
    data, err := readData()
    if err != nil {
        return fmt.Errorf("处理数据失败: %w", err)
    }
    // 处理data...
    return nil
}

// 3. 忽略错误（谨慎使用）
result, _ := someFunctionThatMightFail()
```

## 7. panic 和 recover

类似Java的异常，但只用于严重错误。

```go
func riskyFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("恢复自panic：", r)
        }
    }()
    
    panic("某些严重错误") // 类似throw new RuntimeException
}
```

## 与Java的主要差异

| 特性 | Java | Go | 优势 |
|------|------|----|------|
| 条件语句括号 | 必须有() | 不需要() | 更简洁 |
| 循环类型 | for、while、do-while | 只有for | 统一、简单 |
| switch fallthrough | 默认fallthrough | 默认break | 减少bug |
| 异常处理 | try-catch | error返回值 | 错误更明确 |
| 资源管理 | try-with-resources | defer | 更灵活 |

## 实践任务

1. **猜数字游戏**：使用for循环、if语句和随机数
2. **学生成绩统计**：使用range循环处理成绩数组
3. **简单计算器**：使用switch语句实现四则运算
4. **文件处理**：使用defer确保文件正确关闭
5. **错误处理练习**：实现一个可能出错的函数

## 常见陷阱

1. **range循环的变量重用**：
```go
// 错误：所有goroutine都会打印最后一个值
for _, item := range items {
    go func() {
        fmt.Println(item) // item被重用
    }()
}

// 正确：传递参数
for _, item := range items {
    go func(i string) {
        fmt.Println(i)
    }(item)
}
```

2. **defer参数立即求值**：
```go
func deferTrap() {
    i := 0
    defer fmt.Println(i) // 打印0，不是5
    i = 5
}
```

3. **switch不需要break**：
```go
// 不会执行多个case
switch x {
case 1:
    fmt.Println("一")
case 2:
    fmt.Println("二") // x=1时不会执行
}
```

## 下一步

掌握了控制流后，建议学习：
- 函数定义和调用
- 数据结构（数组、切片、映射）
- 指针和方法

## 参考资源

- [Go Tour - 控制流](https://tour.golang.org/flowcontrol/1)
- [Effective Go - 控制结构](https://golang.org/doc/effective_go#control-structures)
- [Go by Example - 控制流](https://gobyexample.com/) 