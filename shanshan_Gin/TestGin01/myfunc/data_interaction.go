package myfunc

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	name := "hello iam aoao"
	c.HTML(200, "demo01/hello01.html", name)
}

type Student struct {
	Name string
	Age  int
}

func Hello2(c *gin.Context) {
	// 创建结构体实例
	s := Student{
		Name: "aoao",
		Age:  4,
	}
	c.HTML(200, "demo01/hello01.html", s)
}

func Hello3(c *gin.Context) {
	// 定义一个数组：
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30

	c.HTML(200, "demo01/hello01.html", arr)
}

func Hello4(c *gin.Context) {
	// 定义一个结构体类型的数组：
	var arr [3]Student
	arr[0] = Student{
		Name: "alice",
		Age:  1,
	}
	arr[1] = Student{
		Name: "bob",
		Age:  2,
	}
	arr[2] = Student{
		Name: "cathy",
		Age:  3,
	}
	c.HTML(200, "demo01/hello01.html", arr)
}
