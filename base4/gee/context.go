package gee

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type H map[string]interface{}

type Context struct {
	Writer http.ResponseWriter
	Req *http.Request
	Path string
	Method string
	StatusCode int
}
//构建一个上下文实例
func newContext(w http.ResponseWriter,req *http.Request) *Context  {
	return &Context{
	    Writer : w,
		Req: req,
		Path: req.URL.Path,
		Method : req.Method,
	}
}

//获取请求体中的表单数据
func (c *Context) postForm(key string) string{
	return c.Req.FormValue(key)
}
//通过参数名获取请求参数中对应的值
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}
//设置响应状态码
func (c *Context) Status(code int)  {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}
//设置响应头信息
func (c *Context) SetHeader(key string,value string)  {
	c.Writer.Header().Set(key, value)
}
//... 语法
func (c *Context) String(code int,format string,value ...interface{})  {
	c.SetHeader("content-Type","text/plain")
	c.Status(code)
	c.Writer.Write([]byte(fmt.Sprintf(format,value...)))
}

//将需要响应的信息json化之后写入响应writer中
func (c *Context) JSON(code int,obj interface{})  {
	c.SetHeader("Content-Type","application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer,err.Error(),500)
	}
}

func (c *Context) Data(code int,data []byte)  {
	c.Status(code)
	c.Writer.Write(data)
}

func (c *Context) HTML(code int,html string)  {
	c.SetHeader("Content-Type","text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}