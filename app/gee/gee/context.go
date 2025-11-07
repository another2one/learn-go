package gee

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// 上下文 每次请求产生 记录请求信息
type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Path   string
	Method string
	// response info
	StatusCode int
	Params     map[string]string
	// middleware
	handlers []HandFunc // 所有Middelware + handle 顺序排列执行
	index    int
	engine   *Engine
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

type H map[string]interface{}

// Next 中间件主要执行逻辑 666
func (c *Context) Next() {
	c.index++
	for c.index < len(c.handlers) {
		c.handlers[c.index](c)
		c.index++
	}
}

func (c *Context) String(code int, format string, a ...interface{}) {
	html := fmt.Sprintf(format, a...)
	c.Data(code, []byte(html))
}

func (c *Context) HTML(code int, name string, data interface{}) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	if err := c.engine.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Fail(500, err.Error())
	}
}

func (c *Context) Fail(code int, html string) {
	fmt.Fprintf(c.Writer, "<h4>%s</h4>", html)
}

func (c *Context) JSON(code int, h H) {
	c.SetHeader("Content-Type", "application/json")
	msg, err := json.Marshal(h)
	if err != nil {
		log.Fatalf("json 返回错误：%s", err)
	}
	c.Data(code, msg)
}

func (c *Context) Ok(h H) {
	c.JSON(http.StatusOK, H{
		"code": 200,
		"msg":  "操作成功",
		"data": h,
	})
}

func (c *Context) Error(code int, msg string) {
	c.JSON(http.StatusOK, H{
		"code": code,
		"msg":  msg,
		"data": map[string]interface{}{},
	})
}

func (c *Context) ErrorWithData(code int, msg string, h H) {
	c.JSON(http.StatusOK, H{
		"code": code,
		"msg":  msg,
		"data": h,
	})
}

func (c *Context) OkWithMsg(h H, msg string) {
	c.JSON(http.StatusOK, H{
		"code": 200,
		"msg":  msg,
		"data": h,
	})
}

func (c *Context) Query(name string) string {
	return c.Req.URL.Query().Get(name)
}

func (c *Context) Param(key string) string {
	value := c.Params[key]
	return value
}

func (c *Context) PostForm(name string) string {
	return c.Req.PostFormValue(name)
}

func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}
