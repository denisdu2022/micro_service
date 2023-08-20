package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

//定义类对象

type World struct {
}

//绑定类方法

func (w *World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好!"
	return nil
	//修改err 返回对象
	//errors 是错误包,New是New的一个错误对象
	//return errors.New("未知的错误!")
}

func main() {
	//1.注册RPC服务,绑定对象方法
	//err := rpc.RegisterName("hello", new(World))
	//if err != nil {
	//	fmt.Println("注册rpc服务失败!  >>> ", err)
	//	return
	//}

	//调用RegisterServer,传入类对象
	RegisterServer(new(World))

	//2.设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.Listen err ! >>> ", err)
		return
	}

	//先提前调起listener.Close()函数,在当前函数释放占用空间之前执行
	defer listener.Close()
	fmt.Println("开始监听,等待建立连接 .....")

	//3.建立连接
	conn, err := listener.Accept()

	if err != nil {
		fmt.Println("listener.Accept err ! >>> ", err)
		return
	}

	defer conn.Close()

	fmt.Println("连接建立成功 .....")

	//4.绑定服务
	//注意ServerConn()函数是没有返回值的
	//rpc.ServeConn(conn)

	//使用jsonRpc 绑定服务
	jsonrpc.ServeConn(conn)

}
