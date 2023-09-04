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
	Company Company
}

type PrintResult struct {
	Info string
	Err  error
}

func RpcPrintln(employee Employee) {
	/*
		客户端:
				1.建立连接tcp/http
				2.将employee 对象序列化成json字符串-序列化
				3.发送json字符串-调用成功后实际上接收到的是一个二进制的数据
				4.等待服务器发送结果
			    5.将服务器返回的数据解析成PrintResult 对象 -反序列化
		服务端:
				1.监听网络端口 80/8080
				2.读取数据 - 二进制的json数据
				3.对接收到的数据进行反序列化Employee对象
				4.开始处理业务逻辑
				5.将处理的结果PrintResult序列化成二进制数据,进行网络传输
				6.将数据返回
		序列化和反序列化是可以选择的,并不是固定的,不一定要采用json,xml,protobuf,msgpack


		RPC重要的两个点: 传输协议,数据编码协议
		http协议 1.x http2.0
		http协议底层使用的也是tcp协议
		http1.x 短连接,一次性,一旦结果返回连接就断开了
		http2.0既有http的特性,又有长连接的特性 grpc使用的就是http2.0
	*/

}

func main() {

	//调用函数
	//需求:想要把Add函数变成一个远程调用函数,也就是需要把Add函数放在远程服务器上去运行 ,把本地函数调用,改为远程调用

	/*
		1.这个函数的调用参数如何传递-json(json是一种数据格式的协议)或者xml/protobuf/msgpack 编码协议,json并不是一个高效的编码协议
		现在网络调用有两个端 -客户端,服务端
		客户端: 应该将数据传输到gin
		gin服务端:服务端负责解析数据

	*/

	fmt.Println(Add(2, 3))

	//将这个打印的工作放在另一台服务器上,需要将本地内存对象struct传输,但是这样是不行的,可行的方式就是将struct序列化成json
	fmt.Println(Employee{
		Name: "jack",
		Company: Company{
			Name:    "黑马",
			Address: "北京市",
		},
	})
	//远程服务器需要将网络传输来的二进制对象反序列化成struct对象

}
