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