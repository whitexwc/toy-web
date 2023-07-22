package main

import (
	"fmt"
	"net/http"
)

type Server interface {
	// Route 设定一个路由，命中该路由的会执行handlerFunc的代码
	Route(method string, pattern string, handleFunc func(ctx *Context))

	// Start 启动本服务器
	Start(address string) error
}

// sdkHttpServer 基于http库实现
type sdkHttpServer struct {
	Name    string
	handler *HandlerBaseOnMap
}

// Route 注册路由
func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *Context)) {
	// http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	// 	ctx := NewContext(writer, request)
	// 	handleFunc(ctx)
	// })
	key := s.handler.key(method, pattern)
	s.handler.handlers[key] = handleFunc
}

func (s *sdkHttpServer) Start(address string) error {
	//http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name:    name,
		handler: &HandlerBaseOnMap{make(map[string]func(ctx *Context), 0)},
	}
}

// 在没有 context 抽象的情况下，是长这样的
func SignUp(ctx *Context) {
	req := &signUpReq{}

	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	//如果要返回个json串回去，就比较麻烦
	resp := &commonResponse{
		Data: 123,
	}

	err = ctx.WriteJson(http.StatusOK, resp)
	//ctx.W.Write(err.Error())
	if err != nil {
		fmt.Printf("写入响应失败：%v", err)
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}

type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}
