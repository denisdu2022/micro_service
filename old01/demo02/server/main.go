package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	//http://127.0.0.1:8000/add?a=1&b=2
	//返回的格式化: json {"data":3}
	//1. callID的问题: r.URL.Path 2. 数据的传输协议: URL的参数传递协议 3. 网络传输协议:http
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		//解析参数
		_ = r.ParseForm()
		fmt.Println("path: ", r.URL.Path)
		a, _ := strconv.Atoi(r.Form["a"][0])
		b, _ := strconv.Atoi(r.Form["b"][0])
		//设置Header
		w.Header().Set("Content-Type", "application/json")
		//序列化
		jData, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = w.Write(jData)
	})

	//启动监听
	http.ListenAndServe(":8000", nil)
}
