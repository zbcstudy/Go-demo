package file

import (
	"fmt"
	"io/ioutil"
	"testing"
)

// 一次將文件内容全部读取出来 适用文件不大的情况
func TestFileReadAll(t *testing.T) {
	file, err := ioutil.ReadFile("E:\\tmp\\test.txt")
	if err != nil {
		fmt.Println("read file err")
	}
	print(string(file))
}
