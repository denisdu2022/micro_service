package main

import (
	"fmt"
	"io"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type HelloService struct { //类

}

func (s *HelloService) Hello(request string, reply *string) error {
	*reply = "你好: " + request
	return nil
}

func main() {
	//1. 注册处理逻辑  RPC服务
	_ = rpc.RegisterName("HelloService", &HelloService{})

	//2. http URL HandleFunc
	http.HandleFunc("/jsonrpc", func(w http.ResponseWriter, r *http.Request) {
		var conn io.ReadWriteCloser = struct {
			io.Writer
			io.ReadCloser
		}{
			ReadCloser: r.Body,
			Writer:     w,
		}
		rpc.ServeRequest(jsonrpc.NewServerCodec(conn))
	})

	//3.启动服务
	http.ListenAndServe(":8800", nil)
	fmt.Println("Server start!")
}
