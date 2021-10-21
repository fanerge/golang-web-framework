package main

// 实现路由分组控制(Route Group Control)
// 更好的对路由分组进行权限控制
// 以/admin 开头的路由需要鉴权
// 以/api 开头的路由是 RESTful 接口，可以对接第三方平台，需要三方平台鉴权

// test
// curl "http://localhost:9999/v1/hello?name=geektutu"
// curl "http://localhost:9999/v2/hello/geektutu"

import (
	"gee"
	"net/http"
	"fmt"
)

func main() {
	fmt.Print("手机终端体检代码测试")
	r := gee.New() // engine
	r.GET("/index", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gee.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		})

		v1.GET("/hello", func(c *gee.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *gee.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gee.Context) {
			c.JSON(http.StatusOK, gee.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}
	r.Run(":9999")
}
