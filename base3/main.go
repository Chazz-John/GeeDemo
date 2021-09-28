package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r :=gee.New()
	r.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "get根路径请求成功!")
	})
	r.POST("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "post根路径请求成功!")
	})
	r.RUN(":8081")
}
