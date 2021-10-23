package gee

import "fmt"

func MiddleWareB() HandlerFunc {
	return func(c *Context) {
		fmt.Println("中间件B 开始执行")
		c.Next()
		fmt.Println("中间件B 执行结束")
	}
}
