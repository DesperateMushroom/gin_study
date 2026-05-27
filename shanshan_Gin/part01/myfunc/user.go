package myfunc

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.HTML(200, "demo01/hello.html", nil)
}
