package main

import (
	"gin_study/shanshan_Gin/part06/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 写路由
	// 加载html页面：
	r.LoadHTMLGlob("template/**/*")
	r.Static("/s", "static") // 指定js/css文件

	// 定义路由
	r.GET("/userIndex", myfunc.Hello1)
	r.POST("/savefile", myfunc.Hello3)
	r.Run()
}
