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

## 6.数据互动-使用模板文件渲染
1. `Engine`的`.LoadHTMLFiles`方法：(不推荐，因为要加载的文件太多的话要写很多路径文件名)
    - 加载子指定的模板文件
    - 不定长参数，可以传多个字符串，使用这个方法需要指定所有要使用的html文件路径

2. `Engine`的`.LoadHTMLGlob`方法 : （推荐
    - `func (engine *Engine) loadHTMLGlob(pattern string)`
    - 加载子文件夹下的模板文件
    - 只有一个参数，通配符，如：templates/*, 意思是找当前项目路径下templates文件夹下所有的html文件

3. 渲染HTML模板文件：Context的HTML方法：
    - 参数
        - 状态码：当浏览者访问一个网页时，浏览者的浏览器回想网页所在的服务器发出请求。当浏览器接收并显示网页前，此网页所在的服务器会返回一个包含HTTP状态码的信息头(server header)用以响应浏览器请求
            - http状态码详解：http://tool.oschina.net/commons?type=5
            - 1**: 信息，服务器收到请求，需要请求者继续操作
            - 2**：成功，操作被成功接收并且处理
            - 3**：重定向，需要进一步的操作以完成请求
            - 4**：客户端错误，请求包含语法错误或者无法完成请求
            - 5**：服务器错误，服务器在处理请求的过程中发生了错误
        - 渲染文件名
        - 传入参数：空接口可以接收任意类型

4. 多级目录的模板指定
如果有多级目录，比如templates下有demo01和demo02两个目录，如果要使用里面的html文件，必须得在Load的时候指定多级才可以，比如`r.LoadHTMLGlob("templates/**/*")`   
    - 有几级目录，得在通配符上标明
        `r.LoadHTMLGlob("templates/**/*")`
    - 指定html文件，处理第一级的templates路径不需要指定，后面的路径都要指定 
        `c.HTML(200, "demo01/hello01.html", nil)`
    - 在html中define定义目录 
        `{{define "demo01/hello01.html"}}` 
        `{{end}}`

## 7.数据互动-使用静态文件
[1] 指定静态文件的路径
1. 方式1：
    `func (group *RouterGroup) Static(relativePath, root string) IRoutes{}`
    - 第一个参数：相对路径
    - 第二个参数：文件夹名称
    - 含义：这个相对路径映射到哪个文件夹上去
    `r.Static("/s", "/static") // 用‘/s'来替代 /static路径`
2. 方式2：
    `func (group *RouterGroup) StaticFS(relativePath, fs http.FileSystem) IRoutes{}`
    `r.StaticFS("/s", http.Dir("static"))`
[2] 在前端页面引入静态文件：
    `<link rel="stylesheet" href="/s/css/mycss.css">`

## 8.项目结构调整
【1】单独将函数部分提取：
创建文件夹 - 创建go文件-将函数放入：

【2】在main.go中调用即可
    `myfunc.go`

