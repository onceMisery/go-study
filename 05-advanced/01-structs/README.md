# Structs - 结构体

## 📋 学习目标
- 掌握Go语言结构体的定义和使用
- 理解方法的定义和接收者
- 学会结构体嵌套和组合
- 对比Go结构体与Java类的差异
- 掌握结构体的最佳实践

## 🏗️ 结构体基础

### Java类 vs Go结构体对比

**Java类定义:**
```java
public class Person {
    private String name;
    private int age;
    private String email;
    
    // 构造函数
    public Person(String name, int age, String email) {
        this.name = name;
        this.age = age;
        this.email = email;
    }
    
    // Getter和Setter
    public String getName() { return name; }
    public void setName(String name) { this.name = name; }
    public int getAge() { return age; }
    public void setAge(int age) { this.age = age; }
    
    // 实例方法
    public void introduce() {
        System.out.println("我是" + name + "，今年" + age + "岁");
    }
    
    // 静态方法
    public static Person createDefault() {
        return new Person("无名", 0, "");
    }
}
```

**Go结构体定义:**
```go
package main

import "fmt"

// 结构体定义
type Person struct {
    Name  string  // 公开字段（首字母大写）
    Age   int     // 公开字段
    email string  // 私有字段（首字母小写）
}

// 方法定义（值接收者）
func (p Person) Introduce() {
    fmt.Printf("我是%s，今年%d岁\n", p.Name, p.Age)
}

// 方法定义（指针接收者）
func (p *Person) SetEmail(email string) {
    p.email = email
}

func (p Person) GetEmail() string {
    return p.email
}

// 构造函数（约定俗成）
func NewPerson(name string, age int, email string) *Person {
    return &Person{
        Name:  name,
        Age:   age,
        email: email,
    }
}

// 包级别函数（类似Java静态方法）
func CreateDefaultPerson() Person {
    return Person{
        Name:  "无名",
        Age:   0,
        email: "",
    }
}
```

## 🎯 结构体初始化

### 多种初始化方式
```go
type Book struct {
    Title  string
    Author string
    Price  float64
    Pages  int
}

func main() {
    // 方式1：字面量初始化
    book1 := Book{
        Title:  "Go语言学习",
        Author: "张三",
        Price:  99.9,
        Pages:  300,
    }
    
    // 方式2：按顺序初始化
    book2 := Book{"Java进阶", "李四", 89.9, 450}
    
    // 方式3：部分初始化（其他字段为零值）
    book3 := Book{
        Title: "Python入门",
        Price: 79.9,
    }
    
    // 方式4：使用new关键字
    book4 := new(Book)  // 返回*Book
    book4.Title = "C++实战"
    
    // 方式5：构造函数
    book5 := NewBook("Go并发编程", "王五", 120.0, 400)
}

func NewBook(title, author string, price float64, pages int) *Book {
    return &Book{
        Title:  title,
        Author: author,
        Price:  price,
        Pages:  pages,
    }
}
```

**Java对比:**
```java
// Java对象创建
Book book1 = new Book("Go语言学习", "张三", 99.9, 300);
Book book2 = new Book();  // 需要默认构造函数
book2.setTitle("Java进阶");
book2.setPrice(89.9);
```

## 🔧 方法和接收者

### 值接收者 vs 指针接收者
```go
type Counter struct {
    count int
}

// 值接收者 - 不会修改原始结构体
func (c Counter) GetCount() int {
    return c.count
}

// 值接收者 - 无法修改原始结构体
func (c Counter) IncrementValue() {
    c.count++  // 只修改副本
}

// 指针接收者 - 可以修改原始结构体
func (c *Counter) Increment() {
    c.count++
}

// 指针接收者 - 避免大结构体的复制
func (c *Counter) Reset() {
    c.count = 0
}

func main() {
    counter := Counter{count: 0}
    
    fmt.Println(counter.GetCount()) // 0
    
    counter.IncrementValue()
    fmt.Println(counter.GetCount()) // 0 (没有改变)
    
    counter.Increment()
    fmt.Println(counter.GetCount()) // 1 (已改变)
}
```

**选择指导原则:**
- 需要修改结构体：使用指针接收者
- 结构体很大：使用指针接收者（避免复制）
- 结构体很小且不需修改：使用值接收者
- 保持一致性：同一类型的方法使用同一种接收者

**Java对比:**
```java
// Java中所有对象都是引用传递
public class Counter {
    private int count;
    
    public int getCount() { return count; }
    public void increment() { count++; }  // 总是修改原对象
}
```

## 🧩 结构体嵌套和组合

### 嵌套结构体
```go
type Address struct {
    Street  string
    City    string
    Country string
}

type Person struct {
    Name    string
    Age     int
    Address Address  // 嵌套结构体
}

func main() {
    person := Person{
        Name: "张三",
        Age:  30,
        Address: Address{
            Street:  "中山路123号",
            City:    "杭州",
            Country: "中国",
        },
    }
    
    fmt.Printf("%s住在%s%s%s\n", 
        person.Name, 
        person.Address.Country,
        person.Address.City,
        person.Address.Street)
}
```

### 匿名嵌入（组合）
```go
type Animal struct {
    Name string
    Age  int
}

func (a Animal) Speak() {
    fmt.Printf("%s发出声音\n", a.Name)
}

func (a Animal) Sleep() {
    fmt.Printf("%s在睡觉\n", a.Name)
}

// 嵌入Animal（组合而非继承）
type Dog struct {
    Animal  // 匿名嵌入
    Breed string
}

// Dog特有的方法
func (d Dog) Bark() {
    fmt.Printf("%s在汪汪叫\n", d.Name)
}

// 重写父类方法
func (d Dog) Speak() {
    d.Bark()  // 调用Dog特有的行为
}

func main() {
    dog := Dog{
        Animal: Animal{
            Name: "旺财",
            Age:  3,
        },
        Breed: "拉布拉多",
    }
    
    // 可以直接访问嵌入字段
    fmt.Println(dog.Name)    // 直接访问，等同于dog.Animal.Name
    fmt.Println(dog.Age)     // 直接访问
    fmt.Println(dog.Breed)   // Dog自己的字段
    
    // 调用方法
    dog.Speak()  // 调用Dog重写的方法
    dog.Sleep()  // 调用Animal的方法
    dog.Bark()   // 调用Dog特有的方法
}
```

**Java继承对比:**
```java
// Java使用继承
class Animal {
    protected String name;
    protected int age;
    
    public void speak() {
        System.out.println(name + "发出声音");
    }
    
    public void sleep() {
        System.out.println(name + "在睡觉");
    }
}

class Dog extends Animal {
    private String breed;
    
    public void bark() {
        System.out.println(name + "在汪汪叫");
    }
    
    @Override
    public void speak() {
        bark();  // 重写父类方法
    }
}
```

## 🏷️ 结构体标签

### 用于JSON序列化
```go
import (
    "encoding/json"
    "fmt"
)

type User struct {
    ID       int    `json:"id"`
    Name     string `json:"name"`
    Email    string `json:"email,omitempty"`  // 空值时省略
    Password string `json:"-"`                // 不序列化
    Age      int    `json:"age,string"`       // 转为字符串
}

func main() {
    user := User{
        ID:       1,
        Name:     "张三",
        Email:    "",
        Password: "secret123",
        Age:      25,
    }
    
    // 序列化为JSON
    jsonData, _ := json.Marshal(user)
    fmt.Println(string(jsonData))
    // 输出: {"id":1,"name":"张三","age":"25"}
    
    // 从JSON反序列化
    jsonStr := `{"id":2,"name":"李四","email":"lisi@example.com","age":"30"}`
    var newUser User
    json.Unmarshal([]byte(jsonStr), &newUser)
    fmt.Printf("%+v\n", newUser)
}
```

**Java注解对比:**
```java
// Java使用注解
@Entity
@Table(name = "users")
public class User {
    @Id
    @JsonProperty("id")
    private int id;
    
    @JsonProperty("name")
    private String name;
    
    @JsonProperty(value = "email")
    @JsonInclude(JsonInclude.Include.NON_EMPTY)
    private String email;
    
    @JsonIgnore
    private String password;
}
```

## 🎭 接口实现

### 隐式接口实现
```go
// 定义接口
type Speaker interface {
    Speak() string
}

type Walker interface {
    Walk() string
}

// 组合接口
type Animal interface {
    Speaker
    Walker
}

// 结构体
type Cat struct {
    Name string
}

// 实现接口方法（隐式实现）
func (c Cat) Speak() string {
    return c.Name + "说：喵喵"
}

func (c Cat) Walk() string {
    return c.Name + "优雅地走着"
}

type Robot struct {
    Model string
}

func (r Robot) Speak() string {
    return r.Model + "说：beep beep"
}

func (r Robot) Walk() string {
    return r.Model + "机械地行走"
}

// 使用接口
func demonstrateAnimal(a Animal) {
    fmt.Println(a.Speak())
    fmt.Println(a.Walk())
}

func main() {
    cat := Cat{Name: "咪咪"}
    robot := Robot{Model: "T-800"}
    
    // 都实现了Animal接口
    demonstrateAnimal(cat)
    demonstrateAnimal(robot)
}
```

## 📝 实践任务

### 任务1: 基础结构体
1. 定义学生管理系统的结构体
2. 实现方法和构造函数
3. 对比Java类的实现

### 任务2: 组合和嵌入
1. 设计形状类层次结构
2. 使用组合替代继承
3. 实现多层嵌套

### 任务3: 接口和序列化
1. 设计通用的数据处理接口
2. 实现JSON序列化
3. 使用结构体标签

## 🎯 学习要点

### Go结构体特点
1. **组合优于继承**: 使用嵌入实现代码复用
2. **方法接收者**: 值接收者和指针接收者的选择
3. **隐式接口**: 不需要显式声明实现接口
4. **结构体标签**: 元数据支持

### 与Java的主要差异
1. **无类概念**: 结构体+方法而非类
2. **无继承**: 使用组合和嵌入
3. **无构造函数**: 使用工厂函数
4. **隐式接口**: 鸭子类型，无需implements

## 🎯 下一步
- 学习接口和多态的深入应用
- 理解错误处理最佳实践
- 掌握并发编程基础

## 📚 参考资源
- [Go语言规范 - 结构体](https://golang.org/ref/spec#Struct_types)
- [Go语言规范 - 方法](https://golang.org/ref/spec#Method_declarations)
- [Effective Go - 结构体](https://golang.org/doc/effective_go.html#composite_literals) 