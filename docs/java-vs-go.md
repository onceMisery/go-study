# Java与Go语言详细对比

## 语言设计哲学

### Java
- 面向对象为核心
- "一次编写，到处运行"
- 企业级应用导向
- 丰富的生态系统

### Go
- 简洁性和效率为核心
- 快速编译，高效执行
- 云原生和微服务导向
- 内置并发支持

## 语法对比

### 1. 变量声明

**Java:**
```java
// 强类型，显式声明
int age = 25;
String name = "John";
List<String> names = new ArrayList<>();
```

**Go:**
```go
// 类型推断 + 显式类型
var age int = 25
var name string = "John"
var names []string

// 简短声明（类型推断）
age := 25
name := "John"
names := []string{}
```

### 2. 函数定义

**Java:**
```java
public class Calculator {
    public static int add(int a, int b) {
        return a + b;
    }
    
    // 多返回值需要包装类
    public Result divide(int a, int b) {
        if (b == 0) {
            throw new IllegalArgumentException("除数不能为0");
        }
        return new Result(a / b, null);
    }
}
```

**Go:**
```go
// 函数可以在包级别定义
func add(a, b int) int {
    return a + b
}

// 原生支持多返回值
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("除数不能为0")
    }
    return a / b, nil
}
```

### 3. 错误处理

**Java:**
```java
try {
    result = divide(10, 0);
} catch (IllegalArgumentException e) {
    System.err.println("错误: " + e.getMessage());
}
```

**Go:**
```go
result, err := divide(10, 0)
if err != nil {
    fmt.Printf("错误: %s\n", err.Error())
    return
}
```

### 4. 面向对象

**Java:**
```java
public class Animal {
    protected String name;
    
    public Animal(String name) {
        this.name = name;
    }
    
    public void speak() {
        System.out.println(name + " makes a sound");
    }
}

public class Dog extends Animal {
    public Dog(String name) {
        super(name);
    }
    
    @Override
    public void speak() {
        System.out.println(name + " barks");
    }
}
```

**Go:**
```go
// 结构体 + 方法
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Printf("%s makes a sound\n", a.Name)
}

type Dog struct {
    Animal  // 组合而非继承
}

func (d Dog) Speak() {
    fmt.Printf("%s barks\n", d.Name)
}
```

### 5. 接口

**Java:**
```java
public interface Speaker {
    void speak();
}

public class Dog implements Speaker {
    public void speak() {
        System.out.println("Dog barks");
    }
}
```

**Go:**
```go
// 接口自动实现（鸭子类型）
type Speaker interface {
    Speak()
}

type Dog struct{}

func (d Dog) Speak() {
    fmt.Println("Dog barks")
}
// Dog自动实现了Speaker接口
```

### 6. 并发编程

**Java:**
```java
// 线程和线程池
ExecutorService executor = Executors.newFixedThreadPool(10);

executor.submit(() -> {
    System.out.println("任务执行在线程: " + Thread.currentThread().getName());
});

// CompletableFuture
CompletableFuture<String> future = CompletableFuture.supplyAsync(() -> {
    return "异步结果";
});
```

**Go:**
```go
// goroutine（轻量级线程）
go func() {
    fmt.Println("任务执行在goroutine")
}()

// channel通信
ch := make(chan string)
go func() {
    ch <- "异步结果"
}()
result := <-ch
```

## 生态系统对比

### Java生态
| 领域 | 主要框架/库 |
|------|------------|
| Web开发 | Spring Boot, Spring MVC |
| ORM | Hibernate, JPA, MyBatis |
| 依赖注入 | Spring DI, Guice |
| 测试 | JUnit, Mockito, TestNG |
| 构建工具 | Maven, Gradle |
| 微服务 | Spring Cloud, Dubbo |

### Go生态
| 领域 | 主要框架/库 |
|------|------------|
| Web开发 | Gin, Echo, Fiber |
| ORM | GORM, Ent |
| 依赖注入 | Wire, Dig |
| 测试 | testing, Testify |
| 构建工具 | go build, go mod |
| 微服务 | Go-kit, Kratos |

## 性能对比

### 编译和启动
- **Java**: 编译到字节码，JVM启动较慢，热身后性能优秀
- **Go**: 编译到机器码，启动极快，性能稳定

### 内存使用
- **Java**: JVM堆内存，GC可配置但复杂
- **Go**: 直接内存管理，GC简单高效

### 并发性能
- **Java**: 线程较重，需要线程池管理
- **Go**: goroutine极轻量，可创建数百万个

## 适用场景

### Java适合
- 企业级应用开发
- 复杂业务逻辑系统
- 需要丰富生态支持的项目
- 团队已有Java经验的项目

### Go适合
- 云原生应用
- 微服务架构
- 高并发系统
- DevOps工具开发
- 网络编程

## 学习建议

### 从Java转Go的关键点
1. **思维转换**: 从继承转向组合
2. **错误处理**: 适应显式错误处理
3. **并发模型**: 掌握goroutine和channel
4. **接口设计**: 理解隐式接口实现
5. **内存模型**: 理解值类型和指针 