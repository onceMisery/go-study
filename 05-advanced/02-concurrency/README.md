# Go语言并发编程详解

## 学习目标
- 掌握Go的并发模型和原理
- 理解Go与Java并发编程的差异
- 学会使用goroutine和channel
- 掌握并发安全和同步机制

## Java vs Go 并发对比

| 特性 | Java | Go | 主要差异 |
|------|------|----|----------|
| 基本单位 | Thread | Goroutine | Goroutine更轻量级 |
| 通信方式 | 共享内存 | Channel | Go推荐通过通信共享内存 |
| 创建成本 | 高（~2MB栈） | 低（~2KB栈） | Go可创建数百万个goroutine |
| 同步机制 | synchronized, Lock | Mutex, Channel | Go提供更多选择 |
| 线程池 | 手动管理 | 运行时自动调度 | Go更简单 |

## 1. Goroutine 基础

Goroutine是Go的轻量级线程，由Go运行时管理。

### 创建Goroutine
```go
// Java: new Thread(() -> { ... }).start();
// Go: 只需要在函数调用前加go关键字

func sayHello() {
    fmt.Println("Hello from goroutine!")
}

func main() {
    // 启动goroutine
    go sayHello()
    
    // 启动匿名函数goroutine
    go func() {
        fmt.Println("匿名goroutine")
    }()
    
    // 启动带参数的goroutine
    go func(name string) {
        fmt.Printf("Hello, %s!\n", name)
    }("Go")
    
    // 主goroutine等待
    time.Sleep(time.Second)
}
```

### Goroutine调度原理
```go
// M:N调度模型
// M个系统线程 对应 N个goroutine
// Go运行时的调度器负责在系统线程上调度goroutine

func main() {
    // 设置最大系统线程数
    runtime.GOMAXPROCS(4)
    
    // 获取当前goroutine数量
    fmt.Printf("Goroutines: %d\n", runtime.NumGoroutine())
    
    // 启动1000个goroutine
    for i := 0; i < 1000; i++ {
        go func(id int) {
            time.Sleep(time.Millisecond * 100)
            fmt.Printf("Goroutine %d completed\n", id)
        }(i)
    }
    
    fmt.Printf("Goroutines after launch: %d\n", runtime.NumGoroutine())
    time.Sleep(time.Second * 2)
}
```

## 2. Channel 基础

Channel是Go并发编程的核心，用于goroutine间通信。

### 无缓冲Channel
```go
// Java中需要使用BlockingQueue等
// Go的channel是内置的

func main() {
    // 创建无缓冲channel
    ch := make(chan string)
    
    // 启动发送goroutine
    go func() {
        ch <- "Hello, Channel!"  // 发送数据，会阻塞直到有接收者
    }()
    
    // 接收数据
    message := <-ch  // 接收数据，会阻塞直到有数据
    fmt.Println(message)
}
```

### 有缓冲Channel
```go
func main() {
    // 创建带缓冲的channel
    ch := make(chan int, 3)  // 缓冲区大小为3
    
    // 发送数据（不会阻塞，直到缓冲区满）
    ch <- 1
    ch <- 2
    ch <- 3
    
    // 接收数据
    fmt.Println(<-ch)  // 1
    fmt.Println(<-ch)  // 2
    fmt.Println(<-ch)  // 3
}
```

### Channel关闭和检测
```go
func main() {
    ch := make(chan int, 2)
    
    // 发送数据
    go func() {
        for i := 1; i <= 5; i++ {
            ch <- i
        }
        close(ch)  // 关闭channel
    }()
    
    // 接收数据
    for {
        value, ok := <-ch
        if !ok {
            fmt.Println("Channel closed")
            break
        }
        fmt.Printf("Received: %d\n", value)
    }
    
    // 使用range接收（推荐）
    ch2 := make(chan int, 2)
    go func() {
        for i := 1; i <= 3; i++ {
            ch2 <- i
        }
        close(ch2)
    }()
    
    for value := range ch2 {
        fmt.Printf("Range received: %d\n", value)
    }
}
```

## 3. Channel模式

### 生产者-消费者模式
```go
// Java需要使用BlockingQueue
// Go使用channel更简洁

func producer(ch chan<- int) {  // 只发送channel
    for i := 1; i <= 10; i++ {
        ch <- i
        fmt.Printf("Produced: %d\n", i)
        time.Sleep(time.Millisecond * 100)
    }
    close(ch)
}

func consumer(ch <-chan int) {  // 只接收channel
    for value := range ch {
        fmt.Printf("Consumed: %d\n", value)
        time.Sleep(time.Millisecond * 200)
    }
}

func main() {
    ch := make(chan int, 5)  // 缓冲channel
    
    go producer(ch)
    go consumer(ch)
    
    time.Sleep(time.Second * 3)
}
```

### 扇入扇出模式
```go
// 扇出：一个输入，多个输出
func fanOut(input <-chan int) (<-chan int, <-chan int) {
    out1 := make(chan int)
    out2 := make(chan int)
    
    go func() {
        defer close(out1)
        defer close(out2)
        
        for value := range input {
            out1 <- value
            out2 <- value
        }
    }()
    
    return out1, out2
}

// 扇入：多个输入，一个输出
func fanIn(input1, input2 <-chan int) <-chan int {
    output := make(chan int)
    
    go func() {
        defer close(output)
        for {
            select {
            case value, ok := <-input1:
                if !ok {
                    input1 = nil
                    continue
                }
                output <- value
            case value, ok := <-input2:
                if !ok {
                    input2 = nil
                    continue
                }
                output <- value
            }
            if input1 == nil && input2 == nil {
                break
            }
        }
    }()
    
    return output
}
```

## 4. Select语句

select是Go特有的多路复用机制，类似switch但用于channel操作。

### 基本select
```go
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(time.Second)
        ch1 <- "from ch1"
    }()
    
    go func() {
        time.Sleep(time.Second * 2)
        ch2 <- "from ch2"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received:", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received:", msg2)
        }
    }
}
```

### 非阻塞select
```go
func main() {
    ch := make(chan int)
    
    select {
    case value := <-ch:
        fmt.Println("Received:", value)
    default:
        fmt.Println("No data available")  // 立即执行
    }
}
```

### 超时处理
```go
func main() {
    ch := make(chan string)
    
    go func() {
        time.Sleep(time.Second * 2)
        ch <- "data"
    }()
    
    select {
    case data := <-ch:
        fmt.Println("Received:", data)
    case <-time.After(time.Second):
        fmt.Println("Timeout!")
    }
}
```

## 5. 同步原语

虽然Go推荐使用channel，但也提供了传统的同步原语。

### Mutex互斥锁
```go
import "sync"

type Counter struct {
    mu    sync.Mutex
    value int
}

func (c *Counter) Increment() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *Counter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

func main() {
    counter := &Counter{}
    
    // 启动多个goroutine并发修改
    for i := 0; i < 1000; i++ {
        go counter.Increment()
    }
    
    time.Sleep(time.Second)
    fmt.Printf("Final value: %d\n", counter.Value())
}
```

### RWMutex读写锁
```go
type SafeMap struct {
    mu   sync.RWMutex
    data map[string]int
}

func NewSafeMap() *SafeMap {
    return &SafeMap{
        data: make(map[string]int),
    }
}

func (sm *SafeMap) Get(key string) (int, bool) {
    sm.mu.RLock()
    defer sm.mu.RUnlock()
    value, exists := sm.data[key]
    return value, exists
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()
    defer sm.mu.Unlock()
    sm.data[key] = value
}
```

### WaitGroup等待组
```go
func main() {
    var wg sync.WaitGroup
    
    // 启动5个goroutine
    for i := 1; i <= 5; i++ {
        wg.Add(1)  // 增加等待计数
        
        go func(id int) {
            defer wg.Done()  // 完成时减少计数
            
            fmt.Printf("Worker %d starting\n", id)
            time.Sleep(time.Second)
            fmt.Printf("Worker %d done\n", id)
        }(i)
    }
    
    wg.Wait()  // 等待所有goroutine完成
    fmt.Println("All workers completed")
}
```

### Once单次执行
```go
var once sync.Once
var instance *Database

type Database struct {
    connection string
}

func GetDatabase() *Database {
    once.Do(func() {
        fmt.Println("Creating database instance")
        instance = &Database{connection: "localhost:3306"}
    })
    return instance
}
```

## 6. 并发模式

### Worker Pool模式
```go
// Java需要使用ExecutorService
// Go使用goroutine和channel实现

type Job struct {
    ID   int
    Data string
}

type Result struct {
    JobID  int
    Output string
}

func worker(id int, jobs <-chan Job, results chan<- Result) {
    for job := range jobs {
        fmt.Printf("Worker %d processing job %d\n", id, job.ID)
        
        // 模拟工作
        time.Sleep(time.Millisecond * 100)
        
        results <- Result{
            JobID:  job.ID,
            Output: fmt.Sprintf("Processed by worker %d", id),
        }
    }
}

func main() {
    jobs := make(chan Job, 100)
    results := make(chan Result, 100)
    
    // 启动3个worker
    for w := 1; w <= 3; w++ {
        go worker(w, jobs, results)
    }
    
    // 发送任务
    for j := 1; j <= 9; j++ {
        jobs <- Job{ID: j, Data: fmt.Sprintf("data-%d", j)}
    }
    close(jobs)
    
    // 收集结果
    for r := 1; r <= 9; r++ {
        result := <-results
        fmt.Printf("Result: Job %d -> %s\n", result.JobID, result.Output)
    }
}
```

### Pipeline模式
```go
// 阶段1：生成数字
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

// 阶段2：计算平方
func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

// 阶段3：过滤偶数
func filter(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            if n%2 == 0 {
                out <- n
            }
        }
    }()
    return out
}

func main() {
    // 构建pipeline
    numbers := generate(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
    squares := square(numbers)
    evens := filter(squares)
    
    // 消费结果
    for result := range evens {
        fmt.Println(result)  // 输出：4, 16, 36, 64, 100
    }
}
```

## 7. 错误处理和取消

### Context上下文
```go
import "context"

func doWork(ctx context.Context, id int) {
    for {
        select {
        case <-ctx.Done():
            fmt.Printf("Worker %d cancelled: %v\n", id, ctx.Err())
            return
        default:
            fmt.Printf("Worker %d working...\n", id)
            time.Sleep(time.Millisecond * 500)
        }
    }
}

func main() {
    // 带超时的context
    ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
    defer cancel()
    
    // 启动多个worker
    for i := 1; i <= 3; i++ {
        go doWork(ctx, i)
    }
    
    time.Sleep(time.Second * 3)
}
```

## 8. 并发安全

### 数据竞争检测
```bash
# 运行时检测数据竞争
go run -race main.go

# 构建时包含竞争检测
go build -race
```

### 原子操作
```go
import "sync/atomic"

func main() {
    var counter int64
    
    // 并发增加计数器
    for i := 0; i < 1000; i++ {
        go func() {
            atomic.AddInt64(&counter, 1)
        }()
    }
    
    time.Sleep(time.Second)
    fmt.Printf("Counter: %d\n", atomic.LoadInt64(&counter))
}
```

## 性能对比

| 特性 | Java | Go | 优势 |
|------|------|----|------|
| 线程创建 | ~2MB内存 | ~2KB内存 | Go轻量1000倍 |
| 上下文切换 | 较重 | 极轻 | Go性能更好 |
| 并发数量 | 数千 | 数百万 | Go扩展性更好 |
| 编程复杂度 | 较高 | 较低 | Go更简洁 |

## 最佳实践

1. **优先使用channel而不是共享内存**
2. **使用select处理多个channel**
3. **及时关闭不再使用的channel**
4. **使用context处理取消和超时**
5. **避免在goroutine中使用panic**

## 常见陷阱

1. **Goroutine泄漏**：未正确关闭channel或context
2. **数据竞争**：多个goroutine同时访问共享数据
3. **死锁**：循环等待或错误的加锁顺序
4. **Channel阻塞**：向已满的无缓冲channel发送数据

## 实践任务

1. **并发下载器**：使用worker pool下载多个文件
2. **聊天服务器**：使用channel广播消息
3. **限流器**：实现令牌桶算法
4. **缓存系统**：实现并发安全的LRU缓存

## 下一步

掌握了并发编程后，建议学习：
- 网络编程
- 微服务开发
- 性能优化
- 分布式系统

## 参考资源

- [Go并发模式](https://golang.org/doc/codewalk/sharemem/)
- [Effective Go - 并发](https://golang.org/doc/effective_go#concurrency)
- [Go内存模型](https://golang.org/ref/mem) 