package main

import (
	"fmt"
	"sort"
	"strings"
)

// 1. 数组 - 固定长度的数据结构
// Go数组：固定长度，值类型，长度是类型的一部分
// Java数组：固定长度，引用类型

func arrayBasics() {
	fmt.Println("=== 数组基础 ===")

	// 数组声明和初始化
	var arr1 [5]int               // 零值初始化：[0 0 0 0 0]
	arr2 := [5]int{1, 2, 3, 4, 5} // 字面量初始化
	arr3 := [...]int{10, 20, 30}  // 自动推断长度：[3]int
	arr4 := [5]int{1: 10, 3: 30}  // 指定索引初始化：[0 10 0 30 0]

	fmt.Printf("arr1 (零值): %v\n", arr1)
	fmt.Printf("arr2 (字面量): %v\n", arr2)
	fmt.Printf("arr3 (自动长度): %v, 长度: %d\n", arr3, len(arr3))
	fmt.Printf("arr4 (指定索引): %v\n", arr4)

	// 数组访问和修改
	arr2[0] = 100
	fmt.Printf("修改后的arr2: %v\n", arr2)

	// 数组遍历
	fmt.Println("遍历arr2:")
	for i := 0; i < len(arr2); i++ {
		fmt.Printf("  索引%d: %d\n", i, arr2[i])
	}

	// 使用range遍历
	fmt.Println("使用range遍历:")
	for index, value := range arr2 {
		fmt.Printf("  索引%d: %d\n", index, value)
	}
}

func arrayLimitations() {
	fmt.Println("\n=== 数组的限制 ===")

	// 数组长度是类型的一部分
	var arr5 [5]int
	var arr10 [10]int

	fmt.Printf("arr5类型: %T\n", arr5)
	fmt.Printf("arr10类型: %T\n", arr10)

	// arr5 = arr10 // 编译错误：类型不匹配
	fmt.Println("数组长度不同时不能赋值")

	// 数组是值类型，赋值会拷贝整个数组
	original := [3]int{1, 2, 3}
	copy := original
	copy[0] = 100

	fmt.Printf("原数组: %v\n", original)
	fmt.Printf("拷贝后: %v\n", copy)
	fmt.Println("数组是值类型，修改拷贝不影响原数组")
}

// 2. 切片 - 动态数组，Go的核心数据结构
// 类似Java的ArrayList，但更高效

func sliceBasics() {
	fmt.Println("\n=== 切片基础 ===")

	// 切片声明和初始化
	var slice1 []int               // nil切片
	slice2 := []int{}              // 空切片
	slice3 := []int{1, 2, 3, 4, 5} // 字面量初始化
	slice4 := make([]int, 5)       // make创建，长度5，容量5
	slice5 := make([]int, 3, 10)   // make创建，长度3，容量10

	fmt.Printf("slice1 (nil): %v, 长度: %d, 容量: %d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("slice2 (空): %v, 长度: %d, 容量: %d\n", slice2, len(slice2), cap(slice2))
	fmt.Printf("slice3 (字面量): %v, 长度: %d, 容量: %d\n", slice3, len(slice3), cap(slice3))
	fmt.Printf("slice4 (make): %v, 长度: %d, 容量: %d\n", slice4, len(slice4), cap(slice4))
	fmt.Printf("slice5 (make长度容量): %v, 长度: %d, 容量: %d\n", slice5, len(slice5), cap(slice5))

	// nil切片检查
	if slice1 == nil {
		fmt.Println("slice1 是 nil")
	}

	// 切片操作
	slice3[0] = 100
	fmt.Printf("修改后的slice3: %v\n", slice3)
}

func sliceOperations() {
	fmt.Println("\n=== 切片操作 ===")

	// append - 添加元素
	var numbers []int
	fmt.Printf("初始切片: %v (长度: %d, 容量: %d)\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 1)
	fmt.Printf("添加1: %v (长度: %d, 容量: %d)\n", numbers, len(numbers), cap(numbers))

	numbers = append(numbers, 2, 3, 4)
	fmt.Printf("添加2,3,4: %v (长度: %d, 容量: %d)\n", numbers, len(numbers), cap(numbers))

	// 添加另一个切片
	more := []int{5, 6, 7}
	numbers = append(numbers, more...)
	fmt.Printf("添加切片: %v (长度: %d, 容量: %d)\n", numbers, len(numbers), cap(numbers))

	// 切片的切片（子切片）
	sub1 := numbers[1:4]   // 索引1到3
	sub2 := numbers[:3]    // 从开始到索引2
	sub3 := numbers[3:]    // 从索引3到结束
	sub4 := numbers[1:4:6] // 索引1到3，容量到6

	fmt.Printf("原切片: %v\n", numbers)
	fmt.Printf("sub1 [1:4]: %v (长度: %d, 容量: %d)\n", sub1, len(sub1), cap(sub1))
	fmt.Printf("sub2 [:3]: %v (长度: %d, 容量: %d)\n", sub2, len(sub2), cap(sub2))
	fmt.Printf("sub3 [3:]: %v (长度: %d, 容量: %d)\n", sub3, len(sub3), cap(sub3))
	fmt.Printf("sub4 [1:4:6]: %v (长度: %d, 容量: %d)\n", sub4, len(sub4), cap(sub4))
}

func sliceMemoryBehavior() {
	fmt.Println("\n=== 切片内存行为 ===")

	// 切片共享底层数组
	original := []int{1, 2, 3, 4, 5}
	slice := original[1:4]

	fmt.Printf("原切片: %v\n", original)
	fmt.Printf("子切片: %v\n", slice)

	// 修改子切片影响原切片
	slice[0] = 100
	fmt.Printf("修改子切片后原切片: %v\n", original)
	fmt.Printf("子切片: %v\n", slice)

	// 容量扩展时会重新分配内存
	fmt.Println("\n容量扩展演示:")
	small := []int{1, 2}
	fmt.Printf("初始: %v (长度: %d, 容量: %d)\n", small, len(small), cap(small))

	// 添加元素直到超过容量
	for i := 3; i <= 10; i++ {
		small = append(small, i)
		fmt.Printf("添加%d: %v (长度: %d, 容量: %d)\n", i, small, len(small), cap(small))
	}
}

func sliceCopyAndClone() {
	fmt.Println("\n=== 切片拷贝和克隆 ===")

	source := []int{1, 2, 3, 4, 5}

	// 错误的拷贝方式
	wrongCopy := source
	wrongCopy[0] = 100
	fmt.Printf("错误拷贝 - 原切片: %v, 拷贝: %v\n", source, wrongCopy)

	// 正确的拷贝方式1：使用copy函数
	source = []int{1, 2, 3, 4, 5} // 重置
	correctCopy := make([]int, len(source))
	copy(correctCopy, source)
	correctCopy[0] = 200
	fmt.Printf("copy函数 - 原切片: %v, 拷贝: %v\n", source, correctCopy)

	// 正确的拷贝方式2：使用append
	source = []int{1, 2, 3, 4, 5} // 重置
	appendCopy := append([]int(nil), source...)
	appendCopy[0] = 300
	fmt.Printf("append拷贝 - 原切片: %v, 拷贝: %v\n", source, appendCopy)

	// 部分拷贝
	partial := make([]int, 3)
	copied := copy(partial, source)
	fmt.Printf("部分拷贝: %v, 拷贝了%d个元素\n", partial, copied)
}

// 3. 实际应用场景
func practicalExamples() {
	fmt.Println("\n=== 实际应用场景 ===")

	// 动态列表管理
	var todoList []string
	todoList = append(todoList, "学习Go语言")
	todoList = append(todoList, "写代码练习")
	todoList = append(todoList, "阅读文档")

	fmt.Println("待办事项:")
	for i, item := range todoList {
		fmt.Printf("  %d. %s\n", i+1, item)
	}

	// 删除中间元素（模拟删除第2项）
	indexToRemove := 1
	todoList = append(todoList[:indexToRemove], todoList[indexToRemove+1:]...)
	fmt.Println("删除第2项后:")
	for i, item := range todoList {
		fmt.Printf("  %d. %s\n", i+1, item)
	}

	// 在指定位置插入
	insertIndex := 1
	newItem := "复习昨天内容"
	todoList = append(todoList[:insertIndex], append([]string{newItem}, todoList[insertIndex:]...)...)
	fmt.Println("插入新项后:")
	for i, item := range todoList {
		fmt.Printf("  %d. %s\n", i+1, item)
	}
}

// 4. 字符串处理示例
func stringSliceExamples() {
	fmt.Println("\n=== 字符串切片示例 ===")

	// 字符串分割
	text := "Go,Python,Java,JavaScript"
	languages := strings.Split(text, ",")
	fmt.Printf("编程语言: %v\n", languages)

	// 字符串连接
	joined := strings.Join(languages, " | ")
	fmt.Printf("连接结果: %s\n", joined)

	// 字符串过滤
	var shortNames []string
	for _, lang := range languages {
		if len(lang) <= 4 {
			shortNames = append(shortNames, lang)
		}
	}
	fmt.Printf("短名称语言: %v\n", shortNames)

	// 字符串转换
	var upperLanguages []string
	for _, lang := range languages {
		upperLanguages = append(upperLanguages, strings.ToUpper(lang))
	}
	fmt.Printf("大写语言: %v\n", upperLanguages)
}

// 5. 数组和切片的排序
func sortingExamples() {
	fmt.Println("\n=== 排序示例 ===")

	// 整数排序
	numbers := []int{5, 2, 8, 1, 9, 3}
	fmt.Printf("排序前: %v\n", numbers)

	sort.Ints(numbers)
	fmt.Printf("升序排序: %v\n", numbers)

	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Printf("降序排序: %v\n", numbers)

	// 字符串排序
	words := []string{"banana", "apple", "cherry", "date"}
	fmt.Printf("排序前: %v\n", words)

	sort.Strings(words)
	fmt.Printf("字典序排序: %v\n", words)

	// 自定义排序
	people := []struct {
		Name string
		Age  int
	}{
		{"张三", 25},
		{"李四", 20},
		{"王五", 30},
	}

	fmt.Printf("排序前: %+v\n", people)

	// 按年龄排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})
	fmt.Printf("按年龄排序: %+v\n", people)

	// 按姓名排序
	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})
	fmt.Printf("按姓名排序: %+v\n", people)
}

// 6. 性能注意事项
func performanceConsiderations() {
	fmt.Println("\n=== 性能注意事项 ===")

	// 预分配切片容量
	fmt.Println("预分配容量的重要性:")

	// 不好的做法：频繁扩容
	var badSlice []int
	fmt.Printf("初始容量: %d\n", cap(badSlice))
	for i := 0; i < 1000; i++ {
		badSlice = append(badSlice, i)
		if i < 10 || (i+1)%100 == 0 {
			fmt.Printf("添加%d个元素后容量: %d\n", i+1, cap(badSlice))
		}
	}

	// 好的做法：预分配容量
	goodSlice := make([]int, 0, 1000)
	fmt.Printf("\n预分配1000容量，初始容量: %d\n", cap(goodSlice))
	for i := 0; i < 1000; i++ {
		goodSlice = append(goodSlice, i)
	}
	fmt.Printf("添加1000个元素后容量: %d (没有重新分配)\n", cap(goodSlice))

	// 切片泄漏示例
	fmt.Println("\n防止内存泄漏:")
	bigSlice := make([]int, 1000000)
	for i := range bigSlice {
		bigSlice[i] = i
	}

	// 不好的做法：保持对大切片的引用
	// smallPart := bigSlice[:10] // 这会导致整个bigSlice无法被回收

	// 好的做法：拷贝需要的部分
	smallPart := make([]int, 10)
	copy(smallPart, bigSlice[:10])
	fmt.Printf("正确提取小部分: %v\n", smallPart)
}

// 7. 多维数组和切片
func multiDimensionalExamples() {
	fmt.Println("\n=== 多维数组和切片 ===")

	// 二维数组
	var matrix [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			matrix[i][j] = i*3 + j + 1
		}
	}

	fmt.Println("3x3矩阵:")
	for i := 0; i < 3; i++ {
		fmt.Printf("  %v\n", matrix[i])
	}

	// 二维切片
	rows, cols := 4, 5
	matrix2D := make([][]int, rows)
	for i := range matrix2D {
		matrix2D[i] = make([]int, cols)
		for j := range matrix2D[i] {
			matrix2D[i][j] = i*cols + j + 1
		}
	}

	fmt.Printf("\n%dx%d动态矩阵:\n", rows, cols)
	for _, row := range matrix2D {
		fmt.Printf("  %v\n", row)
	}

	// 不规则二维切片
	jaggedArray := [][]string{
		{"A"},
		{"B", "C"},
		{"D", "E", "F"},
		{"G", "H", "I", "J"},
	}

	fmt.Println("\n不规则二维切片:")
	for i, row := range jaggedArray {
		fmt.Printf("  行%d: %v\n", i, row)
	}
}

func main() {
	fmt.Println("Go语言数据结构 - 数组和切片实践")
	fmt.Println("==================================")

	arrayBasics()
	arrayLimitations()
	sliceBasics()
	sliceOperations()
	sliceMemoryBehavior()
	sliceCopyAndClone()
	practicalExamples()
	stringSliceExamples()
	sortingExamples()
	performanceConsiderations()
	multiDimensionalExamples()

	fmt.Println("\n学习要点:")
	fmt.Println("1. 数组长度固定，是值类型；切片长度可变，是引用类型")
	fmt.Println("2. 切片有长度(len)和容量(cap)两个概念")
	fmt.Println("3. 子切片与原切片共享底层数组")
	fmt.Println("4. 使用make或append进行切片操作")
	fmt.Println("5. 注意切片的容量管理，避免频繁重新分配")
	fmt.Println("6. 使用copy()函数进行真正的切片拷贝")
}
