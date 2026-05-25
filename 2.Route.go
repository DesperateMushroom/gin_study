package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	apiGroup := r.Group("api") // 不同功能放在不同组里
	UserGroup(apiGroup)
	r.Run(":8080")
}

func UserView(c *gin.Context) {
	path := c.Request.URL
	fmt.Println(c.Request.Method, path)
}

func UserGroup(r *gin.RouterGroup) {
	r.GET("users", UserView)
	r.POST("users", UserView)
	r.DELETE("users", UserView)
	r.PUT("users", UserView)

}
