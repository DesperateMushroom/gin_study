package main

// 学习gin的json响应格式

import (
	"gin_study/Response/res"

	"github.com/gin-gonic/gin"
)

func main1() {
	// 1. initilization
	r := gin.Default()
	// 2. route
	r.GET("/index", func(ctx *gin.Context) {
		//ctx.JSON(200, gin.H{"code": 0, "msg": "success", "data": gin.H{}})// 我们想对这句做个封装，使用起来更简便
		res.OKWithMsg(ctx, "登陆成功")

	})

	r.GET("/login", func(ctx *gin.Context) {
		res.OKWithMsg(ctx, "登陆成功")

	})
	r.GET("/users", func(ctx *gin.Context) {
		// 用户的话可以传入一个用户的json data
		res.OKWithData(ctx, map[string]any{
			"name": "alice",
		})
	})
	r.POST("/users", func(ctx *gin.Context) {
		// 创建用户，但是验证失败，用fail
		res.FailWithMsg(ctx, "参数错误")
	})

	// gin.H  会变成json对象
	// gin.H 返回一个 "map[string]any" 的类型

	// 3. listen and serve
	r.Run("127.0.0.1:8080")
}
