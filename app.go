package miniweb

import (
	"fmt"
	"net/http"
)

// 全局参数
var (
	// 全局配置对象
	Conf *Config
	// 默认的模板路径
	VIEW string = "./view"
	// 默认模板后缀
	SUFFIX string = ".html"
	// 是否开启模板，默认关闭
	LAYOUT bool = false
	// 默认的layout目录的名称
	LAYOUT_DIR string = VIEW + "/" + "layout"
)

// 导入包的时候，默认初始化全局配置对象
// 注意：这里的初始化是在导入该包的时候
// 也就是说这个对象只会在程序启动是初始化一次
// 在程序运行的过程中修改配置文件是不会生效的
func init() {
	Conf = InitConfig()
	
	// 初始化配置项 section = Global
	var section string = "Global"
	var tmp string =""
	if tmp = Conf.Get(section, "view"); tmp != "" {
		VIEW = tmp
	}
	if tmp = Conf.Get(section, "suffix"); tmp != "" {
		SUFFIX = tmp
	}
	if tmp = Conf.Get(section, "layout"); tmp != "" {
		if tmp == "true" {
			LAYOUT = true
		} else {
			LAYOUT = false
		}
	}
	if tmp = Conf.Get(section, "layout_DIR"); tmp != "" {
		LAYOUT_DIR = VIEW + "/" + tmp
	}
	
}

// 运行服务器
func Run(host string, mux *Router) {
	fmt.Println("Start  services at ", host, ".....")
	
	http.ListenAndServe(host, mux)
}
