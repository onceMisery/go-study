package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 基础goroutine示例
func basicGoroutine() {
	fmt.Println("=== 基础 Goroutine 示例 ===")

	// 普通函数调用
	fmt.Println("开始执行...")

	// 启动goroutine - 异步执行
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Goroutine: %d\n", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// 主goroutine继续执行
	for i := 0; i < 5; i++ {
		fmt.Printf("Main: %d\n", i)
		time.Sleep(150 * time.Millisecond)
	}

	// 等待一下，让goroutine完成
	time.Sleep(1 * time.Second)
	fmt.Println("程序结束")
}

// 使用WaitGroup等待goroutines完成
func waitGroupExample() {
	fmt.Println("\n=== WaitGroup 示例 ===")

	var wg sync.WaitGroup

	// 启动多个goroutines
	for i := 1; i <= 3; i++ {
		wg.Add(1) // 增加等待计数

		go func(id int) {
			defer wg.Done() // 完成时减少计数

			fmt.Printf("Worker %d 开始工作\n", id)
			time.Sleep(time.Duration(id) * time.Second)
			fmt.Printf("Worker %d 完成工作\n", id)
		}(i)
	}

	fmt.Println("等待所有worker完成...")
	wg.Wait() // 等待所有goroutines完成
	fmt.Println("所有工作完成！")
}

// 工作任务结构体
type Task struct {
	ID   int
	Name string
	Data string
}

// 工作者函数
func worker(id int, tasks <-chan Task, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range tasks {
		fmt.Printf("Worker %d 开始处理任务 %d: %s\n", id, task.ID, task.Name)

		// 模拟工作时间
		time.Sleep(time.Duration(task.ID*100) * time.Millisecond)

		// 处理结果
		result := fmt.Sprintf("Worker %d 完成任务 %d: %s -> 处理了 %s",
			id, task.ID, task.Name, task.Data)

		results <- result
	}

	fmt.Printf("Worker %d 退出\n", id)
}

// 工作者池模式
func workerPoolExample() {
	fmt.Println("\n=== Worker Pool 模式 ===")

	const numWorkers = 3
	const numTasks = 8

	// 创建通道
	tasks := make(chan Task, numTasks)
	results := make(chan string, numTasks)

	var wg sync.WaitGroup

	// 启动workers
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, results, &wg)
	}

	// 发送任务
	go func() {
		for i := 1; i <= numTasks; i++ {
			task := Task{
				ID:   i,
				Name: fmt.Sprintf("任务_%d", i),
				Data: fmt.Sprintf("数据_%d", i*10),
			}
			tasks <- task
		}
		close(tasks) // 关闭任务通道
	}()

	// 启动结果收集器
	go func() {
		wg.Wait()      // 等待所有worker完成
		close(results) // 关闭结果通道
	}()

	// 收集结果
	fmt.Println("收集结果:")
	for result := range results {
		fmt.Println(result)
	}
}

// 并发安全的计数器
type SafeCounter struct {
	mu    sync.Mutex
	value int
}

// Increment 增加计数器
func (c *SafeCounter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value 获取当前值
func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.value
}

// 竞态条件示例
func raceConditionExample() {
	fmt.Println("\n=== 竞态条件和并发安全 ===")

	counter := &SafeCounter{}
	var wg sync.WaitGroup

	// 启动100个goroutines，每个增加100次
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				counter.Increment()
			}
		}()
	}

	wg.Wait()
	fmt.Printf("最终计数: %d (期望: 10000)\n", counter.Value())
}

// CPU密集型任务
func cpuIntensiveTask(id int, n int) int {
	result := 0
	for i := 0; i < n; i++ {
		result += i * i
	}
	fmt.Printf("任务 %d 完成，计算了 %d 次，结果: %d\n", id, n, result)
	return result
}

// 并行计算示例
func parallelComputingExample() {
	fmt.Println("\n=== 并行计算示例 ===")

	// 显示可用CPU核心数
	numCPU := runtime.NumCPU()
	fmt.Printf("可用CPU核心数: %d\n", numCPU)

	// 设置使用的CPU核心数
	runtime.GOMAXPROCS(numCPU)

	const numTasks = 8
	const workSize = 1000000

	// 串行计算
	start := time.Now()
	serialResult := 0
	for i := 0; i < numTasks; i++ {
		serialResult += cpuIntensiveTask(i, workSize)
	}
	serialTime := time.Since(start)

	fmt.Printf("串行计算用时: %v\n", serialTime)

	// 并行计算
	start = time.Now()
	results := make(chan int, numTasks)

	for i := 0; i < numTasks; i++ {
		go func(id int) {
			result := cpuIntensiveTask(id, workSize)
			results <- result
		}(i)
	}

	parallelResult := 0
	for i := 0; i < numTasks; i++ {
		parallelResult += <-results
	}

	parallelTime := time.Since(start)

	fmt.Printf("并行计算用时: %v\n", parallelTime)
	fmt.Printf("加速比: %.2fx\n", float64(serialTime)/float64(parallelTime))
	fmt.Printf("结果一致性: %t\n", serialResult == parallelResult)
}

// 超时控制示例
func timeoutExample() {
	fmt.Println("\n=== 超时控制示例 ===")

	timeout := 2 * time.Second

	// 创建一个可能超时的任务
	taskDone := make(chan bool)

	go func() {
		// 模拟长时间运行的任务
		time.Sleep(3 * time.Second)
		taskDone <- true
	}()

	// 使用select实现超时控制
	select {
	case <-taskDone:
		fmt.Println("任务在时间内完成")
	case <-time.After(timeout):
		fmt.Printf("任务超时 (%v)\n", timeout)
	}
}

func main() {
	fmt.Printf("Go版本: %s\n", runtime.Version())
	fmt.Printf("操作系统: %s\n", runtime.GOOS)
	fmt.Printf("CPU架构: %s\n", runtime.GOARCH)

	// 运行各种goroutine示例
	basicGoroutine()
	waitGroupExample()
	workerPoolExample()
	raceConditionExample()
	parallelComputingExample()
	timeoutExample()

	fmt.Println("\n=== Goroutine 示例完成 ===")
}
