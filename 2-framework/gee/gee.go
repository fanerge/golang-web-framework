package gee

import (
	"fmt"
	"net/http"
)

// 定义路由处理函数
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine 中的 router 为 map 实现路由和处理函数的映射
type Engine struct {
	router map[string]HandlerFunc
}

// 创建 engine 返回指针
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func FF() {
	fmt.Print("ffff")
}

// engine 实现 addRoute 方法
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	// addRoute("GET", "/hello", func) => GET-/hello
	key := method + "-" + pattern
	engine.router[key] = handler
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
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}
