package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Context struct {
	W http.ResponseWriter
	R *http.Request
}

func (c *Context) ReadJson(req interface{}) error {
	//帮我读出 body
	//帮我反序列化
	body, err := io.ReadAll(c.R.Body)
	if err != nil {
		fmt.Fprintf(c.W, "read body failed: %v", err)
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	err = json.Unmarshal(body, req)
	if err != nil {
		// 要返回掉，不然就会继续执行后面的代码
		return err
	}
	return nil
}

func (c *Context) WriteJson(code int, resp interface{}) error {
	c.W.WriteHeader(code)
	respJson, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	_, err = c.W.Write(respJson)
	return err
}

func (c *Context) OkJson(resp interface{}) error {
	return c.WriteJson(http.StatusOK, resp)
}

func (c *Context) BadRequestJson(resp interface{}) error {
	return c.WriteJson(http.StatusBadRequest, resp)
}

func NewContext(writer http.ResponseWriter, request *http.Request) *Context {
	return &Context{
		R: request,
		W: writer,
	}
}
