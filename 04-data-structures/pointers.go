package main

import (
	"fmt"
	"unsafe"
)

// 1. 指针基础 - Go中的内存地址
// Java：所有对象都是引用，自动内存管理
// Go：显式指针，手动控制内存访问，但有GC

func pointerBasics() {
	fmt.Println("=== 指针基础 ===")

	// 普通变量
	x := 42
	fmt.Printf("变量x的值: %d\n", x)
	fmt.Printf("变量x的地址: %p\n", &x)
	fmt.Printf("变量x的类型: %T\n", x)

	// 指针变量
	var p *int // 声明int类型的指针
	p = &x     // 获取x的地址
	fmt.Printf("指针p的值(地址): %p\n", p)
	fmt.Printf("指针p指向的值: %d\n", *p) // 解引用
	fmt.Printf("指针p的类型: %T\n", p)

	// 通过指针修改值
	*p = 100
	fmt.Printf("通过指针修改后x的值: %d\n", x)

	// 零值指针
	var nilPointer *int
	fmt.Printf("零值指针: %v\n", nilPointer)
	if nilPointer == nil {
		fmt.Println("指针是nil，不能解引用")
	}
}

func pointerOperations() {
	fmt.Println("\n=== 指针操作 ===")

	// 直接创建指针
	num1 := 10
	num2 := 20

	ptr1 := &num1
	ptr2 := &num2

	fmt.Printf("ptr1指向: %d (地址: %p)\n", *ptr1, ptr1)
	fmt.Printf("ptr2指向: %d (地址: %p)\n", *ptr2, ptr2)

	// 指针比较
	fmt.Printf("ptr1 == ptr2: %t\n", ptr1 == ptr2)
	fmt.Printf("*ptr1 == *ptr2: %t\n", *ptr1 == *ptr2)

	// 指针重新赋值
	ptr1 = ptr2
	fmt.Printf("ptr1重新赋值后指向: %d (地址: %p)\n", *ptr1, ptr1)
	fmt.Printf("现在ptr1 == ptr2: %t\n", ptr1 == ptr2)

	// new函数创建指针
	ptr3 := new(int) // 分配内存并返回指针
	*ptr3 = 30
	fmt.Printf("new创建的指针: %d (地址: %p)\n", *ptr3, ptr3)
}

// 2. 结构体指针
type Person struct {
	Name string
	Age  int
}

func (p *Person) Birthday() {
	p.Age++
	fmt.Printf("%s now is %d years old\n", p.Name, p.Age)
}

func (p Person) GetInfo() string {
	return fmt.Sprintf("Name: %s, Age: %d", p.Name, p.Age)
}

func structPointers() {
	fmt.Println("\n=== 结构体指针 ===")

	// 直接创建结构体
	person1 := Person{Name: "张三", Age: 25}
	fmt.Printf("person1: %+v\n", person1)

	// 获取结构体指针
	personPtr := &person1
	fmt.Printf("指针地址: %p\n", personPtr)
	fmt.Printf("通过指针访问: %+v\n", *personPtr)

	// 通过指针修改字段
	personPtr.Name = "张三丰" // 自动解引用，等同于 (*personPtr).Name
	personPtr.Age = 26
	fmt.Printf("修改后person1: %+v\n", person1)

	// 使用new创建结构体指针
	person2 := new(Person)
	person2.Name = "李四"
	person2.Age = 30
	fmt.Printf("new创建的结构体: %+v\n", *person2)

	// 结构体字面量直接创建指针
	person3 := &Person{Name: "王五", Age: 35}
	fmt.Printf("字面量创建指针: %+v\n", *person3)

	// 方法调用
	person3.Birthday()        // 指针接收者方法
	info := person3.GetInfo() // 值接收者方法，Go自动处理
	fmt.Printf("Info: %s\n", info)
}

// 3. 数组和切片指针
func arraySlicePointers() {
	fmt.Println("\n=== 数组和切片指针 ===")

	// 数组指针
	arr := [3]int{1, 2, 3}
	arrPtr := &arr

	fmt.Printf("数组: %v\n", arr)
	fmt.Printf("数组指针: %p\n", arrPtr)
	fmt.Printf("通过指针访问数组: %v\n", *arrPtr)

	// 修改数组元素
	arrPtr[0] = 100 // 等同于 (*arrPtr)[0] = 100
	fmt.Printf("修改后数组: %v\n", arr)

	// 切片本身就包含指向底层数组的指针
	slice := []int{10, 20, 30}
	fmt.Printf("切片: %v\n", slice)
	fmt.Printf("切片头信息 - 数据指针: %p, 长度: %d, 容量: %d\n",
		&slice[0], len(slice), cap(slice))

	// 切片指针
	slicePtr := &slice
	fmt.Printf("切片指针: %p\n", slicePtr)
	(*slicePtr)[0] = 200
	fmt.Printf("通过切片指针修改: %v\n", slice)
}

// 4. 函数参数中的指针
func modifyByValue(x int) {
	x = 100
	fmt.Printf("函数内修改值: %d\n", x)
}

func modifyByPointer(x *int) {
	*x = 100
	fmt.Printf("函数内通过指针修改: %d\n", *x)
}

func swapValues(a, b *int) {
	*a, *b = *b, *a
}

func functionPointers() {
	fmt.Println("\n=== 函数参数指针 ===")

	num := 50
	fmt.Printf("原始值: %d\n", num)

	// 值传递
	modifyByValue(num)
	fmt.Printf("值传递后: %d\n", num)

	// 指针传递
	modifyByPointer(&num)
	fmt.Printf("指针传递后: %d\n", num)

	// 交换两个值
	a, b := 10, 20
	fmt.Printf("交换前: a=%d, b=%d\n", a, b)
	swapValues(&a, &b)
	fmt.Printf("交换后: a=%d, b=%d\n", a, b)
}

// 5. 指针和内存分配
func memoryAllocation() {
	fmt.Println("\n=== 内存分配 ===")

	// 栈分配 vs 堆分配
	// Go编译器会自动决定变量分配在栈还是堆上

	// 局部变量（通常在栈上）
	localVar := 42
	fmt.Printf("局部变量地址: %p\n", &localVar)

	// 返回局部变量的指针（编译器会将其分配到堆上）
	createPointer := func() *int {
		x := 100 // 这个变量会逃逸到堆上
		return &x
	}

	heapPtr := createPointer()
	fmt.Printf("堆上变量地址: %p, 值: %d\n", heapPtr, *heapPtr)

	// 显式堆分配
	heapVar := new(int)
	*heapVar = 200
	fmt.Printf("new分配地址: %p, 值: %d\n", heapVar, *heapVar)

	// 大对象通常分配在堆上
	largeSlice := make([]int, 10000)
	fmt.Printf("大切片第一个元素地址: %p\n", &largeSlice[0])
}

// 6. 指针和接口
type Writer interface {
	Write(data string)
}

type FileWriter struct {
	filename string
}

func (f *FileWriter) Write(data string) {
	fmt.Printf("写入文件 %s: %s\n", f.filename, data)
}

func interfacePointers() {
	fmt.Println("\n=== 指针和接口 ===")

	// 指针实现接口
	fw := &FileWriter{filename: "test.txt"}

	var w Writer = fw
	w.Write("Hello, Go!")

	// 接口值包含类型信息和值
	fmt.Printf("接口类型: %T\n", w)
	fmt.Printf("接口值: %+v\n", w)

	// 类型断言获取原始指针
	if fileWriter, ok := w.(*FileWriter); ok {
		fmt.Printf("断言成功，文件名: %s\n", fileWriter.filename)
	}
}

// 7. 危险的指针操作 - unsafe包
func unsafePointers() {
	fmt.Println("\n=== 不安全指针操作 ===")

	// 警告：unsafe包的使用需要极其谨慎

	x := int64(100)
	fmt.Printf("原始值: %d\n", x)

	// 获取unsafe.Pointer
	ptr := unsafe.Pointer(&x)
	fmt.Printf("unsafe指针: %p\n", ptr)

	// 类型转换（危险！）
	bytePtr := (*byte)(ptr)
	fmt.Printf("作为byte查看第一个字节: %d\n", *bytePtr)

	// 指针运算（极其危险！）
	nextByte := (*byte)(unsafe.Pointer(uintptr(ptr) + 1))
	fmt.Printf("下一个字节: %d\n", *nextByte)

	fmt.Println("注意: unsafe操作绕过Go的类型安全，可能导致未定义行为")
}

// 8. 指针的实际应用场景
type LinkedListNode struct {
	Data int
	Next *LinkedListNode
}

type LinkedList struct {
	Head *LinkedListNode
	Size int
}

func (ll *LinkedList) Add(data int) {
	newNode := &LinkedListNode{Data: data}
	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
	ll.Size++
}

func (ll *LinkedList) Print() {
	current := ll.Head
	fmt.Print("链表: ")
	for current != nil {
		fmt.Printf("%d", current.Data)
		if current.Next != nil {
			fmt.Print(" -> ")
		}
		current = current.Next
	}
	fmt.Printf(" (长度: %d)\n", ll.Size)
}

func practicalExamples() {
	fmt.Println("\n=== 实际应用场景 ===")

	// 链表实现
	list := &LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Print()

	// 树结构
	type TreeNode struct {
		Value int
		Left  *TreeNode
		Right *TreeNode
	}

	root := &TreeNode{Value: 10}
	root.Left = &TreeNode{Value: 5}
	root.Right = &TreeNode{Value: 15}
	root.Left.Left = &TreeNode{Value: 3}
	root.Left.Right = &TreeNode{Value: 7}

	// 中序遍历
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node != nil {
			inorder(node.Left)
			fmt.Printf("%d ", node.Value)
			inorder(node.Right)
		}
	}

	fmt.Print("二叉树中序遍历: ")
	inorder(root)
	fmt.Println()
}

// 9. 指针性能和最佳实践
func performanceAndBestPractices() {
	fmt.Println("\n=== 性能和最佳实践 ===")

	// 大结构体使用指针传递
	type LargeStruct struct {
		Data [1000]int
		Name string
	}

	// 值传递（低效）
	processByValue := func(ls LargeStruct) {
		// 复制整个结构体，消耗内存和时间
		fmt.Printf("值传递处理，名称: %s\n", ls.Name)
	}

	// 指针传递（高效）
	processByPointer := func(ls *LargeStruct) {
		// 只传递8字节的指针
		fmt.Printf("指针传递处理，名称: %s\n", ls.Name)
	}

	large := LargeStruct{Name: "大结构体"}

	fmt.Println("对于大结构体:")
	processByValue(large)    // 复制1000个int + string
	processByPointer(&large) // 只传递指针

	fmt.Println("\n最佳实践:")
	fmt.Println("1. 大结构体使用指针传递")
	fmt.Println("2. 需要修改原值时使用指针")
	fmt.Println("3. 避免指针的指针，保持简单")
	fmt.Println("4. nil指针检查，避免panic")
	fmt.Println("5. 不要返回局部变量的指针（Go会自动处理逃逸）")
}

func nilPointerHandling() {
	fmt.Println("\n=== nil指针处理 ===")

	var p *int

	// 安全的nil检查
	if p != nil {
		fmt.Printf("指针值: %d\n", *p)
	} else {
		fmt.Println("指针是nil，不能解引用")
	}

	// 安全的方法调用
	type SafeStruct struct {
		value int
	}

	(func(s *SafeStruct) SafeMethod)()
	{
		if s != nil {
			fmt.Printf("安全调用，值: %d\n", s.value)
		} else {
			fmt.Println("接收者是nil")
		}
	}

	var s *SafeStruct
	s.SafeMethod() // 即使s是nil，也不会panic

	s = &SafeStruct{value: 42}
	s.SafeMethod()
}

func main() {
	fmt.Println("Go语言数据结构 - 指针实践")
	fmt.Println("===========================")

	pointerBasics()
	pointerOperations()
	structPointers()
	arraySlicePointers()
	functionPointers()
	memoryAllocation()
	interfacePointers()
	unsafePointers()
	practicalExamples()
	performanceAndBestPractices()
	nilPointerHandling()

	fmt.Println("\n学习要点:")
	fmt.Println("1. 指针存储变量的内存地址")
	fmt.Println("2. 使用&获取地址，*解引用指针")
	fmt.Println("3. 指针的零值是nil")
	fmt.Println("4. 结构体指针可以直接访问字段")
	fmt.Println("5. 大结构体使用指针传递提高性能")
	fmt.Println("6. Go有垃圾回收，无需手动释放内存")
	fmt.Println("7. 始终检查nil指针，避免panic")
	fmt.Println("8. unsafe包提供底层指针操作，但要谨慎使用")
}
