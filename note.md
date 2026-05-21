## go 语法
- `:=`	声明并赋值
- `=`	给已经存在的变量赋值


#### 0.0.0.0 所有内网IP
因为你电脑可能有
- 一台电脑其实可以同时有很多“网卡” docker，wsl..
- localhost `127.0.0.1`
- wifi `192.168.1.5`
- ethernet `10.x.x.x`
- 。。。
以上IP都能访问
用`0.0.0.0`：局域网手机也能访问，Docker 容器也可能访问，外部设备也可能访问
本地开发用`127.0.0.1:8080` 比较安全


#### 为什么服务器常用 0.0.0.0
- 需要让外部用户访问

### Go原生http不方便
- 参数解析与验证：content-type 可能是json，formdata，application。。。参数格式错误也要自己返回
- 路由不太明了：一个路由函数把所有请求方式涵盖 /index，前端希望看到这个路由改用哪种方式去请求
- 响应处理比较原始：现在更多的是返回json格式，原生http需要自己弄json


### 初始gin框架
- 安装第三方库
    - `go get github.com/gin-gonic/gin`
1. 初始化
2. 写路由
3. 监听运行


### Gin的本质：
- 在go的原生http上做封装


## 响应 https://www.fengfengzhidao.com/article/tVPJ1o8BvtodovUy9nK2
- code == 200: 状态码，表示ok，链接通畅
- 响应给前端的json
    - “code” ： 错误码，0-没问题
    - “msg” ：


### map[string]any{}
- map - 映射表
- string - key的类型是string
- any - value可以是任意类型
- {} - 创建一个空map
- struct适合固定字段，而map适合动态字段（json，配置，api返回）


### 关于部署
- 前后端单独部署
- 前端打包之后，后端统一部署


## 10. 查询参数和动态参数
### 查询参数
- ?key=xxx&name=xxx 这种就是查询参数
- 查询参数不是GET请求专属的

### 动态参数
- 用户的个人信息页面，他的路径
- `/users?id=123`  查询参数的模式
- `users/123`   动态参数模式

### 表单参数
一般就是专指form表单

### 文件上传 
- 复杂模式
- 简单模式
- 单文件
- 多文件


### 关于接口测试
常用工具
- postman
- apifox
| 注意：接口测试工具能走通的，前端请求不一定可以走通
- get请求带请求体
- ws加请求头

## 13. 原始内容
- 不同请求体对应的原始内容

#### body阅后即焚问题解决
```go
byteData, _ := io.ReadAll(ctx.Request.Body)
ctx.Request.Body = io.NopCloser(bytes.NewReader(byteData))

name := ctx.PostForm("name")
fmt.Println(name)
```

#### form-data
```
----------------------------753936771297636080885206
Content-Disposition: form-data; name="name"

aoao
----------------------------753936771297636080885206--
```

对应的分隔符就是 `Content-Type:[multipart/form-data; boundary=--------------------------129825813012364991523656]`


### x-www-form-urlencoded
url 编码


### json格式
```
{
    "name":"aoao",
    "age":12
}
```


## 14. bind绑定器
使用binding可以很好的完成参数的绑定

### 查询参数
- `http://127.0.0.1:8080/?name=gege&age=12`
- `ShouldBindQuery()` : 把 URL query 参数自动解析并绑定到 Go struct 里。

### 路径参数
- `users/:id/:name`
- `ShouldBindUri()`：自动把“路径参数（URI parameters）”绑定到 struct。


### 表单参数
- `ShouldBind()`：它会根据 请求的 Content-Type 自动选择绑定方式，把参数填充到 struct 里。
| 注意：不能解析x-www-form-urlencoded的格式

### json参数
- `ShouldBindJSON()`

### header参数