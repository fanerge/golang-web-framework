package gee

import (
	"net/http"
)

// 定义路由处理函数
type HandlerFunc func(*Context)

// Engine 中的 router 为 map 实现路由和处理函数的映射
type Engine struct {
	router *router
}

// 创建 engine 返回指针
func New() *Engine {
	return &Engine{router: newRouter()}
}

// engine 实现 addRoute 方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// engine 实现 GET 方法（addRoute GET 方式简写）
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// engine 实现 POST 方法（addRoute POST 方式简写）
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// 启动服务
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

// 实现 ServeHTTP 方法
// type Handler interface {
// 	ServeHTTP(ResponseWriter, *Request)
// }
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}
