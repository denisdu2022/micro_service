package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"stream_grpc_test/proto"
	"sync"
	"time"
)

func main() {
	//建立连接
	conn, err := grpc.Dial("localhost:8000", grpc.WithInsecure())
	//错误处理
	if err != nil {
		panic(err)
	}

	defer conn.Close()

	//服务端流模式
	//获取客户端
	c := proto.NewGreeterClient(conn)
	//调用
	res, _ := c.GetStream(context.Background(), &proto.StreamReqData{Data: "李白"})
	for {
		a, err := res.Recv() //socket 中的send recv
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("调用返回的值: ", a)
	}

	//客户端流模式
	putStream, _ := c.PutStream(context.Background())
	i := 0
	for {
		i++
		putStream.Send(&proto.StreamReqData{
			Data: fmt.Sprintf("第%d次", i),
		})
		time.Sleep(time.Second)
		if i > 10 {
			break
		}
	}

	//双向流模式
	allStr, _ := c.AllStream(context.Background())
	//使用协程
	wg := sync.WaitGroup{}
	wg.Add(2)
	//接收客户端消息
	go func() {
		defer wg.Done()
		for {
			data, _ := allStr.Recv()
			fmt.Println("收到服务端消息: " + data.Data)
		}
	}()

	//向客户端发送消息
	go func() {
		defer wg.Done()
		for {
			_ = allStr.Send(&proto.StreamReqData{
				Data: "我是客户端",
			})
			time.Sleep(time.Second)
		}
	}()

	wg.Wait()

}
