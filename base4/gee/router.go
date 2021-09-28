package gee

import (
	"log"
	"net/http"
)

//定义router
type router struct {
	handlers map[string]HandlerFunc
}

//router构造函数
func newRouter() *router {
	return &router{
		handlers: make(map[string]HandlerFunc),
	}
}

//将请求的类型和请求路由名通过-拼接在一起,将其对应的路由处理器存入map中
func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}