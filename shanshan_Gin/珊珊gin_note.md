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
        ```go
        // params: 状态码，渲染文件名, 空接口可以接受任意类型
        c.HTML(200, "demo01/hello01.html", nil)
        ```
    - 在html中define定义目录    
        `{{define "demo01/hello01.html"}}`    
        `{{end}}`

## 7.数据互动-使用静态文件
[1] 指定静态文件的路径
1. 方式1：
    `func (group *RouterGroup) Static(relativePath, root string) IRoutes{}`
    - 第一个参数：相对路径
    - 第二个参数：文件夹名称s
    - 含义：这个相对路径映射到哪个文件夹上去   
        `r.Static("/s", "/static") // 用‘/s'来替代 /static路径`   
2. 方式2：
    `func (group *RouterGroup) StaticFS(relativePath, fs http.FileSystem) IRoutes{}`    
    `r.StaticFS("/s", http.Dir("static"))`    
[2] 在前端页面引入静态文件：
    `<link rel="stylesheet" href="/s/css/mycss.css">`

## 8.项目结构调整
【1】单独将函数部分提取：
创建文件夹 - 创建go文件-将函数放入   

【2】在main.go中调用即可
    `myfunc.go`


## 9.数据交互-后端数据给前端-不同类型渲染入页面

### 渲染字符串类型
【1】将要渲染的字符串通过`c.HTML(code, name, interface)`第三个参数传入   
```Go
name := "hello iam aoao"
c.HTML(200, "demo01/hello01.html", name)
```

【2】在页面上利用上下文来获取：   
PS: .代表的就是上下文中你传入的name
`{{.}}`

### 渲染结构体类型
【1】将要渲染的结构体通过`c.HTML(code, name, interface)`第三个参数传入   
```Go
type Student struct {
	Name string
	Age  int
}

func Hello2(c *gin.Context) {
	// 创建结构体实例
	s := Student{
		Name: "aoao",
		Age:  4,
	}
	c.HTML(200, "demo01/hello01.html", s)
}
```

【2】在页面上利用上下文来获取：   
`{{.Name}}<br>{{.Age}}`   

上面的案例传入的是一个结构体实例，以后可能会传入多个结构体，--> 后续利用map来处理


### 渲染数组类型
【1】将要渲染的结构体通过`c.HTML(code, name, interface)`第三个参数传入   
```Go
func Hello3(c *gin.Context) {
	// 定义一个数组：
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	c.HTML(200, "demo01/hello01.html", arr)
}
```
【2】在页面上利用上下文来获取：   
```html
    {{/*这是第一种方式：第一个.代表的是传入的数组的上下文; 第二个.代表遍历的数组的上下文*/}}
    {{/*range . 里的 .	整个数组 arr；      range 内部的 .	当前元素 item*/}}
    {{range .}}
        {{.}}
    {{end}}

    {{/*这是第二种方式：$i-index，$v-value*/}}
    {{range $i,$v := .}}
        {{$i}}
        {{$v}}
    {{end}}
```  


### 渲染结构体数组类型
【1】将要渲染的结构体通过`c.HTML(code, name, interface)`第三个参数传入   
```Go
func Hello4(c *gin.Context) {
	// 定义一个结构体类型的数组：
	var arr [3]Student
	arr[0] = Student{
		Name: "alice",
		Age:  1,
	}
	arr[1] = Student{
		Name: "bob",
		Age:  2,
	}
	arr[2] = Student{
		Name: "cathy",
		Age:  3,
	}
	c.HTML(200, "demo01/hello01.html", arr)
}
```
【2】在页面上利用上下文来获取：   
```HTML
  {{/*这是第一种方式：第一个.代表的是传入的数组的上下文; 第二个.代表遍历的数组的上下文, 代表每一个结构体实例*/}}
    {{/*range . 里的 .	整个数组 arr；      range 内部的 .	当前元素 item*/}}
    {{range .}}
        {{.Name}}
        {{.Age}}
        <br>
    {{end}}

    {{/*这是第二种方式：$i-index，$v-value,如果只用$v接收也是可以的*/}}
    {{range $i,$v := .}}
        {{$i}}：
        {{$v.Name}}-{{$v.Age}}
        <br>
    {{end}}
```

### 渲染map类型
【1】将要渲染的结构体通过`c.HTML(code, name, interface)`第三个参数传入   
```Go
func Hello5(c *gin.Context) {
	// 定义一个Map
	var a map[string]int = make(map[string]int, 3)
	//将键值对存入map
	a["alice"] = 1
	a["bob"] = 2
	a["cathy"] = 3
	c.HTML(200, "demo01/hello01.html", a)
}
```
【2】在页面上利用上下文来获取：  
```HTML
    {{/*获取map中内容，通过key获取value值， .代表上下文中的map*/}}
    {{.alice}}<br>
    {{.bob}}

    {{查看所有key & value}}
    {{range $key, $value := .}}
        key: {{$key}} value: {{$value}} <br>
    {{end}}
```

### 渲染多个结构体类型
【1】解决：将多个结构体类型存入Map中：   

【2】将要渲染的结构体通过`c.HTML(code, name, interface)`第三个参数传入   
```Go
func Hello6(c *gin.Context) {
	// 定义一个Map
	var a map[string]Student = make(map[string]Student, 3)
	//将键值对存入map
	a["no1"] = Student{
		Name: "alice",
		Age:  1,
	}
	a["no2"] = Student{
		Name: "bob",
		Age:  2,
	}
	a["no3"] = Student{
		Name: "cathy",
		Age:  3,
	}
	c.HTML(200, "demo01/hello01.html", a)
}
```

【3】在页面上利用上下文来获取：  
```HTML

    {{/*.代表上下文Map，通过key得到value*/}}
    {{.no1.Name}} = {{.no1.Age}} <br>
    {{.no2.Name}} = {{.no2.Age}} <br>
    {{.no3.Name}} = {{.no3.Age}} <br>
```

### 渲染切片类型
【1】将要渲染的结构体通过`c.HTML(code, name, interface)`第三个参数传入   
```Go
func Hello7(c *gin.Context) {
	// 创建切片
	slice := []int{1, 2, 3, 4, 5, 6}
	c.HTML(200, "demo01/hello01.html", slice)
}
```
【2】在页面上利用上下文来获取：  
和数组的一模一样

