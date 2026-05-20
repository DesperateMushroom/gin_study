package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main1() {
	r := gin.Default()

	// 请求查询 ?key=xxx&name=zzz&name=yyy
	r.GET("", func(ctx *gin.Context) {
		// 查询用query
		name := ctx.Query("name")
		age := ctx.DefaultQuery("age", "25") //给age设定默认值25
		keyList := ctx.QueryArray("key")
		fmt.Println(name, age, keyList)

		// http://127.0.0.1:8080/?key=wawa&key=meme&name=aoao
		// aoao 25 [wawa meme]
	})

	r.Run(":8080")
}
