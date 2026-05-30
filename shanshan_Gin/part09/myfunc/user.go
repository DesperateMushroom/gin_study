package myfunc

import (
	"time"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Age  int
	Name string
}

func Hello(c *gin.Context) {
	// 定义数据
	age := 19
	arr := []int{33, 66, 99}
	flag := true
	username := "aoao"
	now_time := time.Now()
	// 创建结构体实例
	stu := Student{
		Age:  11,
		Name: "Alice",
	}
	// 将age和arr放入map中
	map_data := map[string]any{
		"age":      age,
		"arr":      arr,
		"stu":      stu,
		"flag":     flag,
		"username": username,
		"now_time": now_time,
	}
	c.HTML(200, "demo01/hello.html", map_data)
}

// 定义一个函数:
func Add(num1 int, num2 int) int {
	return num1 + num2
}
