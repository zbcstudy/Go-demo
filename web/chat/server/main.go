package main

import (
	"data/web/chat/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
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

// 处理消息的方法
// 不同的消息 对应不同的处理业务
func serverProcessMes(conn net.Conn, mes *common.Message) (err error) {
	switch mes.Type {
	case common.LoginMesType:
	// 处理登录逻辑
	default:
		fmt.Println("消息无法处理")
	}
}

func process(conn net.Conn) {
	defer conn.Close()
	for {
		message, err := readPkg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端退出了连接")
				return
			} else {
				fmt.Println("conn read err:", err)
			}
		}
		fmt.Println("message:", message)
	}
}

func readPkg(conn net.Conn) (message common.Message, err error) {
	buf := make([]byte, 1024)
	// 读取客户端发送的数据
	_, err = conn.Read(buf[:4])
	if err != nil {
		fmt.Println("conn read error:", err)
		return
	}
	fmt.Println("读到的buf=", buf[:4])

	// 读取长度
	var pkgLength uint32 = binary.BigEndian.Uint32(buf[:4])
	n, err := conn.Read(buf[:pkgLength])
	if n != int(pkgLength) || err != nil {
		fmt.Println("conn read fail error:", err)
		return
	}

	// 将数据 反序列化
	err = json.Unmarshal(buf[:pkgLength], &message)
	return
}
