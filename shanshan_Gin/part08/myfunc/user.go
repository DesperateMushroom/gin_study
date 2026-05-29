package myfunc

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Red1(c *gin.Context) {
	fmt.Println("this is redirect1")
	// 发送一个重定向的请求
	c.Redirect(http.StatusFound, "/redirect2") // 重定向的状态码：3xx
}

func Red2(c *gin.Context) {
	fmt.Println("this is redirect2")
	// 在浏览器响应一个字符串
	c.String(http.StatusOK, "成功重定向：这里是red2")
}
