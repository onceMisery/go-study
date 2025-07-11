# Functions - 函数

## 📋 学习目标
- 掌握Go语言函数的定义和调用
- 理解多返回值特性
- 学会错误处理模式
- 掌握可变参数和匿名函数
- 理解闭包和defer语句

## 🎯 函数基础

### Java vs Go 函数定义对比

**Java方法定义:**
```java
public class Calculator {
    // 静态方法
    public static int add(int a, int b) {
        return a + b;
    }
    
    // 实例方法
    public int multiply(int a, int b) {
        return a * b;
    }
    
    // 方法重载
    public int add(int a, int b, int c) {
        return a + b + c;
    }
}
```

**Go函数定义:**
```go
package main

// 包级别函数（类似Java静态方法）
func add(a, b int) int {
    return a + b
}

// 多参数相同类型简写
func multiply(a, b int) int {
    return a * b
}

// Go没有方法重载，需要不同名称
func addThree(a, b, c int) int {
    return a + b + c
}
```

## 🔄 多返回值特性

### Go的多返回值
```go
// 返回结果和错误
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为0")
    }
    return a / b, nil
}

// 命名返回值
func swap(x, y string) (first, second string) {
    first = y
    second = x
    return  // 裸返回
}

// 多个返回值的解包
result, err := divide(10, 2)
if err != nil {
    fmt.Printf("错误: %v\n", err)
    return
}
fmt.Printf("结果: %.2f\n", result)
```

**Java对比 - 需要包装类:**
```java
// Java需要创建包装类或使用数组
public class DivideResult {
    private double result;
    private String error;
    
    public DivideResult(double result, String error) {
        this.result = result;
        this.error = error;
    }
    // getters...
}

public static DivideResult divide(double a, double b) {
    if (b == 0) {
        return new DivideResult(0, "除数不能为0");
    }
    return new DivideResult(a / b, null);
}
```

## ⚡ 可变参数

### Go可变参数
```go
// 可变参数函数
func sum(numbers ...int) int {
    total := 0
    for _, num := range numbers {
        total += num
    }
    return total
}

// 调用方式
fmt.Println(sum(1, 2, 3))         // 6
fmt.Println(sum(1, 2, 3, 4, 5))   // 15

// 传递切片
nums := []int{1, 2, 3, 4}
fmt.Println(sum(nums...))          // 10
```

**Java对比:**
```java
// Java可变参数
public static int sum(int... numbers) {
    int total = 0;
    for (int num : numbers) {
        total += num;
    }
    return total;
}

// 调用
System.out.println(sum(1, 2, 3));
int[] nums = {1, 2, 3, 4};
System.out.println(sum(nums));
```

## 🎭 匿名函数和闭包

### Go的匿名函数
```go
// 匿名函数
func main() {
    // 立即执行的匿名函数
    result := func(a, b int) int {
        return a + b
    }(3, 4)
    fmt.Println(result) // 7
    
    // 赋值给变量
    add := func(a, b int) int {
        return a + b
    }
    fmt.Println(add(5, 6)) // 11
}

// 闭包示例
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c1 := counter()
    c2 := counter()
    
    fmt.Println(c1()) // 1
    fmt.Println(c1()) // 2
    fmt.Println(c2()) // 1
}
```

**Java对比 - Lambda表达式:**
```java
// Java 8+ Lambda表达式
public static void main(String[] args) {
    // 匿名函数
    BinaryOperator<Integer> add = (a, b) -> a + b;
    System.out.println(add.apply(5, 6)); // 11
    
    // 闭包效果（使用类）
    Supplier<Integer> counter = createCounter();
    System.out.println(counter.get()); // 1
    System.out.println(counter.get()); // 2
}

public static Supplier<Integer> createCounter() {
    AtomicInteger count = new AtomicInteger(0);
    return () -> count.incrementAndGet();
}
```

## ⏰ defer语句

### defer的使用
```go
func readFile(filename string) error {
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()  // 函数返回前执行
    
    // 读取文件操作
    // ...
    
    return nil
}

// 多个defer按LIFO顺序执行
func deferExample() {
    defer fmt.Println("1")
    defer fmt.Println("2") 
    defer fmt.Println("3")
    fmt.Println("函数体")
}
// 输出：函数体 -> 3 -> 2 -> 1
```

**Java对比 - try-with-resources:**
```java
// Java的资源管理
public void readFile(String filename) throws IOException {
    try (FileInputStream file = new FileInputStream(filename)) {
        // 读取文件操作
        // ...
    } // 自动关闭资源
}

// 或传统try-finally
public void readFileOld(String filename) throws IOException {
    FileInputStream file = null;
    try {
        file = new FileInputStream(filename);
        // 读取文件操作
    } finally {
        if (file != null) {
            file.close();
        }
    }
}
```

## 🔧 函数作为值

### 函数类型和传递
```go
// 定义函数类型
type Operation func(int, int) int

// 接受函数作为参数
func calculate(a, b int, op Operation) int {
    return op(a, b)
}

func main() {
    add := func(a, b int) int { return a + b }
    multiply := func(a, b int) int { return a * b }
    
    fmt.Println(calculate(3, 4, add))      // 7
    fmt.Println(calculate(3, 4, multiply)) // 12
}

// 返回函数
func getOperation(opType string) Operation {
    switch opType {
    case "add":
        return func(a, b int) int { return a + b }
    case "multiply":
        return func(a, b int) int { return a * b }
    default:
        return nil
    }
}
```

**Java对比 - 函数式接口:**
```java
// Java函数式编程
@FunctionalInterface
interface Operation {
    int apply(int a, int b);
}

public static int calculate(int a, int b, Operation op) {
    return op.apply(a, b);
}

public static void main(String[] args) {
    Operation add = (a, b) -> a + b;
    Operation multiply = (a, b) -> a * b;
    
    System.out.println(calculate(3, 4, add));      // 7
    System.out.println(calculate(3, 4, multiply)); // 12
}
```

## 🚨 错误处理模式

### Go的错误处理
```go
import (
    "errors"
    "fmt"
)

// 基础错误处理
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 自定义错误类型
type MathError struct {
    Op    string
    Value float64
    Msg   string
}

func (e MathError) Error() string {
    return fmt.Sprintf("%s %.2f: %s", e.Op, e.Value, e.Msg)
}

func sqrt(x float64) (float64, error) {
    if x < 0 {
        return 0, MathError{"sqrt", x, "负数不能开平方"}
    }
    return math.Sqrt(x), nil
}

// 错误处理链
func processData() error {
    result, err := divide(10, 0)
    if err != nil {
        return fmt.Errorf("处理数据失败: %w", err)
    }
    
    // 使用result...
    return nil
}
```

**Java对比 - 异常处理:**
```java
// Java异常处理
public static double divide(double a, double b) throws ArithmeticException {
    if (b == 0) {
        throw new ArithmeticException("Division by zero");
    }
    return a / b;
}

// 自定义异常
class MathException extends Exception {
    public MathException(String message) {
        super(message);
    }
}

public static double sqrt(double x) throws MathException {
    if (x < 0) {
        throw new MathException("负数不能开平方");
    }
    return Math.sqrt(x);
}

// 异常处理
public static void processData() {
    try {
        double result = divide(10, 0);
        // 使用result...
    } catch (ArithmeticException e) {
        System.err.println("处理数据失败: " + e.getMessage());
    }
}
```

## 📝 实践任务

### 任务1: 基础函数
1. 实现数学计算库
2. 练习多返回值
3. 对比Java方法定义

### 任务2: 高级特性
1. 实现函数式编程示例
2. 练习闭包和匿名函数
3. 掌握defer的使用场景

### 任务3: 错误处理
1. 设计错误处理策略
2. 实现自定义错误类型
3. 对比Java异常机制

## 🎯 学习要点

### Go函数特点
1. **多返回值**: 原生支持，无需包装类
2. **defer语句**: 优雅的资源管理
3. **函数是一等公民**: 可作为值传递
4. **显式错误处理**: 返回error而非抛出异常

### 与Java的主要差异
1. **错误处理**: error返回值 vs 异常抛出
2. **方法重载**: Go不支持重载
3. **资源管理**: defer vs try-with-resources
4. **函数定义**: 包级别函数 vs 类方法

## 🎯 下一步
- 学习数据结构（数组、切片、映射）
- 理解接口和多态
- 掌握并发编程

## 📚 参考资源
- [Go语言规范 - 函数](https://golang.org/ref/spec#Function_declarations)
- [Effective Go - 函数](https://golang.org/doc/effective_go.html#functions)
- [Go by Example - 函数](https://gobyexample.com/functions) 