package main

import (
	"gin_study/shanshan_Gin/part03/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 写路由
	// 加载html页面：
	r.LoadHTMLGlob("template/**/*")
	r.Static("/s", "static")

	// 定义路由
	r.GET("/userIndex", myfunc.Hello1)
	r.POST("/getUserInfo", myfunc.Hello2)
	r.Run()
}
