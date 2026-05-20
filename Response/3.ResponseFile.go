package main

import "github.com/gin-gonic/gin"

//这可以学习 但实际架构使用oss或者本地部署oss直接返回下载地址最好
func main3() {

	// 1. 要设置Content-type，唤起浏览器下载
	// 2. 只能是GET请求

	// 1. initialization
	r := gin.Default()

	// 2. route
	r.GET("/Index", func(ctx *gin.Context) {
		ctx.Header("Content-Type", "application/octet-stream")              // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
		ctx.Header("Content-Disposition", "attachment; filename=3.文件下载.go") // 用来指定下载下来的文件名
		ctx.File("3.ResponseFile.go")                                       // 如果文件不存在，会404报错
	})

	// 3. listen to a port
	r.Run(":8080")
}

// 实际开发中用的更多是：前端请求后端接口，然后唤起浏览器下载
// <a href="文件地址" download="文件名">文件下载</a>

// 最好的做法
// 调下载接口的请求，后端不返回实际文件内容，而是生成一个临时下载地址
// 前端构造a标签，再请求这个接口唤起浏览器下载
