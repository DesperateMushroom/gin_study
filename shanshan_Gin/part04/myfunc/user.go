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

	fmt.Println("u_name", u_name)
	fmt.Println("u_pwd", u_pwd)

}

// ajax后端的处理
func Hello3(c *gin.Context) {
	//获取post-ajax请求的数据，获取对应的数据:
	uname := c.PostForm("uname")
	fmt.Println(uname)
	// 如果获取的数据和“aoao“一样，那么急着前端响应-用户名录入重复
	if uname == "aoao" {
		fmt.Println(uname == "aoao")

		// 向浏览器/前端返回数据，返回json格式
		// mapdata := map[string]interface{}{
		// 	"msg": "用户名重复！",
		// }
		// c.JSON(200, mapdata)
		c.JSON(200, gin.H{
			"msg": "用户名重复！",
		})
	} else {
		c.JSON(200, gin.H{"msg": ""})
	}
}
