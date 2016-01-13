MiniWeb
=====
这是一个简易的golang的web框架，简单易用，方便学习

# 使用说明

```go
package main

import (
	C "goBlog/controller"

	web "github.com/buchenglei/miniweb"
)

func main() {
	// 创建一个新的路由器
	mymux := web.NewRouter()
	
	// 处理静态文件
	static := web.StaticFile{}
	mymux.Add("GET", "/static/.*", static.Do)
	
	// 注册路由
	mymux.Add("GET", "/", C.Index)
	mymux.Add("GET", "/admin", C.Admin)
	mymux.Add("GET", "/login", C.Login)
	
	// 启动服务器
	web.Run(":8080", mymux)
}
```
这里是一个实例:<br>
[GoBlog](https://github.com/buchenglei/goBlog)