package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main4() {
	r := gin.Default()

	r.POST("users", func(ctx *gin.Context) {
		// 创建用户
		fileHeader, err := ctx.FormFile("file") // 返回fileHeader 和 error
		// fileHeader 包含：filename，header，size，content，tmpfile

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(fileHeader.Filename) // 文件名字
		fmt.Println(fileHeader.Size)     // 文件大小，单位是字节

		// ========================
		// 复杂模式
		// file, _ := fileHeader.Open() // 文件对象，io流，可读对象
		// byteData, _ := io.ReadAll(file)

		// err = os.WriteFile("xxx.jpg", byteData, 0666) // 保存到当前目录下xxx.jpg, Windows用0666
		// fmt.Println(err)

		// ========================
		// 简单模式
		err = ctx.SaveUploadedFile(fileHeader, "./aaa.jpg")

	})

	r.Run(":8080")
}
