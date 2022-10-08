package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		// 等待客户端连接
		accept, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err:", err)
			panic(err)
		}
		// 一旦有连接 立马开启go处理连接
		go handler(accept)

	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn read error:", err)
		return
	}
	fmt.Println("req:", string(buf[:n]))

	// 往连接中写入数据
	_, err = conn.Write([]byte("hello get request success"))
	if err != nil {
		fmt.Println("conn write error:", err)
	}

}
