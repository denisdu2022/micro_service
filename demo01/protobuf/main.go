package main

import (
	"fmt"
	"githun.com/golang/protobuf/proto"
	"protobuf01/proto"
)

func main() {
	//获取HelloRequest对象并实例化
	req := proto.HelloRequest{
		Name: "李白",
	}

	//序列化
	proto.Marshal()

	fmt.Println(req)
}
