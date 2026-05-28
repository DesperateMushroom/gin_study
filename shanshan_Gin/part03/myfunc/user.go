package myfunc

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Hello1(c *gin.Context) {
	c.HTML(200, "demo01/hello.html", nil)
}

func Hello2(c *gin.Context) {
	// 获取post请求的参数：
	// PostForm方法：作用：通过key得到value数据
	u_name := c.PostForm("username")

	u_pwd := c.PostForm("pwd")

	// DefaultPostForm:作用：当页面中未定义表单元素进行提交给出默认值
	//						如果页面定义了元素但是没有提交数据。那么不会有默认值，会认为是没有提交数据
	u_age := c.DefaultPostForm("age", "18")

	// PostFormArray：如果前端value数据过多可以用数组接收
	lovelang := c.PostFormArray("lovelang")

	// PostFormMap：获取Map数据,参数需要注意：传入的是整个map(而不是具体的key)
	usermap := c.PostFormMap("user")
	fmt.Println("u_name", u_name)
	fmt.Println("u_pwd", u_pwd)
	fmt.Println("u_age", u_age)
	fmt.Println("lovelang", lovelang)
	fmt.Println("usermap", usermap)

}
