package miniweb

import (
	"html/template"
	"encoding/json"
	"net/http"
	"strings"
	"bytes"
	"time"
	"fmt"
	"os"
	"io"
)

type Controller struct {
	// 临时关闭模板功能 默认值为false
	// 临时关闭模板的设置只在当前请求中有效
	CloseLayout bool
	// 指明使用的模板
	Layout string
	// 模板数据
	LayoutData map[string]interface{}
	// 存储模板的缓冲区
	buffer *bytes.Buffer
}

// 用于解析模板的方法
func (c Controller) Render(res Resource, data interface{}) {
	//var tempfile string
	// 判断请求的页面是否已经被缓存
	if isExist, filePath := c.HasCache(res.C, res.M); isExist {
		// 缓存存在则直接输出缓存
		// 将打开的文件内容，复制到响应的对象中去
		file, err := os.Open(filePath)
		if err != nil {
			panic("\n\n存在缓存页面，但缓存页面打开失败，请检查！！！\n\n")
		}
		io.Copy(res.W, file)
		file.Close()
		return
	}
	fmt.Println("没找到缓存文件")
	folder := VIEW + "/" +res.C
	file := folder + "/" + strings.ToLower(res.M) + SUFFIX
	// 创建缓冲区
	c.buffer = bytes.NewBuffer(nil)
	
	// 开启了模板并且没有临时关闭
	if LAYOUT && !c.CloseLayout {
		// 如果开启了模板
		// 先解析相应的页面，再将解析的内容写入解析的模板中，最后输出到浏览器中
		t, err := template.ParseFiles(file)
		if err != nil {
			panic("\n\nError: 模板解析失败\n\t" + err.Error() + "\n\n")
		}
		// 将解析的页面写入缓冲区中
		t.Execute(c.buffer, data)
		// 解析layout
		if c.Layout == "" {
			panic("\n\nError: You opened layout, please set layout name\n\n")
		}
		t, err = template.ParseFiles(LAYOUT_DIR + "/" + c.Layout)
		if err != nil {
			panic("\n\nError: 模板解析失败\n\t" + err.Error() + "\n\n")
		}
		if c.LayoutData == nil {
			c.LayoutData = make(map[string]interface{})
		}
		c.LayoutData["LayoutContent"] = template.HTML(c.buffer.Bytes())
		// 转存完缓冲区中的内容时，需要重置缓冲区
		c.buffer.Reset()
		// 将模板解析完成后也写入缓冲区
		t.Execute(c.buffer, c.LayoutData)
	} else {
		// 没有开启模板就直接解析
		t, err := template.ParseFiles(file)
		if err != nil {
			panic("\n\nError: 模板解析失败\n\t" + err.Error() + "\n\n")
		}
		t.Execute(c.buffer, data)
	}
	
	// 这里统一将缓冲区内容写入res.W，同时保存一份在缓存文件中
	// 响应客户端的请求
	// 设置响应头
	res.W.Header().Add("Content-Type", "text/html; charset=utf-8")
	res.W.WriteHeader(200)
	res.W.Write(c.buffer.Bytes())
	// 判断是否需要缓存页面
	// 调试模式下(debug = true)是不需要创建缓存的
	if !DEBUG {
		// 创建缓存页面
		cachefile := c.CreateCacheFile(res.C, res.M)
		cachefile.Write(c.buffer.Bytes())
		cachefile.Close()
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
func (c Controller) HasCache(cn, mn string) (bool, string) {
	// 先判断是否开启了调试模式
	// 如果开启了调试模式则，不缓存页面
	if DEBUG {
		return false, ""
	}
	// 下面是没有开启调试模式，则缓存存在
	// 直接打开缓存文件
	// TODO 该过程有待优化
	cacheFilePath := "cache/" + cn + "/" + strings.ToLower(mn) + SUFFIX
	// 读取文件信息
	fileinfo, err := os.Stat(cacheFilePath)
	if err == nil || os.IsExist(err) {
		// 当页面缓存存在时		
		modtime := fileinfo.ModTime()
		fmt.Println("modtime", modtime)
		currtime := time.Now()
		fmt.Println("currtime", currtime)
		// 获取缓存文件距今创建多久了
		diff := currtime.Sub(modtime)
		fmt.Println("缓存已经创建了", int(diff / time.Second), "秒.....")
		// 当缓的创建时间距今已经超过设定的秒数，则缓存失效
		if int(diff / time.Second) > Conf.Int("cache", "lifetime") {
			// 解析模板时，会重新创建缓存文件
			return false, ""
		}
		// 这里返回缓存有效，并返回缓存页面的路径
		return true, cacheFilePath
		
	}
	
	// 缓存文件不存在
	return false, ""
}

// 创建缓存文件并返回打开的文件指针
func (c Controller) CreateCacheFile(cn, mn string) *os.File {
	// 判断cache目录是否存在
	_, err := os.Stat("cache/")
	if os.IsNotExist(err) {
		os.Mkdir("cache", 0777)
	}
	// 判断cn目录是否存在
	_, err = os.Stat("cache/" + cn)
	if os.IsNotExist(err) {
		os.Mkdir("cache/" + cn, 0777)
	}
	
	// 创建对应的缓存文件
	f, err := os.Create("cache/" + cn + "/" + strings.ToLower(mn) + SUFFIX)
	if err != nil {
		panic("\n\n\nCreate cache file '" + cn + "/" + mn + "' fail\n\n\n")
	}
	
	return f
}

// 实现页面重定向，临时重定向默认状态码为307
func (c Controller) Redirect(res Resource, url string) {
	http.Redirect(res.W, res.R, url, 307)
}

// 清除指定的Cookie
func (c Controller) ClearCookie(name string, res Resource) {
	// 根据name创建一个过期的cookie
	cookie := &http.Cookie {
		Name: name,
		Value: "",
		Expires: time.Now(),
	}
	http.SetCookie(res.W, cookie)
}