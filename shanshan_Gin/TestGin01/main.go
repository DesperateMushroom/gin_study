package main

import (
	"gin_study/shanshan_Gin/TestGin01/myfunc"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 加载文件：
	// r.LoadHTMLFiles("templates/hello01.html", "templates/hello02.html") // 不推荐
	r.LoadHTMLGlob("templates/**/*") //推荐

	// 指定静态文件：指定css文件
	r.Static("/s", "static") // 用‘/s'来替代 /static路径
	// r.StaticFS("/s", http.Dir("static")) // 这俩效果一样，一般用上面的
	r.GET("/demo01", myfunc.Hello5)

	r.Run()
}
