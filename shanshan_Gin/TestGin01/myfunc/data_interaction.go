package myfunc

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	// params: 状态码，渲染文件名, 空接口可以接受任意类型
	c.HTML(200, "demo01/hello01.html", nil)
}
