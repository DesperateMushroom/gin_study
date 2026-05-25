package main

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var trans ut.Translator

// init() 是什么？
// 程序启动自动执行, 不需要调用, 比 main() 更早执行
// 常用于：初始化配置，注册 validator，初始化数据库，初始化日志
func init() {
	// 创建翻译器
	uni := ut.New(zh.New())            // 这里的zh来自：github.com/go-playground/locales/zh
	trans, _ = uni.GetTranslator("zh") //从翻译器里获取中文 translator

	// 注册翻译器，第二部分：拿到 Gin 的 validator 引擎
	v, ok := binding.Validator.Engine().(*validator.Validate) // binding.Validator.Engine()：Gin 内部有默认 validator。需要类型断言
	if ok {
		_ = zh_translations.RegisterDefaultTranslations(v, trans) // 注册中文翻译
		// 注册：required -> 必填字段，min -> 最小值，max -> 最大值，email -> 邮箱格式错误
	}

	// 第四部分：自定义字段名显示（最重要）
	// field： Name string `json:"name" label:"用户名"` - field就是这个Name字段的反射信息
	v.RegisterTagNameFunc(func(field reflect.StructField) string {
		label := field.Tag.Get("label")
		if label == "" {
			label = field.Name
		}
		name := field.Tag.Get("json")
		return fmt.Sprintf("%s---%s", name, label)
	})

	// 	validator
	// 	↓
	// 字段校验失败
	// 	↓
	// 使用 translator 翻译
	// 	↓
	// 使用 RegisterTagNameFunc 自定义字段名
	// 	↓
	// 生成中文错误

	v.RegisterValidation("fip", func(fl validator.FieldLevel) bool {
		fmt.Println("fl.Field():", fl.Field())
		fmt.Println("fl.FieldName():", fl.FieldName())
		fmt.Println("fl.StructFieldName():", fl.StructFieldName())
		fmt.Println("fl.Parent():", fl.Parent())
		fmt.Println("fl.Top():", fl.Top())
		fmt.Println("fl.Param():", fl.Param())

		ip, ok := fl.Field().Interface().(string)
		if ok && ip != "" {
			// 传了值就去校验是不是ip地址
			ipObj := net.ParseIP(ip)
			return ipObj != nil
		}

		return true // 没传入值就不去校验
	})

}

/*
{
  "name": "name参数必填",
}
*/

func ValidateErr(err error) any {
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return err.Error()
	}
	var m = map[string]any{}
	for _, e := range errs {
		msg := e.Translate(trans)
		_list := strings.Split(msg, "---")
		// 如果err里是fip（自定义的校验）
		if e.Tag() == "fip" {
			m[strings.Split(e.Field(), "---")[0]] = "该ip地址不符合要求" //因为自定义的和之前的msg不同，所以要额外处理
			continue
		}
		m[_list[0]] = _list[1]
	}
	return m
}

type User struct {
	IP string `json:"ip" binding:"fip" label:"ip_addr"`
}

func main() {
	r := gin.Default()
	// 注册路由
	r.POST("/user", func(c *gin.Context) {
		var user User
		if err := c.ShouldBindJSON(&user); err != nil {
			// 参数验证失败
			c.JSON(200, map[string]any{
				"code": 7,
				"msg":  "验证错误",
				"data": ValidateErr(err),
			})
			return
		}

		// 参数验证成功
		c.JSON(http.StatusOK, user)
	})

	// 启动HTTP服务器
	r.Run(":8080")
}
