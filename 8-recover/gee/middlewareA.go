package gee

import "fmt"

func MiddleWareA() HandlerFunc {
	return func(c *Context) {
		fmt.Println("中间件A 开始执行")
		c.Next()
		fmt.Println("中间件A 执行结束")
	}
}
