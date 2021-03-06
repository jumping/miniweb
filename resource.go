//
// 用于包装http.ResponseWriter, *http.Request
// 并添加了一些自定义的处理过程
//
package miniweb

import (
	"net/http"
	"strings"
)

// 包装http.ResponseWriter, *http.Request，名字太长
type Resource struct {
	// 请求中控制器的名称
	C string
	// 对应控制器中的方法
	M string
	W http.ResponseWriter
	R *http.Request
}

// 创建一个新的资源对象
// 并且解析URL中请求的ControllerName和MeathodName
func NewResource(w http.ResponseWriter, r *http.Request) Resource {
	
	// 定义ControllerName和MethodName
	var cn, mn string
	
	path := strings.Split(r.URL.Path, "/")
	
	// 类似如下的请求格式
	// /
	// /Index
	if len(path) == 2 {
		if path[1] == "" {
			cn = "Index"
			mn = "Index"
		} else {
			// 如下请求
			// /ControllerName/MethodName
			cn = path[1]
			mn = "Index"
		}
	} else {
		if path[1] == "" {
			cn = "Index"
		} else {
			cn = path[1]
		}
		
		if path[2] == "" {
			mn = "Index"
		} else {
			mn = path[2]
		}
	}
	
	return Resource {
		C: cn,
		M: mn,
		W: w,
		R: r,
	}
}

// 解析URL中的参数
func (r Resource) URLValue(name string) string {
	if name == "" {
		return ""
	}
	
	values := r.R.URL.Query()
	if value, ok := values[name]; ok {
		return value[0]
	} else {
		return ""
	}
}