package main

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func Index(ctx *gin.Context) {
	// reponse a JSON object
	ctx.JSON(200, Response{
		Code: 0,
		Msg:  "success",
		Data: map[string]any{}, //empty object for now
	})
}

func main() {
	// Gin的本质：在go的原生http上做封装

	gin.SetMode("release") // 取消debug日志信息

	// 1. initialization
	r := gin.Default() // get engine

	// 2. register routes 挂载路由
	r.GET("/index", Index) // same as http.HandleFunc

	// 3. bind to a port, run
	r.Run("127.0.0.1:8080") // same as http.ListenAndServer
}
