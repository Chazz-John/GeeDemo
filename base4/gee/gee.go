// Package gee web基本架构.实现基本的请求响应
package gee

import (
	"net/http"
)

type HandlerFunc func(*Context)

type Engine struct {
	router *router
}

// 实现httpServe接口判断请求方式和请求的路径是否在router中,是则响应请求
func (engine *Engine) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	c := newContext(w, request)
	engine.router.handle(c)
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

// 添加路由
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
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
