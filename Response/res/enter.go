// 让gin http响应更简便

package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

var codeMap = map[int]string{
	1001: "权限错误",
	1002: "角色错误",
}

// 再包一层，好看一点
func response(c *gin.Context, code int, data any, msg string) {
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

//=========================
// 成功响应
// 三种ok响应方式：有data有信息，有data无信息，无data有信息
//=========================
func OK(c *gin.Context, data any, msg string) {
	response(c, 0, data, msg)
}

func OKWithData(c *gin.Context, data any) {
	response(c, 0, data, "success")
}

func OKWithMsg(c *gin.Context, msg string) {
	response(c, 0, gin.H{}, msg)
}

//=========================
// 错误响应
// 三种fail响应方式：有data有信息，有data无信息，无data有信息
// code 根据业务需求指定
//=========================
func Fail(c *gin.Context, code int, data any, msg string) {
	response(c, code, data, msg)
}

func FailWithMsg(c *gin.Context, msg string) {
	response(c, 1001, nil, msg)
}

func FailWithCode(c *gin.Context, code int) {
	// 这里传入的msg一般跟着code走
	// 根据上面的map找msg
	msg, ok := codeMap[code]
	if !ok {
		// 如果没有找到code对应的msg，默认是服务错误
		msg = "服务错误"
	}

	response(c, code, nil, msg)
}
