package file

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestOpenFile(t *testing.T) {
	// 在Windows下 第三个参数不会生效
	file, _ := os.OpenFile("E:\\tmp\\test2.txt", os.O_CREATE|os.O_WRONLY, 0666)
	println(file)

	str := "需要写入文件的信息,abc"

	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		_, _ = writer.WriteString(str + strconv.Itoa(i) + "\r\n")
	}
	err := writer.Flush()
	if err != nil {
		fmt.Println("write success")
	}

}
