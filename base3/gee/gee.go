// Package gee web基本架构.实现基本的请求响应
package gee

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

// 实现httpServe接口判断请求方式和请求的路径是否在router中,是则响应请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	key := request.Method + "-" + request.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, request)
	} else {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "404 NOT FOUNT:%s\n", request.URL)
	}
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

// GET 配置get请求
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST 配置post请求
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// RUN 启动web服务
func (engine *Engine) RUN(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
