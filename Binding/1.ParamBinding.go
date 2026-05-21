package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 查询参数
	r.GET("", func(ctx *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := ctx.ShouldBindQuery(&user)
		fmt.Println(user, err)

		// http://127.0.0.1:8080/?name=gege&age=12
		// {gege 12} <nil>

		// http://127.0.0.1:8080/?name=gege&age=d2
		// {gege 0} strconv.ParseInt: parsing "d2": invalid syntax
	})

	//=================================================

	// 路径参数
	r.GET("users/:id/:name", func(ctx *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			ID   int    `uri:"id"`
		}

		var user User
		err := ctx.ShouldBindUri(&user)
		fmt.Println(user, err)

		// http://127.0.0.1:8080/users/111/嗷嗷
		// {嗷嗷 111} <nil>
	})

	//=================================================

	// 表单参数
	r.POST("form", func(ctx *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := ctx.ShouldBind(&user)
		fmt.Println(user, err)

		//{aoao 12} <nil>
		//{aoao 0} strconv.ParseInt: parsing "ff": invalid syntax
	})

	//=================================================

	// JSON参数
	r.POST("json", func(ctx *gin.Context) {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := ctx.ShouldBindJSON(&user)
		fmt.Println(user, err)

		/*
			{
				"name":"aoao",
				"age":"ee"
			}
			{aoao 0} json: cannot unmarshal string into Go struct field User.age of type int
		*/

		/*
			{
				"name":"aoao",
				"age":12
			}
			{aoao 12} <nil>
		*/
	})

	r.POST("header", func(ctx *gin.Context) {
		type User struct {
			Name        string `header:"Name"`
			Age         int    `header:"Age"`
			UserAgent   string `header:"User-Agent"`
			ContentType string `header:"Content-Type"`
		}
		var user User
		err := ctx.ShouldBindHeader(&user)
		fmt.Println(user, err)

	})
	r.Run(":8080")
}
