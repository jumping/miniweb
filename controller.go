package miniweb

import (
	"html/template"
	"encoding/json"
	"fmt"
)

type Controller struct {
	
}

// 用于解析模板的方法
func (c Controller) Render(res Resource, temp string, data interface{}) {
	t, _ := template.ParseFiles(temp)
	t.Execute(res.W, data)
}

// 向客户端返回JSON编码格式的数据
func (c Controller) RenderJSON(res Resource, data interface{}) {
	res.W.Header().Set("Content-Type", "text/plain; charset=utf-8")
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	} else {
		res.W.WriteHeader(200)
		fmt.Fprintf(res.W, string(jsonData))
	}
}