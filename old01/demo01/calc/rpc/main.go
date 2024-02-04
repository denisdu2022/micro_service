package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

type Company struct {
	Name    string
	Address string
}

type Employee struct {
	Name    string
	company Company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) {
	//rpc中的第二个点:传输协议,数据编码协议
	//http1.x  http2.0协议  http底层也是使用的tcp协议
	/*
			客户端:
				1.建立连接 tcp/http
				2.将employee序列化成json字符串-序列化
				3.发送json字符串-调用成功后实际上接收到的是一个二进制的数据
				4.等待服务器发送结果
				5.将服务器返回的数据解析成PrintResult 对象 -反序列化
			服务端:
				1.监听网络端口80
				2.读取数据- 二进制数据
				3.对读取的数据进行反序列化Employee对象
				4.开始处理业务逻辑
				5.将处理的结果PrintResult 序列化成json二进制数据 -序列化
				6.将服务器返回的数据

		序列化和反序列化是可以选择的,不一定采用json,xml,protobuf,msgpack
	*/

}

func main() {
	//现在想把Add函数变成一个远程的调用函数,也就需要把Add函数放在远程服务器上去运行
	/*
			电商系统,在这里有一段逻辑,这个逻辑是扣减库存,但是库存服务是一个独立的系统,reduce ,那如何调用
			一定会涉及到网络,做成一个web 服务(gin,beego,net/httpserver)
		1.这个函数的调用参数如何传递-json (json是一种数据格式协议,并不是高性能的协议) /xml/protobuf/msgpack
			现在网络调用有两个端,-客户端:应该将数据传输到gin
			gin-服务端,服务端负责解析

	*/
	fmt.Println(Add(2, 3))

	//将这个打印的工作放在另一台服务器上,我需要将本地的内存对象 struct 序列化成二进制对象才能进行网络传输

	fmt.Println(Employee{
		Name: "jack",
		company: Company{
			Name:    "bank",
			Address: "北京市",
		},
	})
	//远程服务需要将二进制对象反解成struct对象
}
