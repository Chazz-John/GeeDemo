package main

import (
	"gee"
)

type i struct {
	username string
	password string
}

func main() {
	r := gee.New()
	r.GET("/", func(context *gee.Context) {
		context.String(200, "请求string接口返回")
	})
	r.POST("/hello", func(context *gee.Context) {
		context.JSON(200, i{
			username: "zhangzhao",
			password: "111111",
		})
	})
	r.RUN(":8081")
}
