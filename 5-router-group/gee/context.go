package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{} //  interface{} 空接口可以放任何东西，有点像 ts 的any

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	Params map[string]string
	// response info
	StatusCode int
}

// 每个请求都需要新建一个 Context
func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// 快捷获取 post 参数
func (c *Context) PostForm(key string) string {
	// c.Req 有 FormValue 方法
	return c.Req.FormValue(key)
}

// 快捷获取 query 参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// 快捷设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// 快捷设置 response header
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// 解析 param
func (c *Context) Param(key string) string {
	value, _ := c.Params[key]
	return value
}

// response string 不定参数，
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
}

// response json  Protobuf
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	// 自动import
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// response html
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	// html 转 字节数组 // 8bit 1byte
	c.Writer.Write([]byte(html))
}

// response 其他数据
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
