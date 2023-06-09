package main

import (
	"baidu_oss/handler"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", handler.UploadHandler)
	http.HandleFunc("/file/upload/success", handler.UploadSuccessHandler)
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("服务启动失败", err)
	}
}
