package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// 加载文件：
	// r.LoadHTMLFiles("templates/hello01.html", "templates/hello02.html") // 不推荐
	r.LoadHTMLGlob("templates/**/*") //推荐
	r.GET("/demo01", Hello)

	r.Run()
}

func Hello(c *gin.Context) {
	// params: 状态码，渲染文件名, 空接口可以接受任意类型
	c.HTML(200, "demo01/hello01.html", nil)
}
