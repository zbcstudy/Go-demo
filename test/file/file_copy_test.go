package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func CopyFile(dest string, src string) (int64, error) {
	// 获取源文件
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("源文件获取失败")
		return 0, nil
	}

	defer srcFile.Close()

	// 获取源文件流
	reader := bufio.NewReader(srcFile)

	// 获取目标文件
	destFile, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		// 目标文件创建失败
		fmt.Println("目标文件创建失败")
		return 0, nil
	}
	defer destFile.Close()
	writer := bufio.NewWriter(destFile)

	// 成功的情况下 返回的是文件的大小
	return io.Copy(writer, reader)
}

func TestFileCopy(t *testing.T) {
	count, _ := CopyFile("E:\\tmp\\copy\\copy.jpeg", "E:\\tmp\\src1.jpeg")
	println(count)
}
