package miniweb

import (
	"html/template"
)

type Controller struct {
	
}

// 用于解析模板的方法
func (c Controller) Render(res Resource, temp string, data interface{}) {
	t, _ := template.ParseFiles(temp)
	t.Execute(res.W, data)
}