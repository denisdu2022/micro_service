package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {

	//返回格式:json {"data":3}
	//1.callID的问题:r.URL.Path
	//2.数据传输协议(URL的参数传递协议) http的协议
	//http://127.0.0.1:8000/add?a=1&b=2
	//3.

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		_ = r.ParseForm() //解析参数
		fmt.Println("Path: ", r.URL.Path)

		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		//设置header
		w.Header().Set("Content-Type", "application/json")
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})

		//返回jData给客户端
		_, _ = w.Write(jData)
	})

	//监听:服务端启动
	http.ListenAndServe(":8000", nil)
}
