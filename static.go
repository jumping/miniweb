package miniweb

import (
	"net/http"
)

// 用于处理静态文件的结构
type StaticFile struct {
	
}

// 处理静态文件的路由对应的路由函数
func (s StaticFile) Static(res Resource) {
	http.ServeFile(res.W, res.R, "." + res.R.URL.Path)
}