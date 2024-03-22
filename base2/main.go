package main

import (
	"fmt"
	"log"
	"net/http"
)

// Engine 定义一个处理request的结构体
type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Headler[%q]= %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND:%s\n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	log.Fatalln(http.ListenAndServe(":8081", engine))
}
