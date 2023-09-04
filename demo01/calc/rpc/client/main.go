package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

//RPC的远程调用,如何做到像本地调用

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a, b int) int {
	//传输协议http
	req := HttpRequest.NewRequest() //初始化req

	res, _ := req.Get(fmt.Sprintf("http://127.0.0.1:8000/%s?a=%d&b=%d", "add", a, b)) //get请求完整的URL

	body, _ := res.Body()

	fmt.Println("[]byte body >>> ", body) //body是[]byte类型

	fmt.Println("string body", string(body))

	rspData := ResponseData{}

	_ = json.Unmarshal(body, &rspData)

	return rspData.Data
}

func main() {
	//req := HttpRequest.NewRequest() //初始化req
	//
	//res, _ := req.Get("http://127.0.0.1:8000/add?a=1&b=2") //get请求完整的URL
	//
	//body, _ := res.Body()
	//
	//fmt.Println("[]byte body >>> ", body) //body是[]byte类型
	//
	//fmt.Println("string body", string(body))

	result := Add(2, 3)

	fmt.Println("result >>>", result)
}
