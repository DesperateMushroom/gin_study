package main

import (
	"fmt"
	"io"
	"net/http"
)

// go的旧http服务不区分请求方式：GET，POST, PUT, DELETE 都会走到Index
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, r.URL.String()) // 查看请求的方式    GET /index?key=123

	if r.Method != "GET" {
		byteData, _ := io.ReadAll(r.Body) //读取请求体
		fmt.Println(string(byteData))
	}
	fmt.Println(r.Header) //请求头     map[Connection:[Keep-Alive] User-Agent:[Mozilla/5.0 (Windows NT; Windows NT 10.0; en-US) WindowsPowerShell/5.1.26100.8115]]

	// 请求 响应都是字符串
	w.Write([]byte("Hello World")) //响应返回
}

func main() {
	// 编写用户访问某个地址触发的函数
	// 例：http://127.0.0.1:8080/index  .../home  所触发的逻辑
	http.HandleFunc("/index", Index) //当访问到xxx/index, 就会触发第二个参数

	fmt.Println("http server running 127.0.0.1:8080")

	// 在主机的所有ip地址下的8080端口打开一个服务
	// 0.0.0.0:8080 和 :8080是一样的
	http.ListenAndServe(":8080", nil) // 监听8080端口并且提供服务
}
