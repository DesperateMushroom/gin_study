package myfunc

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	name := "hello iam aoao"
	c.HTML(200, "demo01/hello01.html", name)
}
