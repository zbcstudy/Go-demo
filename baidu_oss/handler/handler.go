package handler

import (
	"baidu_oss/meta"
	"baidu_oss/util"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
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

		// 获取当前路径
		path, _ := os.Getwd()

		// 条件文件元数据
		fileMeta := meta.FileMeta{
			FileSha1: "",
			FileName: header.Filename,
			FileSize: 0,
			Location: path + "/temp/" + header.Filename,
			UploadAt: time.Now().Format("2006-01-01 15:04:05"),
		}

		newFile, err := os.Create(fileMeta.Location)
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

		// 文件游标移动到文件的任意位置
		newFile.Seek(0, 0)
		sha1 := util.FileSha1(newFile)
		fmt.Println(sha1)
		fileMeta.FileSha1 = sha1

		meta.UpdateFileMeta(fileMeta)

		// 重定向
		http.Redirect(w, r, "/file/upload/success", http.StatusFound)
	}

}

// 上传完成
func UploadSuccessHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "upload finished!")
}

// 获取文件meta信息
func GetFileMetaHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form["fileHash"][0]
	fileMeta := meta.GetFileMeta(fileHash)
	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(data)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form.Get("fileHash")
	fileMeta := meta.GetFileMeta(fileHash)

	f, err := os.Open(fileMeta.Location)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("disposition", "attachment;filename=\""+fileMeta.FileName+"\"")
	w.Write(data)
}

// 更新文件meta信息
func FileMetaUpdateHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	opType := r.Form.Get("op")
	fileHash := r.Form.Get("fileHash")
	filename := r.Form.Get("filename")

	if opType != "0" {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 只支持POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// 获取当前meta
	currentFileMeta := meta.GetFileMeta(fileHash)
	currentFileMeta.FileName = filename
	// 更新
	meta.UpdateFileMeta(currentFileMeta)

	data, err := json.Marshal(currentFileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func FileDeleteHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileHash := r.Form.Get("fileHash")

	// 删除文件
	fileMeta := meta.GetFileMeta(fileHash)
	os.Remove(fileMeta.Location)

	// 删除meta
	meta.RemoveFileMeta(fileHash)

	w.WriteHeader(http.StatusOK)

}
