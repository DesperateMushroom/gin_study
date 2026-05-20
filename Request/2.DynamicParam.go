package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("users/:id", func(ctx *gin.Context) {
		userID := ctx.Param("id") //这个参数要和路径上:之后的对应
		fmt.Println(userID)

		// http://127.0.0.1:8080/users/:1111
		// :1111
		// http://127.0.0.1:8080/users/1111
		// 1111
	})

	r.GET("users/:id/:name", func(ctx *gin.Context) {
		userID := ctx.Param("id") //这个参数要和路径上:之后的对应
		userName := ctx.Param("name")
		fmt.Println(userID, userName)

		// http://127.0.0.1:8080/users/1111/aoao
		// 1111 aoao
	})

	r.Run(":8080")
}
