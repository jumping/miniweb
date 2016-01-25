package miniweb

import (
	"html/template"
	"encoding/json"
	"strings"
	"bytes"
	"fmt"
)

type Controller struct {
	// 指明使用的模板
	Layout string
	// 存储模板的缓冲区
	buffer *bytes.Buffer
}

// 用于解析模板的方法
func (c Controller) Render(res Resource, data interface{}) {
	//var tempfile string
	// 判断请求的页面是否已经被缓存
	if is, _ := c.hasCache(res.C, res.M); is {
		// 缓存存在则直接输出缓存
		// TODO 打开指定的缓存文件，并直接写入到res.W
		
		return
	}
	
	folder := VIEW + "/" +res.C
	file := folder + "/" + strings.ToLower(res.M) + SUFFIX
	
	if LAYOUT {
		buf := make([]byte, 1024)
		c.buffer = bytes.NewBuffer(buf)
		// 如果开启了模板
		// 先解析相应的页面，再将解析的内容写入解析的模板中，最后输出到浏览器中
		t, err := template.ParseFiles(file)
		if err != nil {
			panic("\n\nError: 模板解析失败\n\t" + err.Error() + "\n\n")
		}
		t.Execute(c.buffer, data)
		// 解析layout
		if c.Layout == "" {
			panic("\n\nError: You opened layout, please set layout name\n\n")
		}
		t, err = template.ParseFiles(LAYOUT_DIR + "/" + c.Layout)
		if err != nil {
			panic("\n\nError: 模板解析失败\n\t" + err.Error() + "\n\n")
		}
		layoutdata := make(map[string]template.HTML)
		//fmt.Println(c.buffer.Bytes())
		index := bytes.LastIndexByte(c.buffer.Bytes(), 0)
		layoutdata["LayoutContent"] = template.HTML(c.buffer.Bytes()[index + 1 : c.buffer.Len()])
		// 设置响应头
		res.W.Header().Add("Content-Type", "text/html; charset=utf-8")
		res.W.WriteHeader(200)
		t.Execute(res.W, layoutdata)
	} else {
		// 没有开启模板就直接解析
		t, err := template.ParseFiles(file)
		if err != nil {
			panic("\n\nError: 模板解析失败\n\t" + err.Error() + "\n\n")
		}
		t.Execute(res.W, data)
	}
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

// 判断页面缓存是否存在，如果页面缓存过期则也算不存在
// cn ControllerName
// mn MethodName
func (c Controller) hasCache(cn, mn string) (bool, string) {
	// 先判断缓存目录是否存在
	return false, ""
}