package myfunc

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Hello1(c *gin.Context) {
	c.HTML(200, "demo01/hello.html", nil)
}

func Hello2(c *gin.Context) {
	// 获取前端传入的文件
	file, err := c.FormFile("myfile") // 这里的"myfile"是前端上传文件时<input>的name
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Filename)

	// 加入时间戳：
	time_int := time.Now().Unix()               // 获得当前时间int类型 ~1970.01.01.00.00.00 至今GMT
	time_str := strconv.FormatInt(time_int, 10) // 转成字符串,10进制

	// 保存在我的本地
	c.SaveUploadedFile(file, "./uploaded/"+file.Filename+time_str)

	// 响应
	c.String(200, "文件上传成功")

}

func Hello3(c *gin.Context) {

	// 获取前端传入的文件
	file, err := c.FormFile("file") // 这里的"myfile"是前端上传文件时<input>的name
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Filename)

	// 加入时间戳：
	time_int := time.Now().Unix()               // 获得当前时间int类型 ~1970.01.01.00.00.00 至今GMT
	time_str := strconv.FormatInt(time_int, 10) // 转成字符串,10进制

	// 保存在我的本地
	c.SaveUploadedFile(file, "./uploaded/"+file.Filename+time_str)

	// 响应
	c.String(200, "文件上传成功")

}
