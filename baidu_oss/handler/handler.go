package handler

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传HTML页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			io.WriteString(w, "internal server error")
			return
		}
		io.WriteString(w, string(data))
	} else if r.Method == "POST" {
		// 接收文件流 存储到本地
		file, header, err := r.FormFile("file")
		if err != nil {
			fmt.Printf("fail to get data,err:%s\n", err.Error())
			return
		}
		defer file.Close()
		newFile, err := os.Create("/temp/" + header.Filename)
		if err != nil {
			fmt.Printf("fail to create file,err:%s\n", err.Error())
			return
		}
		defer newFile.Close()
		_, err = io.Copy(newFile, file)
		if err != nil {
			fmt.Printf("fail to save data into file,err:%s\n", err.Error())
			return
		}
	}
	// 重定向
	http.Redirect(w, r, "/file/upload/success", http.StatusFound)
}

// 上传完成
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload finished!")
}
