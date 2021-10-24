package main

// 实现错误处理机制

// test
// http://localhost:9999/
// http://localhost:9999/students

import (
	"fmt"
	"gee"
	"net/http"
)

type student struct {
	Name string
	Age  int8
}

func main() {
	r := gee.New() // engine
	// 将本地文件映射到 assets 路由即可
	r.Static("/assets", "/Users/yuzhenfan/Desktop/coding/golang/src/golang-web-framework/static")
	// html template 地址
	r.LoadHTMLGlob("/Users/yuzhenfan/Desktop/coding/golang/src/golang-web-framework/templates/*")

	// 测试异常
	r.GET("/panic", func(c *gee.Context) {
		// Slice
		names := []string{"yzf", "sdf"}
		c.String(http.StatusOK, names[100])
	})

	r.Use(gee.Logger())
	r.Use(gee.MiddleWareA())
	r.Use(gee.MiddleWareB())
	// 守护进程
	r.Use(gee.Recovery())
	stu1 := &student{
		Name: "yzf",
		Age:  20,
	}
	stu2 := &student{
		Name: "yy",
		Age:  02,
	}
	r.GET("/", func(c *gee.Context) {
		fmt.Printf("路由 handle 执行, %s \n", c.Path)
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gee.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gee.H{
			"title":  "age",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
	// r.GET("/index", func(c *gee.Context) {
	// 	fmt.Printf("路由 handle 执行, %s \n", c.Path)
	// 	c.HTML(http.StatusOK, "<h1>Hello index Page</h1>")
	// })

	v1 := r.Group("/v1")
	{
		// v1.GET("/", func(c *gee.Context) {
		// 	c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
		// })

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
