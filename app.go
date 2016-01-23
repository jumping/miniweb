package miniweb

import (
	"fmt"
	"net/http"
)

// 全局配置对象
var Conf *Config

// 导入包的时候，默认初始化全局配置对象
// 注意：这里的初始化是在导入该包的时候
// 也就是说这个对象只会在程序启动是初始化一次
// 在程序运行的过程中修改配置文件是不会生效的
func init() {
	Conf = InitConfig()
}

// 运行服务器
func Run(host string, mux *Router) {
	fmt.Println("Start  services at ", host, ".....")
	
	http.ListenAndServe(host, mux)
}
