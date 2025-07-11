# Go语言学习项目 - 从Java到Go的进阶之路

## 项目概述

本项目专为有Java开发经验的工程师设计，通过系统性的学习路径和实践项目，帮助你快速掌握Go语言，并成长为高级Go开发工程师。

## 技术选型

### 核心技术栈
- **Go版本**: Go 1.21+
- **包管理**: Go Modules
- **数据库**: MySQL
- **时间处理**: time包（对应Java的LocalDate/LocalDateTime）

### 框架对比
| Java生态 | Go生态 | 用途 |
|---------|--------|------|
| Spring Boot | Gin | Web开发框架 |
| Hibernate/JPA | GORM | ORM框架 |
| Spring DI | Wire | 依赖注入 |
| Spring Config | Viper | 配置管理 |
| JUnit | testing + Testify | 单元测试 |
| Spring Cloud | Go-kit/Kratos | 微服务框架 |

## 项目架构

```
go-demo/
├── README.md                 # 项目总览
├── todo.md                   # 学习进度跟踪
├── data/                     # 学习过程数据存储
├── docs/                     # 学习文档
│   ├── java-vs-go.md        # Java与Go详细对比
│   └── learning-path.md     # 完整学习路线
├── 01-basics/               # 基础语法（1-2周）
│   ├── 01-hello-world/      # Hello World和环境搭建
│   ├── 02-variables/        # 变量声明和作用域
│   ├── 03-data-types/       # 基本数据类型
│   └── 04-operators/        # 运算符
├── 02-control-flow/         # 控制流（1周）
│   ├── 01-conditions/       # 条件语句
│   ├── 02-loops/           # 循环语句
│   └── 03-switch/          # Switch语句
├── 03-functions/            # 函数（1周）
│   ├── function-basics/     # 函数基础
│   ├── advanced-functions/ # 高级函数特性
│   └── closures/           # 闭包
├── 04-data-structures/      # 数据结构（2周）
│   ├── 01-arrays/          # 数组
│   ├── 02-slices/          # 切片（Go特有）
│   ├── 03-maps/            # 映射
│   └── 04-pointers/        # 指针
├── 05-advanced/             # 高级特性（3-4周）
│   ├── 01-structs/         # 结构体
│   ├── 02-interfaces/      # 接口和多态
│   ├── 03-error-handling/  # 错误处理
│   ├── 04-concurrency/     # 并发编程（goroutine、channel）
│   └── 05-file-io/         # 文件操作
├── 06-frameworks/           # 框架学习（4-6周）
│   ├── 01-gin/             # Gin Web框架
│   ├── 02-gorm/            # GORM数据库操作
│   ├── 03-microservices/   # 微服务架构
│   └── 04-testing/         # 测试框架
└── 07-projects/             # 实战项目（6-8周）
    ├── 01-rest-api/        # RESTful API项目
    ├── 02-web-app/         # 完整Web应用
    └── 03-microservice/    # 微服务项目
```

## 学习路线

### 阶段一：基础入门（4-5周）
1. **环境搭建和基础语法**（1-2周）
   - Go环境安装和配置
   - Hello World程序
   - 变量、常量、数据类型
   - 运算符和表达式

2. **控制流和函数**（2周）
   - 条件语句、循环语句
   - 函数定义和调用
   - 函数高级特性

3. **数据结构**（1-2周）
   - 数组和切片
   - 映射和指针
   - 内存管理

### 阶段二：进阶特性（4-5周）
1. **面向对象编程**（2周）
   - 结构体和方法
   - 接口和多态
   - 组合vs继承

2. **错误处理和并发**（2-3周）
   - Go的错误处理模式
   - Goroutine和Channel
   - 并发模式和最佳实践

### 阶段三：框架和生态（4-6周）
1. **Web开发框架**（2-3周）
   - Gin框架基础
   - 中间件和路由
   - RESTful API设计

2. **数据库和ORM**（1-2周）
   - GORM使用
   - 数据库连接池
   - 事务处理

3. **微服务和测试**（1-2周）
   - 微服务架构
   - 单元测试和集成测试

### 阶段四：实战项目（6-8周）
1. **RESTful API项目**（2-3周）
2. **完整Web应用**（2-3周）
3. **微服务项目**（2-3周）

## 学习建议

### 对于Java开发者的建议
1. **思维转换**：Go更注重简洁性和性能，少用继承多用组合
2. **错误处理**：适应Go的显式错误处理，而不是异常机制
3. **并发模型**：理解goroutine和channel，这是Go的核心优势
4. **内存管理**：虽然有GC，但要理解指针和内存布局

### 性能和维护性注意事项
- **性能优化**：关注goroutine泄漏、内存分配、GC压力
- **代码维护**：使用接口提高可测试性，合理使用组合模式
- **边界条件**：处理nil指针、空切片、channel关闭等边界情况
- **监控告警**：集成pprof性能分析和日志监控

## 快速开始

1. 确保已安装Go 1.21+
2. 克隆本项目
3. 按照目录顺序学习
4. 每完成一个模块，更新todo.md进度
5. 实践每个示例代码
6. 对比Java实现思考差异

## 学习资源

- [Go官方文档](https://golang.org/doc/)
- [Go语言圣经](https://gopl.io/)
- [Effective Go](https://golang.org/doc/effective_go.html)
- 本项目提供的对比文档和实践项目

开始你的Go语言学习之旅吧！🚀 