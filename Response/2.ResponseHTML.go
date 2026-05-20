package main

import "github.com/gin-gonic/gin"

func main2() {
	// 1. initialization
	r := gin.Default()

	// 2. route
	// 告诉文件路径，把文件load出来
	r.LoadHTMLGlob("templates/*")           // 这个是传了个文件夹，里面的文件都可以用
	r.LoadHTMLFiles("templates/index.html") // 直接文件地址，加载单个
	r.GET("/Index", func(ctx *gin.Context) {
		// using HTML response
		// second param is a file name
		// third param，can be used in "index.html" file, not really useful in frontend/backend separated
		ctx.HTML(200, "index.html", map[string]any{
			"title": "learning gin",
		})
	})

	// 3. listen to a port
	r.Run(":8080")
}
