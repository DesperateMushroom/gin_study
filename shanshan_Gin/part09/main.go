package main

import (
	"gin_study/shanshan_Gin/part09/myfunc"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 注册函数:FuncMap是html/FuncMap
	r.SetFuncMap(template.FuncMap{
		// 键值对的作用，key指定前端调用的名字，value指定后端对应的函数
		"add": myfunc.Add,
	})
	// 写路由
	// 加载html页面：
	r.LoadHTMLGlob("template/**/*")

	// 定义路由
	r.GET("/userIndex", myfunc.Hello)
	r.Run()
}
