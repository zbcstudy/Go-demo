package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("net dial error:", err)
	}
	defer listen.Close()
	for {
		fmt.Println("开始进行监听")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("list accept error:", err)
			return
		}
		fmt.Println("conn:", conn.RemoteAddr())

		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		// 切片
		buf := make([]byte, 1024)
		// 读取数据
		fmt.Println("服务器端等待客户端发送信息")
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read error:", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
