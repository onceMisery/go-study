package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// 基础channel示例
func basicChannelExample() {
	fmt.Println("=== 基础 Channel 示例 ===")

	// 创建一个字符串类型的channel
	messages := make(chan string)

	// 启动goroutine发送消息
	go func() {
		messages <- "Hello"
		messages <- "World"
		messages <- "from Go"
		close(messages) // 关闭channel
	}()

	// 接收消息
	for msg := range messages {
		fmt.Println("收到消息:", msg)
	}
}

// 缓冲channel示例
func bufferedChannelExample() {
	fmt.Println("\n=== 缓冲 Channel 示例 ===")

	// 创建容量为3的缓冲channel
	ch := make(chan int, 3)

	// 发送数据（不会阻塞，因为有缓冲区）
	ch <- 1
	ch <- 2
	ch <- 3

	fmt.Printf("Channel长度: %d, 容量: %d\n", len(ch), cap(ch))

	// 接收数据
	fmt.Println("接收:", <-ch)
	fmt.Println("接收:", <-ch)
	fmt.Println("接收:", <-ch)
}

// select语句示例
func selectExample() {
	fmt.Println("\n=== Select 语句示例 ===")

	c1 := make(chan string)
	c2 := make(chan string)

	// 启动两个goroutines
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "来自 c1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "来自 c2"
	}()

	// 使用select等待多个channel
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("收到:", msg1)
		case msg2 := <-c2:
			fmt.Println("收到:", msg2)
		case <-time.After(3 * time.Second):
			fmt.Println("超时")
			return
		}
	}
}

// 单向channel示例
func unidirectionalChannelExample() {
	fmt.Println("\n=== 单向 Channel 示例 ===")

	// 双向channel
	ch := make(chan int)

	// 启动发送者
	go sender(ch)
	// 启动接收者
	receiver(ch)
}

// 发送者函数（只能发送）
func sender(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("发送: %d\n", i)
		ch <- i
		time.Sleep(500 * time.Millisecond)
	}
	close(ch)
}

// 接收者函数（只能接收）
func receiver(ch <-chan int) {
	for num := range ch {
		fmt.Printf("接收: %d\n", num)
	}
}

// 生产者-消费者模式
func producerConsumerExample() {
	fmt.Println("\n=== 生产者-消费者模式 ===")

	const bufferSize = 5
	const numProducers = 2
	const numConsumers = 3
	const numItems = 20

	// 创建缓冲channel
	jobs := make(chan int, bufferSize)
	results := make(chan string, bufferSize)

	var wg sync.WaitGroup

	// 启动生产者
	for p := 1; p <= numProducers; p++ {
		wg.Add(1)
		go producer(p, jobs, &wg, numItems/numProducers)
	}

	// 启动消费者
	for c := 1; c <= numConsumers; c++ {
		wg.Add(1)
		go consumer(c, jobs, results, &wg)
	}

	// 启动结果收集器
	go func() {
		wg.Wait()
		close(results)
	}()

	// 收集结果
	fmt.Println("处理结果:")
	for result := range results {
		fmt.Println(result)
	}
}

// 生产者函数
func producer(id int, jobs chan<- int, wg *sync.WaitGroup, count int) {
	defer wg.Done()
	defer func() {
		fmt.Printf("Producer %d 完成\n", id)
	}()

	for i := 0; i < count; i++ {
		job := id*1000 + i // 生成唯一的job ID
		fmt.Printf("Producer %d 生产 job %d\n", id, job)
		jobs <- job
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	}
}

// 消费者函数
func consumer(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		fmt.Printf("Consumer %d 完成\n", id)
	}()

	for job := range jobs {
		fmt.Printf("Consumer %d 开始处理 job %d\n", id, job)

		// 模拟处理时间
		processingTime := time.Duration(rand.Intn(1000)) * time.Millisecond
		time.Sleep(processingTime)

		result := fmt.Sprintf("Consumer %d 完成 job %d (用时 %v)",
			id, job, processingTime)
		results <- result
	}
}

// 扇入模式（Fan-in）
func fanInExample() {
	fmt.Println("\n=== 扇入模式 (Fan-in) ===")

	// 创建两个输入channel
	input1 := make(chan string)
	input2 := make(chan string)

	// 启动数据源
	go func() {
		for i := 1; i <= 5; i++ {
			input1 <- fmt.Sprintf("输入1-%d", i)
			time.Sleep(200 * time.Millisecond)
		}
		close(input1)
	}()

	go func() {
		for i := 1; i <= 5; i++ {
			input2 <- fmt.Sprintf("输入2-%d", i)
			time.Sleep(300 * time.Millisecond)
		}
		close(input2)
	}()

	// 扇入：合并多个channel
	output := fanIn(input1, input2)

	// 接收合并后的数据
	for data := range output {
		fmt.Println("扇入接收:", data)
	}
}

// 扇入函数：合并多个channel为一个
func fanIn(inputs ...<-chan string) <-chan string {
	output := make(chan string)
	var wg sync.WaitGroup

	// 为每个输入channel启动一个goroutine
	for _, input := range inputs {
		wg.Add(1)
		go func(ch <-chan string) {
			defer wg.Done()
			for data := range ch {
				output <- data
			}
		}(input)
	}

	// 启动goroutine等待所有输入完成后关闭输出
	go func() {
		wg.Wait()
		close(output)
	}()

	return output
}

// 扇出模式（Fan-out）
func fanOutExample() {
	fmt.Println("\n=== 扇出模式 (Fan-out) ===")

	// 创建输入数据
	input := make(chan int)

	// 启动数据源
	go func() {
		for i := 1; i <= 10; i++ {
			input <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(input)
	}()

	// 扇出：分发到多个worker
	const numWorkers = 3
	outputs := fanOut(input, numWorkers)

	// 收集所有worker的结果
	var wg sync.WaitGroup
	for i, output := range outputs {
		wg.Add(1)
		go func(workerID int, ch <-chan int) {
			defer wg.Done()
			for data := range ch {
				result := data * data // 计算平方
				fmt.Printf("Worker %d 处理 %d -> %d\n", workerID, data, result)
				time.Sleep(200 * time.Millisecond)
			}
		}(i+1, output)
	}

	wg.Wait()
}

// 扇出函数：将一个channel分发到多个channel
func fanOut(input <-chan int, numWorkers int) []<-chan int {
	outputs := make([]chan int, numWorkers)
	outputsRead := make([]<-chan int, numWorkers)

	// 创建输出channels
	for i := range outputs {
		outputs[i] = make(chan int)
		outputsRead[i] = outputs[i]
	}

	// 启动分发器
	go func() {
		defer func() {
			for _, output := range outputs {
				close(output)
			}
		}()

		i := 0
		for data := range input {
			// 轮询分发
			outputs[i%numWorkers] <- data
			i++
		}
	}()

	return outputsRead
}

// 管道模式 (Pipeline)
func pipelineExample() {
	fmt.Println("\n=== 管道模式 (Pipeline) ===")

	// 创建数据源
	numbers := generateNumbers(1, 10)

	// 构建处理管道
	squared := square(numbers)
	filtered := filter(squared, func(n int) bool { return n > 20 })

	// 消费最终结果
	fmt.Println("管道处理结果:")
	for result := range filtered {
		fmt.Printf("最终结果: %d\n", result)
	}
}

// 生成数字
func generateNumbers(start, end int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= end; i++ {
			ch <- i
			time.Sleep(100 * time.Millisecond)
		}
	}()
	return ch
}

// 计算平方
func square(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			result := num * num
			fmt.Printf("平方: %d -> %d\n", num, result)
			output <- result
		}
	}()
	return output
}

// 过滤器
func filter(input <-chan int, predicate func(int) bool) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			if predicate(num) {
				fmt.Printf("过滤通过: %d\n", num)
				output <- num
			} else {
				fmt.Printf("过滤拒绝: %d\n", num)
			}
		}
	}()
	return output
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Go Channel 综合示例 ===")

	basicChannelExample()
	bufferedChannelExample()
	selectExample()
	unidirectionalChannelExample()
	producerConsumerExample()
	fanInExample()
	fanOutExample()
	pipelineExample()

	fmt.Println("\n=== Channel 示例完成 ===")
}
