package main

import (
	"gee"
)

type i struct {
	usernaem string
	password string
}

func main() {
	r :=gee.New()
	r.GET("/", func(context *gee.Context) {
		context.String(200, "请求string接口返回")
	})
	r.POST("/hello", func(context *gee.Context) {
		context.JSON(200,i{
			usernaem: "zhangzhao",
			password: "111111",
		})
	})
	r.RUN(":8081")
}
