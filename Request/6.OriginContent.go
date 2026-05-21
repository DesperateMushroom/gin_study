package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("", func(ctx *gin.Context) {
		// context 封装了http的request和writerRepsonse
		byteData, _ := io.ReadAll(ctx.Request.Body) // 请求体 body是stream流，不是普通字符串
		fmt.Println(string(byteData))               // string() 把byte切片转换成字符串

		// 注意：读了之后，body就没了，阅后即焚
		// 想要他保留body部分，要获取一个ReadCloser的对象
		//把已经读取出来的 byteData，重新包装成一个新的 HTTP Body 流，再塞回 Request.Body
		ctx.Request.Body = io.NopCloser(bytes.NewReader(byteData))

		// 获取分隔符
		fmt.Println(ctx.Request.Header)
		// map[Accept:[*/*]
		// Accept-Encoding:[gzip, deflate, br]
		// Connection:[keep-alive]
		// Content-Length:[267]
		// Content-Type:[multipart/form-data; boundary=--------------------------129825813012364991523656]
		// Postman-Token:[ec5ec080-821d-487d-9091-b5c1c126716b]
		// User-Agent:[PostmanRuntime/7.54.0]]
		// 如果我们要切割，我们要按照Content-Type - boundary这个切割

		name := ctx.PostForm("name")
		fmt.Println(name)

	})

	r.Run(":8080")
}
