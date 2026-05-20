package main

import "github.com/gin-gonic/gin"

// 静态文件都会放在一个目录（static，upload）里
// 浏览器可以直接访问到
func main() {
	r := gin.Default()

	// 第一个参数是别名，第二个才是实际路径
	r.Static("st", "static") // 访问st的路径，全部都会到static文件里去
	// ex: 127.0.0.1/st/abc.txt

	r.StaticFile("abc", "static/abc.txt")
	r.Run()
}

// 静态文件的路径，不能再被路由使用了 r.GET("abc") 不行！
