package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main3() {
	r := gin.Default()

	r.POST("users", func(ctx *gin.Context) {
		// 创建用户
		name := ctx.PostForm("name") // 分不清传还是没传，name是空字符和没传都是 ""
		// 传没传
		age, ok := ctx.GetPostForm("age") // ok==true: 传了
		fmt.Println(name)
		fmt.Println(age, ok)

		// http://127.0.0.1:8080/users
		// form-data: name - aoao; age - 11
		// output: aoao
		// 		   11 true
	})

	r.Run(":8080")
}
