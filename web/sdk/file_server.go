package main

import (
	"fmt"
	"net/http"
)

// 一个简单的文件服务器
func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("服务启动成功")
	}
}
