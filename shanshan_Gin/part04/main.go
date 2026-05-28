package main

import (
	"gin_study/shanshan_Gin/part04/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 写路由
	// 加载html页面：
	r.LoadHTMLGlob("template/**/*")
	// 指定js文件 和 css文件
	r.Static("/s", "static")

	// 定义路由
	r.GET("/userIndex", myfunc.Hello1)
	r.POST("/getUserInfo", myfunc.Hello2)
	r.POST("/ajaxpost", myfunc.Hello3)
	r.Run()
}
