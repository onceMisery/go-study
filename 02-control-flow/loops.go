package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Product 商品结构体 - 用于演示循环操作
type Product struct {
	ID    int
	Name  string
	Price float64
	Stock int
}

// 1. 基本for循环
// Go只有for循环，没有while和do-while（与Java不同）
func basicLoops() {
	fmt.Println("=== 基本for循环 ===")

	// 传统三段式for循环（类似Java）
	fmt.Println("1. 传统for循环:")
	for i := 0; i < 5; i++ {
		fmt.Printf("  第%d次循环\n", i+1)
	}

	// 模拟while循环（Java: while(condition)）
	fmt.Println("2. 模拟while循环:")
	count := 0
	for count < 3 {
		fmt.Printf("  count = %d\n", count)
		count++
	}

	// 无限循环（相当于Java的while(true)）
	fmt.Println("3. 无限循环示例（会自动退出）:")
	counter := 0
	for {
		counter++
		fmt.Printf("  无限循环第%d次\n", counter)
		if counter >= 3 {
			break // 使用break退出循环
		}
	}
}

// 2. range循环 - Go特有的迭代方式
// 相当于Java的foreach循环
func rangeLoops() {
	fmt.Println("\n=== range循环 ===")

	// 数组/切片的range循环
	numbers := []int{10, 20, 30, 40, 50}

	fmt.Println("1. 遍历切片（获取索引和值）:")
	// 类似 for (int i = 0; i < array.length; i++)
	for index, value := range numbers {
		fmt.Printf("  索引%d: 值%d\n", index, value)
	}

	fmt.Println("2. 只遍历值（忽略索引）:")
	// for (Type value : collection)
	for _, value := range numbers {
		fmt.Printf("  值: %d\n", value)
	}

	fmt.Println("3. 只遍历索引（忽略值）:")
	for index := range numbers {
		fmt.Printf("  索引: %d\n", index)
	}

	// 字符串的range循环
	fmt.Println("4. 遍历字符串:")
	text := "Hello世界"
	for i, char := range text {
		fmt.Printf("  位置%d: 字符%c (Unicode: %d)\n", i, char, char)
	}

	// Map的range循环
	fmt.Println("5. 遍历Map:")
	scores := map[string]int{
		"张三": 85,
		"李四": 92,
		"王五": 78,
	}
	for name, score := range scores {
		fmt.Printf("  %s: %d分\n", name, score)
	}
}

// 3. 循环控制语句 - break, continue
// 与Java的用法相同
func loopControl() {
	fmt.Println("\n=== 循环控制语句 ===")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("1. continue示例（跳过偶数）:")
	for _, num := range numbers {
		if num%2 == 0 {
			continue // 跳过本次循环
		}
		fmt.Printf("  奇数: %d\n", num)
	}

	fmt.Println("2. break示例（找到第一个大于5的数就退出）:")
	for _, num := range numbers {
		if num > 5 {
			fmt.Printf("  找到第一个大于5的数: %d\n", num)
			break // 退出循环
		}
		fmt.Printf("  检查数字: %d\n", num)
	}
}

// 4. 嵌套循环
// 多层循环的使用，包括标签的使用
func nestedLoops() {
	fmt.Println("\n=== 嵌套循环 ===")

	fmt.Println("1. 九九乘法表:")
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d×%d=%d  ", j, i, i*j)
		}
		fmt.Println()
	}

	fmt.Println("2. 使用标签控制嵌套循环:")
	// Go支持标签，可以指定break或continue的目标循环
OuterLoop:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i == 2 && j == 2 {
				fmt.Printf("  在(%d,%d)处跳出外层循环\n", i, j)
				break OuterLoop // 跳出外层循环
			}
			fmt.Printf("  (%d,%d) ", i, j)
		}
		fmt.Println()
	}
}

// 5. 实际业务场景中的循环应用
func businessLoops() {
	fmt.Println("\n=== 业务场景循环应用 ===")

	// 商品库存管理示例
	products := []Product{
		{1, "笔记本电脑", 5999.99, 10},
		{2, "手机", 2999.99, 0},
		{3, "平板电脑", 1999.99, 5},
		{4, "智能手表", 999.99, 15},
		{5, "耳机", 299.99, 0},
	}

	fmt.Println("1. 库存检查:")
	var outOfStock []Product
	var lowStock []Product
	totalValue := 0.0

	for _, product := range products {
		totalValue += product.Price * float64(product.Stock)

		if product.Stock == 0 {
			outOfStock = append(outOfStock, product)
			fmt.Printf("  ❌ %s 缺货\n", product.Name)
		} else if product.Stock < 5 {
			lowStock = append(lowStock, product)
			fmt.Printf("  ⚠️  %s 库存不足(%d件)\n", product.Name, product.Stock)
		} else {
			fmt.Printf("  ✅ %s 库存充足(%d件)\n", product.Name, product.Stock)
		}
	}

	fmt.Printf("\n库存总价值: ¥%.2f\n", totalValue)
	fmt.Printf("缺货商品数量: %d\n", len(outOfStock))
	fmt.Printf("库存不足商品数量: %d\n", len(lowStock))

	// 批量处理示例
	fmt.Println("\n2. 批量价格调整（打折优惠）:")
	discountRate := 0.1 // 10%折扣
	for i := range products {
		originalPrice := products[i].Price
		products[i].Price = originalPrice * (1 - discountRate)
		fmt.Printf("  %s: ¥%.2f -> ¥%.2f\n",
			products[i].Name, originalPrice, products[i].Price)
	}
}

// 6. 性能相关的循环优化
// 演示循环中的性能注意事项
func performanceLoops() {
	fmt.Println("\n=== 循环性能优化 ===")

	// 大数据量处理示例
	size := 1000000
	data := make([]int, size)

	// 初始化数据
	fmt.Println("1. 高效的数据初始化:")
	start := time.Now()
	for i := range data {
		data[i] = rand.Intn(1000)
	}
	duration := time.Since(start)
	fmt.Printf("  初始化%d个元素耗时: %v\n", size, duration)

	// 避免在循环中重复计算
	fmt.Println("2. 避免重复计算（优化前后对比）:")

	// 优化前：每次循环都计算len(data)
	start = time.Now()
	sum1 := 0
	for i := 0; i < len(data); i++ { // 不推荐：每次都调用len()
		sum1 += data[i]
	}
	duration1 := time.Since(start)

	// 优化后：提前计算长度
	start = time.Now()
	sum2 := 0
	length := len(data)
	for i := 0; i < length; i++ { // 推荐：提前计算长度
		sum2 += data[i]
	}
	duration2 := time.Since(start)

	// 最优：使用range
	start = time.Now()
	sum3 := 0
	for _, value := range data { // 最推荐：使用range
		sum3 += value
	}
	duration3 := time.Since(start)

	fmt.Printf("  方法1(重复计算len): %v, 结果: %d\n", duration1, sum1)
	fmt.Printf("  方法2(提前计算len): %v, 结果: %d\n", duration2, sum2)
	fmt.Printf("  方法3(使用range): %v, 结果: %d\n", duration3, sum3)
}

// 7. 错误处理中的循环
// 在循环中处理可能的错误
func errorHandlingLoops() {
	fmt.Println("\n=== 循环中的错误处理 ===")

	// 模拟可能失败的操作
	operations := []string{"op1", "op2", "fail", "op4", "op5"}

	processOperation := func(op string) error {
		if op == "fail" {
			return fmt.Errorf("操作%s失败", op)
		}
		return nil
	}

	fmt.Println("1. 遇到错误继续处理:")
	successCount := 0
	failCount := 0

	for _, op := range operations {
		if err := processOperation(op); err != nil {
			fmt.Printf("  ❌ %v\n", err)
			failCount++
			continue // 继续处理下一个
		}
		fmt.Printf("  ✅ 操作%s成功\n", op)
		successCount++
	}

	fmt.Printf("  处理结果: 成功%d个, 失败%d个\n", successCount, failCount)

	fmt.Println("2. 遇到错误立即停止:")
	for i, op := range operations {
		if err := processOperation(op); err != nil {
			fmt.Printf("  在第%d个操作时发生错误: %v\n", i+1, err)
			break // 立即停止
		}
		fmt.Printf("  ✅ 操作%s成功\n", op)
	}
}

func main() {
	fmt.Println("Go语言控制流 - 循环语句实践")
	fmt.Println("===============================")

	// 执行各种循环示例
	basicLoops()
	rangeLoops()
	loopControl()
	nestedLoops()
	businessLoops()
	performanceLoops()
	errorHandlingLoops()

	fmt.Println("\n学习要点:")
	fmt.Println("1. Go只有for循环，可以模拟while和do-while")
	fmt.Println("2. range是Go特有的迭代方式，非常强大")
	fmt.Println("3. 支持break和continue，以及标签控制")
	fmt.Println("4. 注意循环中的性能优化（避免重复计算）")
	fmt.Println("5. 在循环中正确处理错误和边界条件")
}
