package main

// 使用 Trie 树实现动态路由(dynamic route)解析。
// 支持两种模式:name和*filepath
// /hello/:name

// test
// curl "http://localhost:9999/hello/yzf"
// curl "http://localhost:9999/assets/css/index.css"

import (
	"gee"
	"net/http"
)

func main() {
	r := gee.New() // engine
	// r.GET("/", func(c *gee.Context) {
	// 	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	// })

	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=yzf
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/user", func(c *gee.Context) {
		// expect /hello?name=yzf
		c.String(http.StatusOK, "user")
	})

	r.GET("/hello/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/hello/:name/:part", func(c *gee.Context) {
		// expect /hello/yzf/sss
		c.String(http.StatusOK, "hello %s, %s %s\n", c.Param("name"), c.Param("part"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")

}
