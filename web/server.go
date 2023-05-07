package web

import "net/http"

// 确保一定实现了Server接口
var _ Server = &HTTPServer{}

type HandleFunc func(ctx Context)

type Server interface {
	http.Handler
	// Start 启动服务器
	// addr 是监听地址。如果只指定端口，可以使用 ":8081"
	// 或者 "localhost:8082"
	Start(addr string) error

	// AddRoute 增加路由注册的功能
	AddRoute(method string, path string, handleFunc HandleFunc)
}

type HTTPServer struct {
}

func (h *HTTPServer) AddRoute(method string, path string, handleFunc HandleFunc) {
	// 这里注册到路由树中
	panic("implement me")
}

// ServeHTTP 处理请求的入口
func (h *HTTPServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	ctx := &Context{
		Req:  request,
		Resp: writer,
	}
	h.serve(ctx)
}

func (h *HTTPServer) serve(ctx *Context) {
	// TODO
}

func (h *HTTPServer) Start(addr string) error {
	//TODO implement me
	panic("implement me")
}

func (h *HTTPServer) Post(path string, handler HandleFunc) {
	h.AddRoute(http.MethodPost, path, handler)
}

func (h *HTTPServer) Get(path string, handler HandleFunc) {
	h.AddRoute(http.MethodGet, path, handler)
}
