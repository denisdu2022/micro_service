package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"stream_grpc_test/proto"
	"sync"
	"time"
)

const POST = ":8000"

type server struct {
	proto.UnimplementedGreeterServer
}

//参照之前一元调用的方法
//func (s *server) GetStream(ctx context.Context, req proto.StreamReqData) (*proto.StreamResData, error) {
//	return nil, nil
//}

//服务端流模式

func (s *server) GetStream(req *proto.StreamReqData, res proto.Greeter_GetStreamServer) error {
	i := 0
	for {
		i++
		_ = res.Send(&proto.StreamResData{
			Data: fmt.Sprintf("%v", time.Now()),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	return nil
}

//客户端流模式

func (s *server) PutStream(cliStr proto.Greeter_PutStreamServer) error {

	for {
		if a, err := cliStr.Recv(); err != nil {
			fmt.Println(err)
			break
		} else {
			fmt.Println(a.Data)
		}

	}
	return nil
}

//双向流模式

func (s *server) AllStream(allStr proto.Greeter_AllStreamServer) error {
	//使用协程
	wg := sync.WaitGroup{}
	wg.Add(2)
	//接收客户端消息
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到客户端消息: " + data.Data)
		}
	}()

	//向客户端发送消息
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamResData{
				Data: "我是服务器",
			})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

	return nil
}

func main() {
	//启动服务
	listener, err := net.Listen("tcp", POST)
	//错误处理
	if err != nil {
		panic(err)
	}

	//创建grpc的server
	s := grpc.NewServer()
	//注册grpc的服务
	proto.RegisterGreeterServer(s, &server{})

	//绑定服务
	err = s.Serve(listener)
	if err != nil {
		panic("failed to start grpc:" + err.Error())
	}
}
