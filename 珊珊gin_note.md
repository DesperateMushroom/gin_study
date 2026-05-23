https://www.bilibili.com/video/BV1CV411P78J?spm_id_from=333.788.videopod.episodes&vd_source=95e010711fed90c1308d22313b50b336&p=3

## 第一个gin项目
1. 创建项目目录地址
2. 设置proxy：GoLang在GoModule里设置，vscode在cmd设置go env GOPROXY
3. 下载Gin框架：`go get -u github.com/gin-gonic/gin`
4. 配置file watcher （golang）自动导入啥的
5. 写代码，默认三步
```go
// 1. initialization
r := gin.Default() // get engine

// 2. register routes 挂载路由
r.GET("/index", Index) // same as http.HandleFunc

// 3. bind to a port, run
r.Run("127.0.0.1:8080") // same as http.ListenAndServer
```
6. 运行程序
7. 浏览器访问


## 4. 运行原理
### r := gin.Default()
- `Default()` 返回的是一个引擎，它是框架非常重要的数据结构，是框架的入口
- 引擎 - 框架核心发送机 - 默认服务器 - 整个web服务都是由它来驱动的
- Default()底层调用了New(), 相当于New()升级，New()返回的是一个引擎，在此基础上多增加了中间件处理-engine.Use(Logger(), Recovery())
- 可以用 gin.New() 替代


### r.GET("/", func(context *gin.Context){})
- 路由：通过访问“/”该路径的GET请求，走着一条处理逻辑，走对应函数中的内容
- “/”：路由规则     函数：路由函数
- 路由请求：GET。POST。DELETE。PATCH。PUT。HEAD。OPTIONS。ANY
- 函数：可以直接写匿名函数，还可以在外部定义函数使用
- `gin.Context`: 把请求和响应都封装到`gin.Context`上下文环境中
- func里：对页面的渲染效果（多种），你要给浏览器响应什么效果
- 函数：路径按照自己的项目规则去定义 "/xxx" "/yyy”

### r.Run()
- 启动引擎，服务器启动
- Run可以传入参数：host+port
- 中间拼接的冒号不要忘记


