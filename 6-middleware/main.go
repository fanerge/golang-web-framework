package main

// 实现 Web 框架的中间件(Middlewares)机制(洋葱模型)
// 实现通用的Logger中间件，能够记录请求到响应所花费的时间

// test
// curl "http://localhost:9999/v1/hello?name=geektutu"
// curl "http://localhost:9999/v2/hello/geektutu"

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New() // engine
	r.Use(gee.Logger())
	r.Use(gee.MiddleWareA())
	r.Use(gee.MiddleWareB())
	r.GET("/index", func(c *gee.Context) {
		fmt.Printf("路由 handle 执行, %s \n", c.Path)
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
