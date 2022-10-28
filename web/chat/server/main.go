package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("服务器在8889端口监听")

	listen, err := net.Listen("tcp", ":8889")
	defer listen.Close()

	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}

	//等待客户端连接
	for {
		fmt.Println("等待客户端连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen accept err:", err)
		}

		// 获取连接成功 起一个协程保持通讯
		go process(conn)
	}

}

func process(conn net.Conn) {
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		// 读取客户端发送的数据
		n, err := conn.Read(buf[:4])
		if n != 4 || err != nil {
			fmt.Println("conn read error:", err)
			return
		}
		fmt.Println("读到的buf=", buf[:4])
	}
}
