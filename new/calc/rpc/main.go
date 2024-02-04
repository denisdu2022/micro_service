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
	//客户端:
	//1.建立连接:tcp/http
	//2.将employee对象序列化成json字符串-序列化
	//3.发送json字符串 -调用成功后实际上你接収到的是一个二进制的数据
	//4.等待服务器发送结果
	//5.将服务器返回的数据解析成printResult对象 -反序列化

	//服务端:
	//1.监听网络端口80
	//2.读取数据-二进制的json数据
	//3.对数据进行解析:反序列化
	//4.开始处理业务逻辑
	//5.将处理的结果PrintResult序列化成json二进制数据-序列化
	//6.将处理后的数据返回

	//序列化和反序列化是可以选择的,不一定采用json,xml,protobuf,msgpack

}

func main() {
	//现在想把Add函数变成一个远程的函数调用,也就意味着需要把Add函数放在远程服务器上去运行

	/*
		原本的电商系统,这里有一段逻辑,这个逻辑是扣减库存,但是库存服务是一个独立的系统,这就需要远程调用

		1.这个函数的调用参数该如何传递-json (json是一种数据格式的协议)/xml/protobuf/msgpack
		现在网络调用有两个端-客户端,应该做什么?将数据传输到gin
		gin服务端,服务端负责解析数据
		(json并不是高性能的编码协议)


	*/
	fmt.Println(Add(1, 2))

	fmt.Println(Employee{
		Name: "李白",
		company: Company{
			Name:    "大唐",
			Address: "长安街",
		},
	})
}
