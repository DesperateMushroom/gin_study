package main

import (
	"gin_study/shanshan_Gin/part08/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 写路由
	// 定义路由
	r.GET("/redirect1", myfunc.Red1)
	r.GET("/redirect2", myfunc.Red2)

	r.Run()
}
