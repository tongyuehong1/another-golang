package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// 设置访问的路由
	http.HandleFunc("/hello", handleHello)
	// 打印在terminal中
	fmt.Println("serving on http://localhost:7777/hello")
	// http.ListenAndServe设置设置监听的端口
	log.Fatal(http.ListenAndServe("localhost:7777", nil))
}

func handleHello(w http.ResponseWriter, req *http.Request) {
	// 每次进入这个网页，都会在terminal中打印
	log.Println("serving", req.URL)
	// 这个写入到w的是输出到客户端的
	fmt.Fprintln(w, "Hello, 世界!")
}
