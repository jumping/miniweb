//
// 创建和处理路由信息
// 可以处理的URL格式：
// /ControllerName/MethodName/params
//
package miniweb

import (
	"net/http"
	//"regexp"
	"reflect"
	"fmt"
)

// Router 路由器
type Router struct {
	// 处理业务逻辑
	mux  map[string]RouterFunc
	// 处理不同的状态
	status interface{}
}

// RouterFunc 用于路由响应函数的定义
// 将请求的资源路径和对应的处理函数关联
type RouterFunc ControllerInter

// NewRouter 创建一个新的路由器
// 初始化一个路由器以及相关默认的行为
func NewRouter() *Router {
	r := new(Router)
	r.mux = make(map[string]RouterFunc)
	r.status = Status{}
	
	return r
}


// Add 添加路由规则
// 根据请求自动调用指定controller中对应的方法
// 路由格式：
// /ContollerName
// 由框架自动调用对应Controller中的包含的请求方法
func (r *Router) Add(url string, controller RouterFunc) {
	r.mux[url] = controller
}

// 匹配路由规则
/*func (r Router) match(method, url string) (RouterFunc, bool) {
	entry := r.mux[method]
	for x, f := range entry {
		if ok, err := regexp.MatchString("^" + x + "$", url); ok && err == nil {
			return f, true
		}
	}
	
	return nil, false
}*/

// AddStatus 添加自定义的状态处理函数
/*func (r *Router) AddStatus(status int, function RouterFunc) {
	r.status[status] = function
}
// HandleStatus 处理不同的状态信息
func (r Router) HandleStatus(status int, res Resource) {
	if f, ok := r.status[status]; ok {
		res.W.WriteHeader(status)
		f(res)
	} else {
		res.W.WriteHeader(400)
		r.status[400](res)
	}
}*/

// 实现Handler接口
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 包装http.ResponseWriter, *http.Request
	resource := NewResource(w, req)
	c := r.mux["/" + resource.C]
	if c == nil {
		fmt.Println("No path router:", resource.C)
	} else {
		// 反射处理
		// 用于处理URL中的ControllerName和MethodName
		controller := reflect.ValueOf(c).Elem()
		params := make([]reflect.Value,1)
		params[0] = reflect.ValueOf(resource)
		method := controller.MethodByName(resource.M)
		if method.IsValid() {
			method.Call(params)
		} else {
			fmt.Println("Invaild request method:", resource.M)
		}
	}
}
