package function

import (
	"errors"
	"testing"
)

// 自定义错误
func readFile(name string) (err error) {
	if name == "config.ini" {
		return nil
	} else {
		// 自定义错误
		return errors.New("读取文件错误")
	}
}

func TestError(t *testing.T) {
	err := readFile("config.config")
	if err != nil {
		// 如果读取文件发生错误，输出错误 并终止代码执行
		panic(err)
	}
	println("程序执行完毕")
}
