package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("users", func(ctx *gin.Context) {
		// 上传多个文件
		form, err := ctx.MultipartForm() // 返回一个map,包含多个文件, value就是
		// type Form struct {
		// Value map[string][]string
		// File  map[string][]*FileHeader

		if err != nil {
			fmt.Println(err)
			return
		}

		// fileHeader 包含：filename，header，size，content，tmpfile
		for key, headers := range form.File {
			fmt.Println(key, headers) //file [0x3818e134d440 0x3818e134d4a0 0x3818e134d320] 只有一个key和一个header array
			for _, header := range headers {
				fmt.Println(header.Filename)
				ctx.SaveUploadedFile(header, "./"+header.Filename)
			}
		}

	})

	r.Run(":8080")
}
