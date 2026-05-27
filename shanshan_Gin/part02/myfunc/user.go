package myfunc

import "github.com/gin-gonic/gin"

func Hello1(c *gin.Context) {
	// 获取路径中的参数值
	id := c.Param("id")
	c.String(200, "获取路径上拼接的参数值, %s", id)
}

func Hello2(c *gin.Context) {
	// 获取路径中的参数值
	id := c.Param("id")
	c.String(200, "获取路径上拼接的参数值, %s", id) // 能获取就给你拼接，没获取就算了
}

func Hello3(c *gin.Context) {
	// 获取路径中的参数值：通过key获取对应的value
	id := c.Query("id")
	name := c.Query("name")

	c.String(200, "获取路径上拼接的参数值, %s,%s", id, name) // 能获取就给你拼接，没获取就算了
}

func Hello4(c *gin.Context) {
	// 获取路径中的参数值：通过key获取对应的value
	id := c.DefaultQuery("id", "000")
	name := c.DefaultQuery("name", "no name")

	c.String(200, "获取路径上拼接的参数值, %s,%s", id, name) // 没获取到的用默认值
}

func Hello5(c *gin.Context) {
	// 获取路径中的参数值：通过key获取对应的value的多个参数
	idValues := c.QueryArray("id")

	c.String(200, "获取路径上拼接的参数值, %s", idValues) // 没获取到的用默认值
}

func Hello6(c *gin.Context) {
	// 获取路径中的参数值：通过key获取对应的value的多个参数
	usersMap := c.QueryMap("users")

	c.String(200, "获取路径上拼接的参数值, %s", usersMap) // 没获取到的用默认值
}
