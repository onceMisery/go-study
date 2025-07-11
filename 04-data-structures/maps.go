package main

import (
	"fmt"
	"sort"
	"strings"
)

// 1. Map基础 - Go的哈希表实现
// 类似Java的HashMap，但语法更简洁

func mapBasics() {
	fmt.Println("=== Map基础 ===")

	// Map声明和初始化
	var map1 map[string]int      // nil map，不能写入
	map2 := make(map[string]int) // 空map，可以写入
	map3 := map[string]int{      // 字面量初始化
		"apple":  100,
		"banana": 50,
		"orange": 75,
	}
	map4 := make(map[string]int, 10) // 预分配容量（提示性）

	fmt.Printf("map1 (nil): %v, 长度: %d\n", map1, len(map1))
	fmt.Printf("map2 (空): %v, 长度: %d\n", map2, len(map2))
	fmt.Printf("map3 (字面量): %v, 长度: %d\n", map3, len(map3))
	fmt.Printf("map4 (预分配): %v, 长度: %d\n", map4, len(map4))

	// nil map检查
	if map1 == nil {
		fmt.Println("map1 是 nil，不能写入")
	}

	// Map操作
	map2["hello"] = 100
	map2["world"] = 200
	fmt.Printf("添加元素后的map2: %v\n", map2)
}

func mapOperations() {
	fmt.Println("\n=== Map操作 ===")

	// 创建学生成绩Map
	scores := make(map[string]int)

	// 添加元素
	scores["张三"] = 85
	scores["李四"] = 92
	scores["王五"] = 78
	scores["赵六"] = 88

	fmt.Printf("学生成绩: %v\n", scores)

	// 访问元素
	score := scores["张三"]
	fmt.Printf("张三的成绩: %d\n", score)

	// 安全访问（检查键是否存在）
	if score, exists := scores["张三"]; exists {
		fmt.Printf("张三的成绩: %d\n", score)
	} else {
		fmt.Println("张三不存在")
	}

	// 访问不存在的键
	if score, exists := scores["陈七"]; exists {
		fmt.Printf("陈七的成绩: %d\n", score)
	} else {
		fmt.Printf("陈七不存在，默认值: %d\n", score) // 返回零值
	}

	// 修改元素
	scores["张三"] = 90
	fmt.Printf("修改后张三的成绩: %d\n", scores["张三"])

	// 删除元素
	delete(scores, "王五")
	fmt.Printf("删除王五后: %v\n", scores)

	// 遍历Map
	fmt.Println("遍历所有学生成绩:")
	for name, score := range scores {
		fmt.Printf("  %s: %d分\n", name, score)
	}

	// 只遍历键
	fmt.Println("所有学生姓名:")
	for name := range scores {
		fmt.Printf("  %s\n", name)
	}
}

// 2. 复杂数据类型作为值
type Student struct {
	Name    string
	Age     int
	Courses []string
	Grades  map[string]float64
}

func complexValueTypes() {
	fmt.Println("\n=== 复杂数据类型作为值 ===")

	// 学生信息Map
	students := make(map[string]Student)

	students["S001"] = Student{
		Name:    "张三",
		Age:     20,
		Courses: []string{"数学", "物理", "化学"},
		Grades: map[string]float64{
			"数学": 85.5,
			"物理": 92.0,
			"化学": 78.5,
		},
	}

	students["S002"] = Student{
		Name:    "李四",
		Age:     19,
		Courses: []string{"数学", "英语", "历史"},
		Grades: map[string]float64{
			"数学": 88.0,
			"英语": 95.5,
			"历史": 82.0,
		},
	}

	fmt.Println("学生信息:")
	for id, student := range students {
		fmt.Printf("  学号: %s\n", id)
		fmt.Printf("  姓名: %s, 年龄: %d\n", student.Name, student.Age)
		fmt.Printf("  课程: %v\n", student.Courses)
		fmt.Printf("  成绩: %v\n", student.Grades)

		// 计算平均分
		var total float64
		for _, grade := range student.Grades {
			total += grade
		}
		average := total / float64(len(student.Grades))
		fmt.Printf("  平均分: %.2f\n\n", average)
	}
}

// 3. Map作为缓存
func mapAsCache() {
	fmt.Println("\n=== Map作为缓存 ===")

	// 斐波那契数列缓存
	cache := make(map[int]int)

	var fibonacci func(n int) int
	fibonacci = func(n int) int {
		if n <= 1 {
			return n
		}

		// 检查缓存
		if result, exists := cache[n]; exists {
			fmt.Printf("缓存命中: fib(%d) = %d\n", n, result)
			return result
		}

		// 计算并缓存
		result := fibonacci(n-1) + fibonacci(n-2)
		cache[n] = result
		fmt.Printf("计算并缓存: fib(%d) = %d\n", n, result)
		return result
	}

	fmt.Println("计算斐波那契数列:")
	for i := 0; i <= 10; i++ {
		fmt.Printf("fib(%d) = %d\n", i, fibonacci(i))
	}

	fmt.Printf("\n缓存内容: %v\n", cache)
}

// 4. 字符统计和词频分析
func textAnalysis() {
	fmt.Println("\n=== 文本分析 ===")

	text := "Go语言是Google开发的编程语言。Go语言简洁、高效、并发性强。"

	// 字符频率统计
	charCount := make(map[rune]int)
	for _, char := range text {
		charCount[char]++
	}

	fmt.Println("字符频率统计:")
	for char, count := range charCount {
		if char != ' ' && char != '。' { // 跳过空格和句号
			fmt.Printf("  '%c': %d次\n", char, count)
		}
	}

	// 词频统计
	words := strings.Fields(strings.ReplaceAll(text, "。", ""))
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	fmt.Println("\n词频统计:")
	for word, count := range wordCount {
		fmt.Printf("  '%s': %d次\n", word, count)
	}
}

// 5. Map的排序
func mapSorting() {
	fmt.Println("\n=== Map排序 ===")

	// 创建测试数据
	population := map[string]int{
		"北京": 2154,
		"上海": 2424,
		"广州": 1491,
		"深圳": 1344,
		"杭州": 1036,
		"成都": 1658,
	}

	fmt.Printf("原始数据: %v\n", population)

	// 按键排序
	var cities []string
	for city := range population {
		cities = append(cities, city)
	}
	sort.Strings(cities)

	fmt.Println("\n按城市名排序:")
	for _, city := range cities {
		fmt.Printf("  %s: %d万人\n", city, population[city])
	}

	// 按值排序
	type CityPopulation struct {
		City string
		Pop  int
	}

	var cityPops []CityPopulation
	for city, pop := range population {
		cityPops = append(cityPops, CityPopulation{city, pop})
	}

	// 按人口倒序排序
	sort.Slice(cityPops, func(i, j int) bool {
		return cityPops[i].Pop > cityPops[j].Pop
	})

	fmt.Println("\n按人口倒序排序:")
	for _, cp := range cityPops {
		fmt.Printf("  %s: %d万人\n", cp.City, cp.Pop)
	}
}

// 6. 多级Map（嵌套Map）
func nestedMaps() {
	fmt.Println("\n=== 嵌套Map ===")

	// 公司部门员工信息
	company := make(map[string]map[string]string)

	// 初始化部门
	company["研发部"] = make(map[string]string)
	company["销售部"] = make(map[string]string)
	company["人事部"] = make(map[string]string)

	// 添加员工
	company["研发部"]["张三"] = "高级工程师"
	company["研发部"]["李四"] = "架构师"
	company["研发部"]["王五"] = "初级工程师"

	company["销售部"]["赵六"] = "销售经理"
	company["销售部"]["陈七"] = "销售代表"

	company["人事部"]["孙八"] = "HR经理"

	fmt.Println("公司组织架构:")
	for dept, employees := range company {
		fmt.Printf("  %s:\n", dept)
		for name, position := range employees {
			fmt.Printf("    %s - %s\n", name, position)
		}
	}

	// 查找特定员工
	targetName := "张三"
	found := false
	for dept, employees := range company {
		if position, exists := employees[targetName]; exists {
			fmt.Printf("\n找到员工 %s: %s部门, 职位: %s\n", targetName, dept, position)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("\n未找到员工: %s\n", targetName)
	}
}

// 7. Set模拟（使用Map）
func setSimulation() {
	fmt.Println("\n=== Set模拟 ===")

	// Go没有内置Set，使用map[type]bool模拟
	tags := make(map[string]bool)

	// 添加元素
	addToSet := func(set map[string]bool, item string) {
		set[item] = true
	}

	// 检查元素是否存在
	contains := func(set map[string]bool, item string) bool {
		return set[item]
	}

	// 删除元素
	removeFromSet := func(set map[string]bool, item string) {
		delete(set, item)
	}

	// 获取所有元素
	getItems := func(set map[string]bool) []string {
		var items []string
		for item := range set {
			items = append(items, item)
		}
		sort.Strings(items) // 排序以便输出一致
		return items
	}

	// 测试Set操作
	addToSet(tags, "Go")
	addToSet(tags, "Python")
	addToSet(tags, "Java")
	addToSet(tags, "Go") // 重复添加

	fmt.Printf("标签集合: %v\n", getItems(tags))
	fmt.Printf("包含'Go': %t\n", contains(tags, "Go"))
	fmt.Printf("包含'C++': %t\n", contains(tags, "C++"))

	removeFromSet(tags, "Python")
	fmt.Printf("删除Python后: %v\n", getItems(tags))

	// Set运算
	tags2 := map[string]bool{
		"Go":         true,
		"JavaScript": true,
		"C++":        true,
	}

	// 交集
	intersection := make(map[string]bool)
	for tag := range tags {
		if tags2[tag] {
			intersection[tag] = true
		}
	}
	fmt.Printf("交集: %v\n", getItems(intersection))

	// 并集
	union := make(map[string]bool)
	for tag := range tags {
		union[tag] = true
	}
	for tag := range tags2 {
		union[tag] = true
	}
	fmt.Printf("并集: %v\n", getItems(union))
}

// 8. Map的性能和注意事项
func mapPerformanceNotes() {
	fmt.Println("\n=== Map性能注意事项 ===")

	// 1. Map是引用类型
	original := map[string]int{"a": 1, "b": 2}
	copy := original
	copy["c"] = 3

	fmt.Printf("原Map: %v\n", original)
	fmt.Printf("拷贝Map: %v\n", copy)
	fmt.Println("Map是引用类型，修改拷贝会影响原Map")

	// 2. 零值处理
	var nilMap map[string]int
	safeMap := make(map[string]int)

	// nilMap["key"] = 1 // 运行时panic
	safeMap["key"] = 1

	fmt.Printf("安全Map: %v\n", safeMap)

	// 3. 并发安全性
	fmt.Println("\n注意: Map不是并发安全的")
	fmt.Println("多goroutine访问需要使用sync.RWMutex或sync.Map")

	// 4. 键类型限制
	fmt.Println("\n键类型必须是可比较的:")
	fmt.Println("- 可以: 基本类型、数组、结构体（字段都可比较）")
	fmt.Println("- 不可以: 切片、Map、函数")
}

func realWorldExamples() {
	fmt.Println("\n=== 实际应用场景 ===")

	// HTTP头部处理
	headers := map[string]string{
		"Content-Type":  "application/json",
		"Authorization": "Bearer token123",
		"User-Agent":    "Go-HTTP-Client",
		"Accept":        "application/json",
	}

	fmt.Println("HTTP请求头:")
	for key, value := range headers {
		fmt.Printf("  %s: %s\n", key, value)
	}

	// 配置管理
	config := map[string]interface{}{
		"database_host":   "localhost",
		"database_port":   5432,
		"debug_mode":      true,
		"max_connections": 100,
		"timeout":         30.5,
	}

	fmt.Println("\n应用配置:")
	for key, value := range config {
		fmt.Printf("  %s: %v (%T)\n", key, value, value)
	}

	// 用户会话管理
	sessions := make(map[string]map[string]interface{})

	sessionID := "sess_abc123"
	sessions[sessionID] = map[string]interface{}{
		"user_id":    12345,
		"username":   "张三",
		"login_time": "2024-01-15 10:30:00",
		"role":       "admin",
	}

	fmt.Println("\n用户会话:")
	if session, exists := sessions[sessionID]; exists {
		fmt.Printf("  会话ID: %s\n", sessionID)
		for key, value := range session {
			fmt.Printf("    %s: %v\n", key, value)
		}
	}
}

func main() {
	fmt.Println("Go语言数据结构 - Map实践")
	fmt.Println("========================")

	mapBasics()
	mapOperations()
	complexValueTypes()
	mapAsCache()
	textAnalysis()
	mapSorting()
	nestedMaps()
	setSimulation()
	mapPerformanceNotes()
	realWorldExamples()

	fmt.Println("\n学习要点:")
	fmt.Println("1. Map是引用类型，零值是nil")
	fmt.Println("2. 使用make()创建可写入的Map")
	fmt.Println("3. 访问不存在的键返回零值")
	fmt.Println("4. 使用value, ok := map[key]检查键是否存在")
	fmt.Println("5. 使用delete()函数删除键值对")
	fmt.Println("6. Map遍历顺序是随机的")
	fmt.Println("7. Map不是并发安全的")
	fmt.Println("8. 键类型必须是可比较的")
}
