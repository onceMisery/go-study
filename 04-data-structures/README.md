# Go语言数据结构详解

## 学习目标
- 掌握Go语言的主要数据结构
- 理解Go与Java集合框架的差异
- 学会选择合适的数据结构
- 掌握数据结构的高级用法

## Java vs Go 数据结构对比

| 数据结构 | Java | Go | 主要差异 |
|----------|------|----|----------|
| 数组 | 引用类型，可变长度 | 值类型，固定长度 | 内存模型完全不同 |
| 动态数组 | ArrayList | slice | Go的slice更强大 |
| 映射 | HashMap/TreeMap | map | Go的map是内置类型 |
| 集合 | HashSet/TreeSet | 需要自实现 | Go没有内置set |
| 链表 | LinkedList | 需要自实现 | Go没有内置链表 |
| 栈/队列 | Stack/Queue | 需要自实现 | Go用slice实现 |

## 1. 数组 (Array)

Go的数组是值类型，长度固定，这与Java数组有根本区别。

### 数组声明和初始化
```go
// Java: int[] arr = new int[5];
var arr [5]int  // 声明长度为5的整数数组，零值为[0 0 0 0 0]

// 初始化
var arr1 = [5]int{1, 2, 3, 4, 5}
arr2 := [5]int{1, 2, 3, 4, 5}
arr3 := [...]int{1, 2, 3, 4, 5}  // 自动推断长度

// 指定索引初始化
arr4 := [5]int{0: 1, 2: 3, 4: 5}  // [1 0 3 0 5]
```

### 数组操作
```go
// 访问和修改
arr := [3]int{1, 2, 3}
fmt.Println(arr[0])  // 1
arr[0] = 10
fmt.Println(arr)     // [10 2 3]

// 数组长度
fmt.Println(len(arr))  // 3

// 遍历数组
for i := 0; i < len(arr); i++ {
    fmt.Printf("arr[%d] = %d\n", i, arr[i])
}

// 使用range
for index, value := range arr {
    fmt.Printf("索引：%d，值：%d\n", index, value)
}
```

### 数组是值类型
```go
// 重要：数组赋值会复制整个数组
arr1 := [3]int{1, 2, 3}
arr2 := arr1  // 复制整个数组
arr2[0] = 100
fmt.Println(arr1)  // [1 2 3] - 不受影响
fmt.Println(arr2)  // [100 2 3]

// 函数传参也会复制
func modifyArray(arr [3]int) {
    arr[0] = 999  // 不会影响原数组
}
```

## 2. 切片 (Slice)

切片是Go最重要的数据结构，类似Java的ArrayList但更强大。

### 切片基础
```go
// Java: List<Integer> list = new ArrayList<>();
var slice []int  // 声明切片，零值为nil

// 初始化
slice1 := []int{1, 2, 3, 4, 5}
slice2 := make([]int, 5)      // 长度为5，容量为5
slice3 := make([]int, 3, 10)  // 长度为3，容量为10
```

### 切片操作
```go
// 添加元素
slice := []int{1, 2, 3}
slice = append(slice, 4)       // [1 2 3 4]
slice = append(slice, 5, 6, 7) // [1 2 3 4 5 6 7]

// 合并切片
slice2 := []int{8, 9, 10}
slice = append(slice, slice2...)  // [1 2 3 4 5 6 7 8 9 10]

// 切片操作
numbers := []int{0, 1, 2, 3, 4, 5}
fmt.Println(numbers[1:4])   // [1 2 3] - 左闭右开
fmt.Println(numbers[:3])    // [0 1 2] - 从开始到索引3
fmt.Println(numbers[2:])    // [2 3 4 5] - 从索引2到结束
fmt.Println(numbers[:])     // [0 1 2 3 4 5] - 完整切片
```

### 切片的内部结构
```go
// 切片包含：指针、长度、容量
slice := make([]int, 3, 5)
fmt.Printf("长度：%d，容量：%d\n", len(slice), cap(slice))

// 扩容机制
slice = []int{1, 2, 3}
fmt.Printf("容量：%d\n", cap(slice))  // 3
slice = append(slice, 4)
fmt.Printf("容量：%d\n", cap(slice))  // 6 (通常翻倍)
```

### 切片拷贝
```go
// 浅拷贝：共享底层数组
slice1 := []int{1, 2, 3, 4, 5}
slice2 := slice1[1:4]  // [2 3 4]
slice2[0] = 100
fmt.Println(slice1)    // [1 100 3 4 5] - 受影响

// 深拷贝
slice3 := make([]int, len(slice1))
copy(slice3, slice1)   // 复制所有元素
slice3[0] = 999
fmt.Println(slice1)    // [1 100 3 4 5] - 不受影响
```

## 3. 映射 (Map)

Go的映射类似Java的HashMap。

### 映射基础
```go
// Java: Map<String, Integer> map = new HashMap<>();
var userAges map[string]int  // 声明map，零值为nil

// 初始化
userAges = make(map[string]int)
userAges["Alice"] = 25
userAges["Bob"] = 30

// 直接初始化
userAges2 := map[string]int{
    "Alice": 25,
    "Bob":   30,
    "Carol": 28,
}
```

### 映射操作
```go
userAges := map[string]int{
    "Alice": 25,
    "Bob":   30,
}

// 访问
age := userAges["Alice"]          // 25
ageNotExist := userAges["David"]  // 0 (零值)

// 检查键是否存在
age, exists := userAges["Alice"]
if exists {
    fmt.Printf("Alice的年龄是%d\n", age)
}

// 添加和修改
userAges["Carol"] = 28  // 添加
userAges["Alice"] = 26  // 修改

// 删除
delete(userAges, "Bob")

// 遍历
for name, age := range userAges {
    fmt.Printf("%s: %d岁\n", name, age)
}
```

### 映射的高级用法
```go
// 嵌套映射
students := map[string]map[string]int{
    "张三": {"数学": 90, "英语": 85},
    "李四": {"数学": 88, "英语": 92},
}

// 使用结构体作为值
type Student struct {
    Name string
    Age  int
}

studentMap := map[int]Student{
    1001: {Name: "张三", Age: 20},
    1002: {Name: "李四", Age: 21},
}

// 统计字符出现次数
text := "hello world"
charCount := make(map[rune]int)
for _, char := range text {
    charCount[char]++
}
```

## 4. 结构体 (Struct)

结构体是Go中组织数据的主要方式，替代Java的类。

### 结构体定义
```go
// Java: class Person { ... }
type Person struct {
    Name string
    Age  int
    Email string
}

// 嵌套结构体
type Address struct {
    Street  string
    City    string
    Country string
}

type Employee struct {
    Person          // 匿名字段，实现继承效果
    ID       int
    Position string
    Address  Address  // 命名字段
}
```

### 结构体操作
```go
// 创建结构体
var p1 Person  // 零值：{Name: "", Age: 0, Email: ""}

// 初始化
p2 := Person{
    Name:  "张三",
    Age:   25,
    Email: "zhangsan@example.com",
}

p3 := Person{"李四", 30, "lisi@example.com"}  // 按顺序

// 访问和修改
fmt.Println(p2.Name)  // 张三
p2.Age = 26

// 结构体指针
p4 := &Person{Name: "王五", Age: 28}
fmt.Println(p4.Name)   // 可以直接访问，Go自动解引用
fmt.Println((*p4).Name) // 显式解引用
```

## 5. 接口 (Interface)

Go的接口实现了多态，但与Java接口不同。

### 接口定义
```go
// 定义接口
type Writer interface {
    Write([]byte) (int, error)
}

type Reader interface {
    Read([]byte) (int, error)
}

// 组合接口
type ReadWriter interface {
    Reader
    Writer
}
```

### 接口实现
```go
// 文件类型
type File struct {
    name string
}

// 实现Writer接口（隐式实现）
func (f *File) Write(data []byte) (int, error) {
    fmt.Printf("写入文件%s: %s\n", f.name, string(data))
    return len(data), nil
}

// 使用接口
func writeData(w Writer, data []byte) {
    w.Write(data)
}

func main() {
    file := &File{name: "test.txt"}
    writeData(file, []byte("Hello, Go!"))
}
```

## 6. 通道 (Channel)

通道是Go特有的数据结构，用于goroutine通信。

### 通道基础
```go
// 创建通道
ch := make(chan int)        // 无缓冲通道
ch2 := make(chan int, 10)   // 有缓冲通道

// 发送和接收
go func() {
    ch <- 42  // 发送
}()
value := <-ch  // 接收

// 关闭通道
close(ch)

// 检查通道是否关闭
value, ok := <-ch
if !ok {
    fmt.Println("通道已关闭")
}
```

## 7. 自定义数据结构

### 实现栈
```go
type Stack []int

func (s *Stack) Push(value int) {
    *s = append(*s, value)
}

func (s *Stack) Pop() (int, bool) {
    if len(*s) == 0 {
        return 0, false
    }
    index := len(*s) - 1
    value := (*s)[index]
    *s = (*s)[:index]
    return value, true
}

func (s *Stack) IsEmpty() bool {
    return len(*s) == 0
}
```

### 实现集合
```go
type Set map[string]bool

func NewSet() Set {
    return make(Set)
}

func (s Set) Add(item string) {
    s[item] = true
}

func (s Set) Remove(item string) {
    delete(s, item)
}

func (s Set) Contains(item string) bool {
    return s[item]
}

func (s Set) ToSlice() []string {
    result := make([]string, 0, len(s))
    for item := range s {
        result = append(result, item)
    }
    return result
}
```

## 性能对比和选择指南

| 操作 | 数组 | 切片 | 映射 | 建议 |
|------|------|------|------|------|
| 按索引访问 | O(1) | O(1) | O(1) | 都很快 |
| 插入 | 不支持 | O(1)* | O(1) | 切片或映射 |
| 删除 | 不支持 | O(n) | O(1) | 看需求 |
| 查找 | O(n) | O(n) | O(1) | 需要快速查找用映射 |
| 内存使用 | 最少 | 中等 | 最多 | 根据场景选择 |

## 实践任务

1. **学生管理系统**：使用结构体和映射
2. **LRU缓存**：结合映射和双向链表
3. **词频统计**：使用映射统计文本中单词频率
4. **图的表示**：使用邻接表表示图结构

## 常见陷阱

1. **切片共享底层数组**
2. **映射并发访问不安全**
3. **接口nil判断**
4. **结构体比较**

## 下一步

掌握了数据结构后，建议学习：
- 方法和接口的深入使用
- 并发编程（goroutine和channel）
- 错误处理和包管理

## 参考资源

- [Go Tour - 数据结构](https://tour.golang.org/moretypes/1)
- [Effective Go - 数据](https://golang.org/doc/effective_go#data)
- [Go Slices: usage and internals](https://golang.org/blog/slices-intro) 