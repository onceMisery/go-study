package main

import (
	"fmt"
	"time"
)

// Person 结构体 - 类似Java中的类
// 在Go中，结构体字段首字母大写表示公共访问权限
type Person struct {
	Name     string // 公共字段 (类似Java的public)
	age      int    // 私有字段 (类似Java的private)
	Email    string
	Birthday time.Time
}

// NewPerson 构造函数模式 - Go没有构造函数，使用函数返回结构体指针
// 类似Java: public Person(String name, int age, String email)
func NewPerson(name string, age int, email string) *Person {
	return &Person{
		Name:     name,
		age:      age,
		Email:    email,
		Birthday: time.Now().AddDate(-age, 0, 0), // 估算生日
	}
}

// GetAge 获取年龄 - 类似Java的getter方法
// 在Go中，方法接收者写在func关键字后面
func (p *Person) GetAge() int {
	return p.age
}

// SetAge 设置年龄 - 类似Java的setter方法
// 使用指针接收者可以修改结构体
func (p *Person) SetAge(age int) {
	if age >= 0 && age <= 150 {
		p.age = age
	}
}

// Greet 问候方法 - 展示方法定义
func (p *Person) Greet() string {
	return fmt.Sprintf("你好，我是 %s，今年 %d 岁", p.Name, p.age)
}

// IsAdult 判断是否成年 - 值接收者示例
// 值接收者不能修改结构体，但性能更好
func (p Person) IsAdult() bool {
	return p.age >= 18
}

// String 实现Stringer接口 - 类似Java的toString()
func (p *Person) String() string {
	return fmt.Sprintf("Person{Name: %s, Age: %d, Email: %s}",
		p.Name, p.age, p.Email)
}

// 演示嵌入字段（匿名字段）
type Address struct {
	Street   string
	City     string
	Country  string
	PostCode string
}

// PersonWithAddress 展示结构体嵌入 - 类似继承但是组合
type PersonWithAddress struct {
	Person      // 嵌入Person结构体
	Address     // 嵌入Address结构体
	PhoneNumber string
}

// GetFullInfo 获取完整信息
func (p *PersonWithAddress) GetFullInfo() string {
	return fmt.Sprintf("%s\n地址: %s, %s, %s %s\n电话: %s",
		p.Greet(), p.Street, p.City, p.Country, p.PostCode, p.PhoneNumber)
}

func main() {
	// 创建Person实例
	person1 := NewPerson("张三", 25, "zhangsan@example.com")
	fmt.Println("=== 基础结构体操作 ===")
	fmt.Println(person1)
	fmt.Println(person1.Greet())
	fmt.Printf("是否成年: %t\n", person1.IsAdult())

	// 修改年龄
	person1.SetAge(30)
	fmt.Printf("修改后年龄: %d\n", person1.GetAge())

	// 字面量初始化
	person2 := &Person{
		Name:  "李四",
		age:   22,
		Email: "lisi@example.com",
	}
	fmt.Println("\n=== 字面量初始化 ===")
	fmt.Println(person2.Greet())

	// 结构体嵌入示例
	fmt.Println("\n=== 结构体嵌入 (组合) ===")
	personWithAddr := &PersonWithAddress{
		Person: Person{
			Name:  "王五",
			age:   28,
			Email: "wangwu@example.com",
		},
		Address: Address{
			Street:   "长安街1号",
			City:     "北京",
			Country:  "中国",
			PostCode: "100000",
		},
		PhoneNumber: "13800138000",
	}

	fmt.Println(personWithAddr.GetFullInfo())
	// 直接访问嵌入字段的方法
	fmt.Printf("嵌入字段调用: %s\n", personWithAddr.Greet())
}
