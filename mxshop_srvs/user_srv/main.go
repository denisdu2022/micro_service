package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"user_srv/handler"
	"user_srv/proto"
)

func main() {
	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 8080, "端口号")

	flag.Parse()
	fmt.Println("ip: ", *IP)
	fmt.Println("port: ", *Port)

	//获取grcp的server
	server := grpc.NewServer()
	//注册grpc
	proto.RegisterUserServer(server, &handler.UserServer{})

	//启动监听
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen: " + err.Error())
	}

	//绑定监听
	err = server.Serve(lis)
	if err != nil {
		panic("failed to start grpc: " + err.Error())
	}

}
