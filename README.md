# golang-web-framework
基于 go lang 搭建 WEB 开发框架，适合 go lang 初学者。

# Web 框架功能
1.  框架雏形
2.  设计请求上下文 Context 并提供 response JSON、HTML、String 等快捷方法，路由 Router 封装
3.  路由功能完善
  基础功能（路由注册、路由处理函数匹配及调用）
  使用 Trie 树实现动态路由(dynamic route)解析（支持两种模式:name和*filepath）
4.  路由分组控制（更好实现分组路由的权限控制）
5.  中间件设计（🧅模型）
6.  支持模版引擎
7.  错误恢复及 log 函数调用栈

# 使用
```
git clone https://github.com/fanerge/golang-web-framework.git
cd golang-web-framework
cd 1-http/2-framework/3-context/...
go run main.go
main.go 文件中有测试场景使用 curl
```

