# Hello World - Go语言入门

## 📋 学习目标
- 安装和配置Go开发环境
- 理解Go项目结构和模块系统
- 编写第一个Go程序
- 掌握基本的Go命令行工具

## 🔧 环境准备

### 1. 安装Go
- 访问 [golang.org](https://golang.org/dl/) 下载Go安装包
- Windows: 下载.msi文件双击安装
- 验证安装: `go version`

### 2. 设置环境变量
```bash
# Windows (PowerShell)
$env:GOPATH = "D:\go"
$env:GOROOT = "C:\Go"

# 验证环境
go env GOPATH
go env GOROOT
```

### 3. 开发工具配置
- **VS Code**: 安装Go扩展
- **GoLand**: JetBrains的Go IDE
- **Vim/Neovim**: 配置go-vim插件

## 📚 Go项目结构

### Java项目结构对比
**Java (Maven):**
```
project/
├── src/
│   ├── main/
│   │   └── java/
│   │       └── com/example/
│   │           └── Main.java
│   └── test/
├── pom.xml
└── target/
```

**Go (Modules):**
```
project/
├── go.mod          # 依赖管理（类似pom.xml）
├── go.sum          # 依赖校验
├── main.go         # 主程序入口
├── internal/       # 私有包
├── pkg/           # 公共包
└── cmd/           # 可执行程序
```

## 🚀 第一个Go程序

### 创建项目
```bash
# 1. 创建项目目录
mkdir hello-world-go
cd hello-world-go

# 2. 初始化Go模块
go mod init hello-world

# 3. 创建主程序文件
# 见 main.go 文件
```

### Java vs Go 对比

**Java版本:**
```java
package com.example;

public class HelloWorld {
    public static void main(String[] args) {
        System.out.println("Hello, World!");
        
        // 变量声明
        String name = "Java";
        int year = 2024;
        
        System.out.printf("Hello from %s in %d!%n", name, year);
    }
}
```

**Go版本:**
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
    
    // 变量声明
    var name string = "Go"
    var year int = 2024
    
    // 或者使用简短声明
    name := "Go"
    year := 2024
    
    fmt.Printf("Hello from %s in %d!\n", name, year)
}
```

## 🔍 关键差异分析

### 1. 包声明
- **Java**: `package com.example;` (域名反转)
- **Go**: `package main` (简单包名，main表示可执行程序)

### 2. 导入语句
- **Java**: `import java.util.*;`
- **Go**: `import "fmt"` (导入标准库或模块路径)

### 3. 主函数
- **Java**: `public static void main(String[] args)`
- **Go**: `func main()` (更简洁，包级别函数)

### 4. 输出语句
- **Java**: `System.out.println()`
- **Go**: `fmt.Println()` (需要导入fmt包)

### 5. 变量声明
- **Java**: 类型在前 `String name = "value"`
- **Go**: 类型在后 `var name string = "value"` 或 `name := "value"`

## 📝 实践任务

### 任务1: 基础Hello World
1. 创建基础的Hello World程序
2. 使用不同的输出格式
3. 添加命令行参数处理

### 任务2: 项目结构
1. 创建符合Go约定的项目结构
2. 理解go.mod文件的作用
3. 实践包的导入和使用

### 任务3: 对比练习
1. 将一个简单的Java程序转换为Go
2. 记录转换过程中的差异
3. 总结Go的简洁性体现

## 🛠️ Go命令行工具

### 基本命令
```bash
# 运行程序
go run main.go

# 构建程序
go build                    # 生成可执行文件
go build -o hello.exe      # 指定输出文件名

# 安装程序到GOPATH/bin
go install

# 格式化代码
go fmt

# 检查代码
go vet

# 测试
go test

# 下载依赖
go mod download

# 整理依赖
go mod tidy
```

### Java vs Go 工具对比
| 功能 | Java | Go |
|------|------|-----|
| 编译 | javac | go build |
| 运行 | java | go run |
| 包管理 | Maven/Gradle | go mod |
| 格式化 | IDE插件 | go fmt |
| 代码检查 | SpotBugs/PMD | go vet |

## 💡 学习要点

### Go的优势
1. **编译速度快**: 比Java编译快很多
2. **部署简单**: 单一可执行文件，无需JVM
3. **语法简洁**: 更少的样板代码
4. **内置工具**: 格式化、测试、文档生成等

### 需要适应的地方
1. **包管理**: 不同于Maven/Gradle的依赖管理
2. **项目结构**: 更扁平的包结构
3. **错误处理**: 没有异常机制
4. **面向对象**: 没有类，使用结构体和接口

## 🎯 下一步
完成Hello World后，继续学习：
- 变量和数据类型
- 函数定义和调用
- 包的创建和使用
- 基本的输入输出操作

## 📚 参考资源
- [Go官方教程](https://golang.org/doc/tutorial/)
- [Go by Example](https://gobyexample.com/)
- [A Tour of Go](https://tour.golang.org/)
- [Effective Go](https://golang.org/doc/effective_go.html) 