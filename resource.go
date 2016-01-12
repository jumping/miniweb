//
// 用于包装http.ResponseWriter, *http.Request
// 并添加了一些自定义的处理过程
//
package miniweb

import (
	"net/http"
)

// 包装http.ResponseWriter, *http.Request，名字太长
type Resource struct {
	W http.ResponseWriter
	R *http.Request
}

func NewResource(w http.ResponseWriter, r *http.Request) Resource {
	return Resource {
		W: w,
		R: r,
	}
}