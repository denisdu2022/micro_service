package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"protobuf_test/proto"
)

//protobuf和json的直观对比

type Hello struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Courses []string `json:"courses"`
}

func main() {
	//protobuf 实例化
	req := helloworld.HelloRequest{
		Name:    "李白",
		Age:     23,
		Courses: []string{"go", "gin", "微服务"},
	}
	rsp, _ := proto.Marshal(&req) //具体的编码是如何做到的:protobuf的原理 varint
	//fmt.Println("[]byte rsp: ", rsp)
	//fmt.Println("string rsp: ", string(rsp))
	//fmt.Println("len rsp: ", len(rsp))

	//json实例化
	//jsonStruct := Hello{
	//	Name:    "李白",
	//	Age:     23,
	//	Courses: []string{"go", "gin", "微服务"},
	//}
	//
	//jsonRsp, _ := json.Marshal(jsonStruct)
	//fmt.Println("[]byte jsonRsp: ", jsonRsp)
	//fmt.Println("string jsonRsp: ", string(jsonRsp))
	//fmt.Println("len jsonRsp: ", len(jsonRsp))

	//反向解析
	newReq := helloworld.HelloRequest{}
	_ = proto.Unmarshal(rsp, &newReq)
	fmt.Println(newReq.Name, newReq.Age, newReq.Courses)
}
