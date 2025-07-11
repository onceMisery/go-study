# Gin Web Framework - Gin Web框架

## 📋 学习目标
- 掌握Gin框架的基本使用
- 理解路由和中间件机制
- 学会处理HTTP请求和响应
- 对比Spring Boot的实现方式
- 掌握RESTful API开发

## 🚀 Gin框架介绍

### 什么是Gin
Gin是Go语言的一个高性能HTTP Web框架，类似于Java的Spring Boot，但更轻量级。

**特点:**
- 高性能 (比其他Go框架快40倍)
- 支持中间件
- 崩溃处理
- JSON验证
- 路由组
- 错误管理
- 内置渲染

### Spring Boot vs Gin对比

**Spring Boot应用:**
```java
@RestController
@RequestMapping("/api")
public class UserController {
    
    @GetMapping("/users")
    public List<User> getUsers() {
        return userService.findAll();
    }
    
    @PostMapping("/users")
    public User createUser(@RequestBody User user) {
        return userService.save(user);
    }
    
    @GetMapping("/users/{id}")
    public ResponseEntity<User> getUser(@PathVariable Long id) {
        User user = userService.findById(id);
        if (user == null) {
            return ResponseEntity.notFound().build();
        }
        return ResponseEntity.ok(user);
    }
}

@SpringBootApplication
public class Application {
    public static void main(String[] args) {
        SpringApplication.run(Application.class, args);
    }
}
```

**Gin应用:**
```go
package main

import (
    "net/http"
    "strconv"
    
    "github.com/gin-gonic/gin"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
    Email string `json:"email"`
}

var users = []User{
    {ID: 1, Name: "张三", Email: "zhangsan@example.com"},
    {ID: 2, Name: "李四", Email: "lisi@example.com"},
}

func main() {
    // 创建Gin路由器
    r := gin.Default()
    
    // API路由组
    api := r.Group("/api")
    {
        api.GET("/users", getUsers)
        api.POST("/users", createUser)
        api.GET("/users/:id", getUser)
    }
    
    // 启动服务器
    r.Run(":8080")
}

func getUsers(c *gin.Context) {
    c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
    var newUser User
    if err := c.ShouldBindJSON(&newUser); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    newUser.ID = len(users) + 1
    users = append(users, newUser)
    c.JSON(http.StatusCreated, newUser)
}

func getUser(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }
    
    for _, user := range users {
        if user.ID == id {
            c.JSON(http.StatusOK, user)
            return
        }
    }
    
    c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
}
```

## 🛣️ 路由系统

### 基础路由
```go
func main() {
    r := gin.Default()
    
    // HTTP方法路由
    r.GET("/get", handleGet)
    r.POST("/post", handlePost)
    r.PUT("/put", handlePut)
    r.DELETE("/delete", handleDelete)
    r.PATCH("/patch", handlePatch)
    r.HEAD("/head", handleHead)
    r.OPTIONS("/options", handleOptions)
    
    // 参数路由
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(200, gin.H{"user_id": id})
    })
    
    // 通配符路由
    r.GET("/files/*filepath", func(c *gin.Context) {
        filepath := c.Param("filepath")
        c.JSON(200, gin.H{"filepath": filepath})
    })
    
    r.Run(":8080")
}
```

### 路由组
```go
func main() {
    r := gin.Default()
    
    // 简单路由组
    v1 := r.Group("/v1")
    {
        v1.GET("/users", getUsersV1)
        v1.POST("/users", createUserV1)
    }
    
    // 带中间件的路由组
    v2 := r.Group("/v2")
    v2.Use(authMiddleware())
    {
        v2.GET("/users", getUsersV2)
        v2.POST("/users", createUserV2)
        v2.PUT("/users/:id", updateUserV2)
        v2.DELETE("/users/:id", deleteUserV2)
    }
    
    r.Run()
}
```

**Spring Boot路由对比:**
```java
// Spring Boot使用注解定义路由
@RestController
@RequestMapping("/api/v1")
public class UserController {
    
    @GetMapping("/users")
    public List<User> getUsers() { ... }
    
    @PostMapping("/users")
    public User createUser(@RequestBody User user) { ... }
}

// 版本控制
@RestController
@RequestMapping("/api/v2")
public class UserV2Controller {
    
    @GetMapping("/users")
    @PreAuthorize("hasRole('USER')")  // 类似Gin中间件
    public List<User> getUsers() { ... }
}
```

## 🔗 中间件系统

### 内置中间件
```go
func main() {
    // gin.Default() 包含Logger和Recovery中间件
    r := gin.Default()
    
    // 或者使用gin.New()创建无中间件的实例
    r = gin.New()
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    
    r.GET("/test", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "test"})
    })
    
    r.Run()
}
```

### 自定义中间件
```go
// 认证中间件
func authMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing token"})
            c.Abort()  // 停止后续处理
            return
        }
        
        // 验证token逻辑
        if !validateToken(token) {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        
        // 设置用户信息到上下文
        c.Set("user_id", getUserIDFromToken(token))
        c.Next()  // 继续处理
    }
}

// 日志中间件
func loggerMiddleware() gin.HandlerFunc {
    return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
        return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
            param.ClientIP,
            param.TimeStamp.Format(time.RFC1123),
            param.Method,
            param.Path,
            param.Request.Proto,
            param.StatusCode,
            param.Latency,
            param.Request.UserAgent(),
            param.ErrorMessage,
        )
    })
}

// CORS中间件
func corsMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Header("Access-Control-Allow-Origin", "*")
        c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
        
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        
        c.Next()
    }
}

func main() {
    r := gin.New()
    
    // 全局中间件
    r.Use(loggerMiddleware())
    r.Use(gin.Recovery())
    r.Use(corsMiddleware())
    
    // 特定路由使用中间件
    protected := r.Group("/admin")
    protected.Use(authMiddleware())
    {
        protected.GET("/users", getAdminUsers)
        protected.DELETE("/users/:id", deleteUser)
    }
    
    r.Run()
}
```

**Spring Boot中间件对比:**
```java
// Spring Boot使用Filter或Interceptor
@Component
public class AuthFilter implements Filter {
    @Override
    public void doFilter(ServletRequest request, ServletResponse response, 
                        FilterChain chain) throws IOException, ServletException {
        HttpServletRequest req = (HttpServletRequest) request;
        String token = req.getHeader("Authorization");
        
        if (token == null || !validateToken(token)) {
            HttpServletResponse res = (HttpServletResponse) response;
            res.setStatus(HttpServletResponse.SC_UNAUTHORIZED);
            return;
        }
        
        chain.doFilter(request, response);
    }
}

// 或使用拦截器
@Component
public class AuthInterceptor implements HandlerInterceptor {
    @Override
    public boolean preHandle(HttpServletRequest request, 
                           HttpServletResponse response, 
                           Object handler) throws Exception {
        // 认证逻辑
        return true;
    }
}
```

## 📨 请求处理

### 参数绑定
```go
type LoginRequest struct {
    Username string `json:"username" binding:"required"`
    Password string `json:"password" binding:"required,min=6"`
}

type QueryParams struct {
    Page  int    `form:"page"`
    Size  int    `form:"size"`
    Query string `form:"q"`
}

func handleLogin(c *gin.Context) {
    var req LoginRequest
    
    // 绑定JSON数据
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 处理登录逻辑
    token, err := authenticate(req.Username, req.Password)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"token": token})
}

func handleSearch(c *gin.Context) {
    var params QueryParams
    
    // 绑定查询参数
    if err := c.ShouldBindQuery(&params); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    // 设置默认值
    if params.Page == 0 {
        params.Page = 1
    }
    if params.Size == 0 {
        params.Size = 10
    }
    
    results := searchData(params.Query, params.Page, params.Size)
    c.JSON(http.StatusOK, gin.H{"data": results})
}

func handleUpload(c *gin.Context) {
    // 单文件上传
    file, err := c.FormFile("file")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
        return
    }
    
    // 保存文件
    dst := fmt.Sprintf("./uploads/%s", file.Filename)
    if err := c.SaveUploadedFile(file, dst); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
        return
    }
    
    c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
```

**Spring Boot参数绑定对比:**
```java
@PostMapping("/login")
public ResponseEntity<?> login(@Valid @RequestBody LoginRequest request) {
    // 自动验证和绑定
    String token = authService.authenticate(request.getUsername(), request.getPassword());
    return ResponseEntity.ok(Map.of("token", token));
}

@GetMapping("/search")
public ResponseEntity<?> search(@RequestParam(defaultValue = "1") int page,
                               @RequestParam(defaultValue = "10") int size,
                               @RequestParam String q) {
    // Spring自动转换参数类型
    List<Result> results = searchService.search(q, page, size);
    return ResponseEntity.ok(Map.of("data", results));
}
```

## 🎨 响应处理

### 不同格式响应
```go
func handleResponse(c *gin.Context) {
    format := c.Query("format")
    
    data := map[string]interface{}{
        "message": "Hello World",
        "time":    time.Now(),
        "data":    []int{1, 2, 3, 4, 5},
    }
    
    switch format {
    case "xml":
        c.XML(http.StatusOK, data)
    case "yaml":
        c.YAML(http.StatusOK, data)
    case "html":
        c.HTML(http.StatusOK, "index.html", data)
    default:
        c.JSON(http.StatusOK, data)
    }
}

func handleFile(c *gin.Context) {
    filename := c.Param("filename")
    filepath := fmt.Sprintf("./files/%s", filename)
    
    // 检查文件是否存在
    if _, err := os.Stat(filepath); os.IsNotExist(err) {
        c.JSON(http.StatusNotFound, gin.H{"error": "File not found"})
        return
    }
    
    // 返回文件
    c.File(filepath)
}

func handleDownload(c *gin.Context) {
    filename := "report.pdf"
    filepath := fmt.Sprintf("./reports/%s", filename)
    
    // 设置下载头
    c.Header("Content-Description", "File Transfer")
    c.Header("Content-Disposition", "attachment; filename="+filename)
    c.File(filepath)
}
```

## 🔧 配置和部署

### 应用配置
```go
type Config struct {
    Server struct {
        Port string `json:"port"`
        Host string `json:"host"`
    } `json:"server"`
    
    Database struct {
        Driver   string `json:"driver"`
        Host     string `json:"host"`
        Port     int    `json:"port"`
        Name     string `json:"name"`
        Username string `json:"username"`
        Password string `json:"password"`
    } `json:"database"`
    
    JWT struct {
        Secret     string `json:"secret"`
        Expiration int    `json:"expiration"`
    } `json:"jwt"`
}

func loadConfig() (*Config, error) {
    var config Config
    
    // 读取配置文件
    file, err := os.Open("config.json")
    if err != nil {
        return nil, err
    }
    defer file.Close()
    
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(&config); err != nil {
        return nil, err
    }
    
    return &config, nil
}

func main() {
    // 加载配置
    config, err := loadConfig()
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }
    
    r := gin.Default()
    
    // 设置信任的代理
    r.SetTrustedProxies([]string{"127.0.0.1"})
    
    // 设置HTML模板
    r.LoadHTMLGlob("templates/*")
    
    // 静态文件服务
    r.Static("/static", "./static")
    
    // 启动服务器
    addr := fmt.Sprintf("%s:%s", config.Server.Host, config.Server.Port)
    log.Printf("Server starting on %s", addr)
    r.Run(addr)
}
```

**Spring Boot配置对比:**
```yaml
# application.yml
server:
  port: 8080
  host: localhost

spring:
  datasource:
    driver-class-name: com.mysql.cj.jdbc.Driver
    url: jdbc:mysql://localhost:3306/mydb
    username: user
    password: password

jwt:
  secret: my-secret-key
  expiration: 86400
```

```java
@Configuration
@ConfigurationProperties(prefix = "jwt")
public class JwtConfig {
    private String secret;
    private int expiration;
    // getters and setters
}
```

## 📝 实践任务

### 任务1: 基础API
1. 创建用户管理API
2. 实现CRUD操作
3. 添加参数验证

### 任务2: 中间件开发
1. 实现JWT认证中间件
2. 添加请求日志中间件
3. 实现限流中间件

### 任务3: 完整应用
1. 开发Blog API系统
2. 集成数据库
3. 添加文件上传功能

## 🎯 学习要点

### Gin框架特点
1. **高性能**: 基于httprouter，性能优异
2. **简洁**: API设计简单易用
3. **中间件**: 强大的中间件系统
4. **灵活**: 可扩展性强

### 与Spring Boot的差异
1. **轻量级**: Gin更轻量，Spring Boot功能更全
2. **配置**: Gin需手动配置，Spring Boot自动配置
3. **生态**: Spring生态更丰富
4. **学习曲线**: Gin更简单，Spring Boot更复杂

## 🎯 下一步
- 学习GORM数据库操作
- 掌握微服务架构
- 理解容器化部署

## 📚 参考资源
- [Gin官方文档](https://gin-gonic.com/)
- [Gin GitHub仓库](https://github.com/gin-gonic/gin)
- [Go Web编程实战](https://github.com/astaxie/build-web-application-with-golang) 