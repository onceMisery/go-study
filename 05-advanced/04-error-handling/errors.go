package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// ========== 基础错误处理 ==========

// 基础错误创建和处理
func basicErrorHandling() {
	fmt.Println("=== 基础错误处理 ===")

	// 使用errors.New创建错误
	err1 := errors.New("这是一个简单错误")
	fmt.Printf("错误1: %v\n", err1)

	// 使用fmt.Errorf创建格式化错误
	username := "admin"
	err2 := fmt.Errorf("用户 '%s' 不存在", username)
	fmt.Printf("错误2: %v\n", err2)

	// 错误处理模式
	result, err := divide(10, 0)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("除法结果: %.2f\n", result)
	}

	result, err = divide(10, 2)
	if err != nil {
		fmt.Printf("除法错误: %v\n", err)
	} else {
		fmt.Printf("除法结果: %.2f\n", result)
	}
}

// divide 除法函数，演示错误返回
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("除数不能为零: %.2f / %.2f", a, b)
	}
	return a / b, nil
}

// ========== 自定义错误类型 ==========

// 网络错误类型
type NetworkError struct {
	URL     string
	Code    int
	Message string
	Time    time.Time
}

// Error 实现error接口
func (e *NetworkError) Error() string {
	return fmt.Sprintf("网络错误 [%s] %d: %s (时间: %v)",
		e.URL, e.Code, e.Message, e.Time.Format("15:04:05"))
}

// IsTimeout 检查是否为超时错误
func (e *NetworkError) IsTimeout() bool {
	return e.Code == 408
}

// IsServerError 检查是否为服务器错误
func (e *NetworkError) IsServerError() bool {
	return e.Code >= 500
}

// 业务逻辑错误
type BusinessError struct {
	Operation string
	Reason    string
	Code      string
}

func (e *BusinessError) Error() string {
	return fmt.Sprintf("业务错误 [%s]: %s (代码: %s)", e.Operation, e.Reason, e.Code)
}

// ========== 错误包装和解包 ==========

// 错误包装示例 (Go 1.13+)
func errorWrappingExample() {
	fmt.Println("\n=== 错误包装和解包 ===")

	// 模拟一个可能失败的操作
	err := performOperation("重要任务")
	if err != nil {
		fmt.Printf("操作失败: %v\n", err)

		// 检查是否包含特定错误
		var netErr *NetworkError
		if errors.As(err, &netErr) {
			fmt.Printf("  -> 网络错误详情: URL=%s, Code=%d\n", netErr.URL, netErr.Code)
			if netErr.IsTimeout() {
				fmt.Println("  -> 这是一个超时错误")
			}
		}

		var bizErr *BusinessError
		if errors.As(err, &bizErr) {
			fmt.Printf("  -> 业务错误详情: 操作=%s, 代码=%s\n", bizErr.Operation, bizErr.Code)
		}

		// 检查是否是某种类型的错误
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("  -> 文件不存在错误")
		}
	}
}

func performOperation(taskName string) error {
	// 模拟网络调用失败
	if err := callAPI("https://api.example.com/data"); err != nil {
		return fmt.Errorf("执行任务 '%s' 时网络调用失败: %w", taskName, err)
	}

	// 模拟业务逻辑失败
	if err := validateData(); err != nil {
		return fmt.Errorf("执行任务 '%s' 时数据验证失败: %w", taskName, err)
	}

	return nil
}

func callAPI(url string) error {
	// 模拟网络错误
	return &NetworkError{
		URL:     url,
		Code:    408,
		Message: "请求超时",
		Time:    time.Now(),
	}
}

func validateData() error {
	// 模拟业务错误
	return &BusinessError{
		Operation: "数据验证",
		Reason:    "缺少必填字段",
		Code:      "MISSING_FIELD",
	}
}

// ========== 错误处理策略 ==========

// 文件操作错误处理
func fileOperationExample() {
	fmt.Println("\n=== 文件操作错误处理 ===")

	// 尝试读取文件
	content, err := readFileWithRetry("test.txt", 3)
	if err != nil {
		fmt.Printf("读取文件失败: %v\n", err)

		// 根据错误类型采取不同处理策略
		if os.IsNotExist(err) {
			fmt.Println("  -> 文件不存在，尝试创建默认文件")
			if createErr := createDefaultFile("test.txt"); createErr != nil {
				fmt.Printf("  -> 创建默认文件失败: %v\n", createErr)
			} else {
				fmt.Println("  -> 默认文件创建成功")
			}
		} else if os.IsPermission(err) {
			fmt.Println("  -> 权限不足，请检查文件权限")
		} else {
			fmt.Println("  -> 其他文件系统错误")
		}
	} else {
		fmt.Printf("文件内容: %s\n", content)
	}
}

// 带重试的文件读取
func readFileWithRetry(filename string, maxRetries int) (string, error) {
	var lastErr error

	for i := 0; i < maxRetries; i++ {
		content, err := readFile(filename)
		if err == nil {
			return content, nil
		}

		lastErr = err
		fmt.Printf("第 %d 次尝试失败: %v\n", i+1, err)

		// 如果是文件不存在，不需要重试
		if os.IsNotExist(err) {
			break
		}

		// 等待一段时间再重试
		time.Sleep(time.Duration(i+1) * 100 * time.Millisecond)
	}

	return "", fmt.Errorf("重试 %d 次后仍然失败: %w", maxRetries, lastErr)
}

func readFile(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	content := make([]byte, 1024)
	n, err := file.Read(content)
	if err != nil && err != io.EOF {
		return "", err
	}

	return string(content[:n]), nil
}

func createDefaultFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	defaultContent := "这是一个默认文件内容\n创建时间: " + time.Now().Format("2006-01-02 15:04:05")
	_, err = file.WriteString(defaultContent)
	return err
}

// ========== 错误聚合 ==========

// MultiError 多错误聚合
type MultiError struct {
	errors []error
}

func (m *MultiError) Error() string {
	if len(m.errors) == 0 {
		return "无错误"
	}

	if len(m.errors) == 1 {
		return m.errors[0].Error()
	}

	result := fmt.Sprintf("发生 %d 个错误:\n", len(m.errors))
	for i, err := range m.errors {
		result += fmt.Sprintf("  %d. %v\n", i+1, err)
	}
	return result
}

func (m *MultiError) Add(err error) {
	if err != nil {
		m.errors = append(m.errors, err)
	}
}

func (m *MultiError) HasErrors() bool {
	return len(m.errors) > 0
}

// 批量处理示例
func batchProcessExample() {
	fmt.Println("\n=== 批量处理错误聚合 ===")

	tasks := []string{"任务1", "任务2", "任务3", "任务4", "任务5"}
	var multiErr MultiError

	for _, task := range tasks {
		if err := processTask(task); err != nil {
			multiErr.Add(fmt.Errorf("处理 %s 失败: %w", task, err))
		}
	}

	if multiErr.HasErrors() {
		fmt.Printf("批量处理完成，但有错误:\n%v", multiErr.Error())
	} else {
		fmt.Println("批量处理全部成功")
	}
}

func processTask(taskName string) error {
	// 模拟一些任务会失败
	switch taskName {
	case "任务2":
		return errors.New("网络连接失败")
	case "任务4":
		return errors.New("数据格式错误")
	default:
		return nil
	}
}

// ========== panic 和 recover ==========

func panicRecoverExample() {
	fmt.Println("\n=== Panic 和 Recover 示例 ===")

	// 安全执行可能panic的函数
	result := safeExecute(func() interface{} {
		return riskyOperation(10, 0)
	})

	if result != nil {
		fmt.Printf("安全执行结果: %v\n", result)
	}

	result = safeExecute(func() interface{} {
		return riskyOperation(10, 2)
	})

	if result != nil {
		fmt.Printf("安全执行结果: %v\n", result)
	}
}

// safeExecute 安全执行函数，捕获panic
func safeExecute(fn func() interface{}) (result interface{}) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("捕获到panic: %v\n", r)
			result = nil
		}
	}()

	return fn()
}

func riskyOperation(a, b int) int {
	if b == 0 {
		panic("除数为零会导致panic")
	}
	return a / b
}

// ========== 实用工具函数 ==========

// Must 包装器 - 如果有错误就panic（用于初始化）
func Must[T any](value T, err error) T {
	if err != nil {
		panic(err)
	}
	return value
}

// 使用Must的示例
func mustExample() {
	fmt.Println("\n=== Must 包装器示例 ===")

	// 在确定不会出错的地方使用Must
	port := Must(strconv.Atoi("8080"))
	fmt.Printf("端口号: %d\n", port)

	// 这会导致panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Must捕获到panic: %v\n", r)
		}
	}()

	invalidPort := Must(strconv.Atoi("invalid"))
	fmt.Printf("无效端口号: %d\n", invalidPort) // 这行不会执行
}

// ========== 主函数 ==========

func main() {
	fmt.Println("=== Go 错误处理综合示例 ===")

	basicErrorHandling()
	errorWrappingExample()
	fileOperationExample()
	batchProcessExample()
	panicRecoverExample()
	mustExample()

	fmt.Println("\n=== 错误处理示例完成 ===")
}
