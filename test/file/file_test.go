package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFileCreate(t *testing.T) {
	file, _ := os.Create("test.txt")
	fmt.Println(file.Name())
	// 文件写入 返回写入数据的长度
	write, err2 := file.Write([]byte("first write"))
	if err2 != nil {
		fmt.Println("write file err:", err2)
	} else {
		fmt.Println("write:", write)
	}
	_ = file.Close()

	diskFile, err := os.Open("E:\\tmp\\test.txt")

	defer diskFile.Close()

	if err != nil {
		fmt.Println("open file error=", err)
	}
	fmt.Printf("diskFile:%v\n", diskFile)

	reader := bufio.NewReader(diskFile)

	for {
		// 读到换行结束
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(content)
	}

}
