MiniWeb
=====
这是一个简易的golang的web框架，简单易用，方便学习

# 使用说明

## 安装
```
go get https://github.com/buchenglei/miniweb
```
## 基本操作
```go
package main

// 导入包
import (
	web "github.com/buchenglei/miniweb"
)

// 创建控制器
type Index struct {
	web.Controller
}

// 响应首页的处理函数
func (i Index)Index(res web.Resource) {
	data := make(map[string]string)
	data["BlogTitle"] = web.Conf.Get("Blog", "name")
	// 如果使用了Layout，可使用下面的操作临时关闭Layout
	// 该操作仅在当前请求中生效
	// i.CloseLayout = true
	i.Layout = "index.html"
	i.Render(res, data)
}

// 创建一个新的路由器
mymux := web.NewRouter()
mymux.Add("/Index", &Index{})

// 启动项目
web.Run(server, mymux)
```

## 配置项
```go
// 在项目目录下创建config.ini配置文件
// 通过该方法获取配置
web.Conf.Get("section", "key")
```
### 配置文件的基本结构

	[server]
	host = 127.0.0.1
	port = 8080
	[Global]
	layout = true
	debug = true
	[cache]
	lifetime = 100
	

### 默认配置
```go

// 默认的模板路径
VIEW string = "./view"
// 默认模板后缀
SUFFIX string = ".html"
// 是否开启模板，默认关闭
LAYOUT bool = false
// 默认的layout目录的名称
LAYOUT_DIR string = VIEW + "/" + "layout"
// 是否开启调试模式，默认关闭调试模式
DEBUG bool = false

```
这些配置项都可以在config.ini的[Global]里修改

## 依赖
[goconfig](http://github.com/Unknwon/goconfig)

## 实例
一个使用miniweb实现的博客系统:<br>
[GoBlog](https://github.com/buchenglei/goBlog)