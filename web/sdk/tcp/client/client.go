package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		fmt.Println("client dial error:", err)
		return
	}
	// 客户端可以发送单行数据
	reader := bufio.NewReader(os.Stdin)

	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("read string error:", err)
		return
	}
	_, err = conn.Write([]byte(str))
	if err != nil {
		fmt.Println("conn write error:", err)
	}
}
