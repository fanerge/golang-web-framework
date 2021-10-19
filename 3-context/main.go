package main

// 将路由(router)独立出来，方便之后增强
// 设计上下文(Context)，封装 Request 和 Response
// 提供对 JSON、HTML 等返回类型的支持
// 路由 handle 一般都有 context 并且 有一些快捷方式

// test
// curl -i http://localhost:9999/
// curl "http://localhost:9999/hello?name=yzf"
// curl "http://localhost:9999/login" -X POST -d 'username=geektutu&password=1234'
// curl "http://localhost:9999/xxx"

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New() // engine
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
