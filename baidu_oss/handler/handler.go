package handler

import (
	"github.com/containerd/containerd"
	"io/ioutil"
	"net/http"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		// 返回上传HTML页面
		data, err := ioutil.ReadFile("./static/view/index.html")
		if err != nil {
			_ = ioutil.WriteFile("index.html", data, 'w')
		}
	} else if r.Method == "POST" {

	}
	containerd.New()
	w.Write([]byte("1234"))
}
