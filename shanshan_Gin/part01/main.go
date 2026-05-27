package main

import (
	"gin_study/shanshan_Gin/part01/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 加载html文件：
	r.LoadHTMLGlob("template/**/*")

	// 指定静态文件
	r.Static("/s", "static")

	//写路由：
	r.GET("/demo", myfunc.Hello)
	r.Run()
}
