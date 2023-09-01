package main

import (
	"context" //上下文   goroutine(go 程) 之间用来进行数据传递api包
	"demo02/pb"
	"google.golang.org/grpc"
	"net"
)

//定义类

type Children struct {
}

//接口绑定类方法

func (c *Children) SayHello(ctx context.Context, teacher *pb.Teacher) (*pb.Teacher, error) {
	reply := &pb.Teacher{
		Name: "您好," + teacher.Name,
		Age:  20,
	}

	return reply, nil
}

func main() {

	//初始化Grpc对象
	grpcServer := grpc.NewServer()
	//注册grpc服务
	//pb.RegisterSayNameServer(grpcServer, new())

	//设置服务端监听
	lis, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}

	defer lis.Close()

	//绑定服务
	grpcServer.Serve(lis)

}
