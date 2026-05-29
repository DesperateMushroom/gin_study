
#### 注意：
利用ajax上传文件的话，需要在ajax中添加两个参数：
- `contentType:false`: 默认为true，当设置为true时，jquery ajax提交的时候不会序列号data，而是直接使用data
- `processData:false`: 目的时防止上传文件中出现分界符导致服务器无法正确识别文件起始位置

## 上传单个文件
### 代码：
- frontend:
```HTML
<body>
    <span> this is red </span>
    <br>
    定义一个用户的form表单：
    <form action="/savefile" method="post">
        <input type="file" class="myfile" multiple="multiple">

        <input type="button" value="提交按钮" id="btn">
    </form>
    <script>
        // 获取按钮
        var btn = document.getElementById("btn");
        // 给按钮加入一个单击事件：
        btn.onclick = function(){
            console.log($(".myfile"));  // $(".myfile") 获取到class为myfile的所有文件，形式以数组形式存放
            console.log($(".myfile")[0]); // 获取myfile class文件的第一个
            // 在input type=“file” 里加上 multiple="multiple" ==> 可以一次性上传多个文件
            console.log($(".myfile")[0].files); // 获取到class为myfile的所有文件中的第一个文件中上传的所有文件
            console.log($(".myfile")[0].files[0]);// 获取到class为myfile的所有文件中的第一个文件中上传的所有文件的第一个
            // 创建存放form表单的数据：
            var form_data = new FormData();
            // 在form_data里添加要传入后端的元素：
            form_data.append("file", $(".myfile")[0].files[0])
            // 利用ajax向后端传递数据
            $.ajax({
                url:"/savefile",
                type:"POST",
                data:form_data,
                contentType:false,
                processData:false,
                success : function(info){
                    // 后台响应成功会调用函数
                    // info = 后台响应的数据封装到info里，info名字可以随意
                },
                fail : function(){}
            })
        }
    </script> 
</body>
```

- backend:
```Go

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

```

--- 
## 上传多个文件
### 代码
- frontend：
```HTML
<body>
    <span> this is red </span>
    <br>
    定义一个用户的form表单：
    <form action="/savefile" method="post">
        <input type="file" class="myfile" >
        <input type="file" class="myfile" >
        <input type="file" class="myfile" >

        <input type="button" value="提交按钮" id="btn">
    </form>
    <script>
        // 获取按钮
        var btn = document.getElementById("btn");
        // 给按钮加入一个单击事件：
        btn.onclick = function(){
            // 创建存放form表单的数据：
            var form_data = new FormData();
            // 获取多个文件
            var myfiles = $(".myfile");
            // 对多个文件进行遍历，每一个添加到form_data中去
            for(var i = 0; i<myfiles.length; i++){
                // 在form_data里添加要传入后端的元素：
                form_data.append("file", $(".myfile")[i].files[0])
            }


            
            // 利用ajax向后端传递数据
            $.ajax({
                url:"/savefile",
                type:"POST",
                data:form_data,
                contentType:false,
                processData:false,
                success : function(info){
                    // 后台响应成功会调用函数
                    // info = 后台响应的数据封装到info里，info名字可以随意
                },
                fail : function(){}
            })
        }
    </script> 
</body>
```

- backend:
```Go

func Hello1(c *gin.Context) {
	c.HTML(200, "demo01/hello.html", nil)
}

func Hello(c *gin.Context) {

	// 1. c.MultipartForm() —— Gin 推荐的安全快捷方式
	form, e := c.MultipartForm()
	if e != nil {
		fmt.Println(e)
		return
	}
	// 在form表单中获取name相同的文件
	files := form.File["file"] // File是个map，通过key获取value部分

	// files 就是name相同的多个文件：挨个处理，遍历处理
	for _, file := range files {
		// 获取前端传入的文件
		fmt.Println(file.Filename)

		// 加入时间戳：
		time_int := time.Now().Unix()               // 获得当前时间int类型 ~1970.01.01.00.00.00 至今GMT
		time_str := strconv.FormatInt(time_int, 10) // 转成字符串,10进制

		// 保存在我的本地
		c.SaveUploadedFile(file, "./uploaded/"+time_str+file.Filename)
	}

	// 响应json..
	// ...

}


func main() {
	r := gin.Default()
	// 写路由
	// 加载html页面：
	r.LoadHTMLGlob("template/**/*")
	r.Static("/s", "static") // 指定js/css文件

	// 定义路由
	r.GET("/userIndex", myfunc.Hello1)
	r.POST("/savefile", myfunc.Hello)
	r.Run()
}
```