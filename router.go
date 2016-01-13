//
// 创建和处理路由信息
//

package miniweb

import (
	"net/http"
	"regexp"
)

// Router 路由器
type Router struct {
	// 处理业务逻辑
	mux  map[string]map[string]RouterFunc
	// 处理不同的状态
	status map[int]RouterFunc
}

// RouterFunc 用于路由响应函数的定义
// 将请求的资源路径和对应的处理函数关联
type RouterFunc func(res Resource)

// NewRouter 创建一个新的路由器
// 初始化一个路由器以及相关默认的行为
func NewRouter() *Router {
	r := new(Router)
	r.mux = make(map[string]map[string]RouterFunc)
	r.status = make(map[int]RouterFunc)
	// 初始化默认状态处理函数
	r.status[400] = status400
	r.status[404] = status404
	r.status[405] = status405
	
	return r
}


// Add 添加路由规则
// 需要指明method = ["GET", "POST", ...]
func (r *Router) Add(method, url string, function RouterFunc) {
	if r.mux[method] == nil {
		r.mux[method] = make(map[string]RouterFunc)
	}
	r.mux[method][url] = function
}

// 匹配路由规则
func (r Router) match(method, url string) (RouterFunc, bool) {
	entry := r.mux[method]
	for x, f := range entry {
		if ok, err := regexp.MatchString("^" + x + "$", url); ok && err == nil {
			return f, true
		}
	}
	
	return nil, false
}

// AddStatus 添加自定义的状态处理函数
func (r *Router) AddStatus(status int, function RouterFunc) {
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
}

// 实现Handler接口
func (r Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 包装http.ResponseWriter, *http.Request
	resource := NewResource(w, req)
	// 判断请求的方法是否存在
	if _, ok := r.mux[resource.R.Method]; ok {
		// 在对应的方法中匹配对应的URL
		if f, ok := r.match(resource.R.Method, resource.R.URL.Path); ok {
			// 将标准库中的响应和请求包装成Resource结构
			// 并交给响应的路由函数去处理
			f(resource)
		} else {
			// 处理资源不存在的状态404 Not Found
			r.HandleStatus(404, resource)
		}
	} else {
		// 当前请求的方法不在允许的方法列表中
		// 需要处理状态405 Method Not Allowed
		r.HandleStatus(405, resource)
	}
}
