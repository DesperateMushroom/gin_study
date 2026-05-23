package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// binding 参数校验
func main() {
	r := gin.Default()

	//=================================================

	// 表单参数
	r.POST("form", func(ctx *gin.Context) {
		type User struct {
			Name string `form:"name" binding:"required,min=3,max=5"` // 为空校验
			Age  int    `form:"age"`
		}
		var user User
		err := ctx.ShouldBind(&user)
		fmt.Println(user, err)
		if err != nil {
			ctx.String(200, err.Error())
			return
		}
		ctx.JSON(200, user)
		// 为空校验
		//错误：{ 11} Key: 'User.Name' Error:Field validation for 'Name' failed on the 'required' tag
		//成功：{aoao 11} <nil>
	})

	//=================================================

	// JSON参数
	r.POST("json", func(ctx *gin.Context) {
		type User struct {
			Pwd   string   `json:"Pwd"`
			RePwd string   `json:"RePwd" binding:"required,eqfield=Pwd"`
			IP    []string `json:"IP" binding:"min=1,dive,required,ip"`
		}
		var user User
		err := ctx.ShouldBindJSON(&user)
		fmt.Println(user, err)
		if err != nil {
			ctx.String(200, err.Error())
			return
		}
		ctx.JSON(200, user)
	})

	r.Run(":8080")
}
