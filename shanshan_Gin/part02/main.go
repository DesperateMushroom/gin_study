package main

import (
	"gin_study/shanshan_Gin/part02/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//写路由：
	// 路由规则中要求你传入id的参数，那么就必须你在访问的时候必须传入参数值
	r.GET("/demo1/:id", myfunc.Hello1)

	// 如果利用*占位符，路径是否带参数就不重要了
	r.GET("/demo2/*id", myfunc.Hello2)

	// 如果路径中以键值对形式传入参数的话，在路由规则中就不用做文章了，不用进行任何操作
	r.GET("/demo3", myfunc.Hello3)

	// 如果路径中以键值对形式传入参数的话，在路由规则中就不用做文章了，不用进行任何操作
	r.GET("/demo4", myfunc.Hello4)

	//数组
	r.GET("/demo5", myfunc.Hello5)

	//map
	r.GET("/demo6", myfunc.Hello6)
	r.Run()
}
