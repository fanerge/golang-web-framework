package main

// redis

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// 创建redis连接池
var pool *redis.Pool

func main() {
	// 链接，默认端口 6379
	// c, err := redis.Dial("tcp", "localhost:6379")
	// if err != nil {
	// 	fmt.Println("err", err)
	// 	return
	// }
	// fmt.Println("redis conn success")
	// defer c.Close()

	// String类型Set、Get操作
	// _, err = c.Do("set", "abc", 101)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// r, err := redis.Int(c.Do("get", "abc"))
	// if err != nil {
	// 	fmt.Println("get abc failed,", err)
	// 	return
	// }
	// fmt.Println(r)

	// String批量操作
	// _, err = c.Do("mset", "abc", 200, "efg", 400)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// r, err := redis.Ints(c.Do("mget", "abc", "efg"))
	// if err != nil {
	// 	fmt.Println("get abc failed,", err)
	// 	return
	// }

	// for _, v := range r {
	// 	fmt.Println(v)
	// }

	// 设置过期时间
	// c.Do("set", "abc", 101)
	// _, err = c.Do("expire", "abc", 10)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(c.Do("get", "abc"))
	// time.Sleep(10 * time.Second)
	// fmt.Println(c.Do("get", "abc"))

	// List队列操作
	// _, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// r, err := redis.String(c.Do("lpop", "book_list"))
	// if err != nil {
	// 	fmt.Println("get abc failed,", err)
	// 	return
	// }
	// fmt.Println(r)

	// Hash表
	// _, err = c.Do("hset", "books", "abc", 100, "sss", 200)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// r, err := redis.Int(c.Do("hget", "books", "abc"))
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(r)

	// Redis连接池
	// 从连接池，取一个链接
	c := pool.Get()
	defer c.Close() // 函数运行结束 ，把连接放回连接池
	_, err := c.Do("set", "abc", 200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("get", "abc"))
	if err != nil {
		fmt.Println("get abc faild :", err)
		return
	}
	fmt.Println(r)
	//关闭连接池
	pool.Close()

}

func init() {
	pool = &redis.Pool{ //实例化一个连接池
		MaxIdle: 16, //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:   0,   //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) { //要连接的redis数据库
			return redis.Dial("tcp", "localhost:6379")
		},
	}
}
