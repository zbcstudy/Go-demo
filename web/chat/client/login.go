package main

import (
	"data/web/chat/common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"strconv"
	"time"
)

func Login(userId int, password string) error {

	fmt.Printf("userId:%v,password:%v\n", userId, password)

	// 连接到服务器端
	conn, err := net.Dial("tcp", ":8889")
	if err != nil {
		fmt.Println("net dial err:", err)
	}

	defer conn.Close()

	// 消息结构体
	var message common.Message
	message.Type = common.LoginMesType

	// login结构体
	var loginMes = common.LoginMes{
		UserId:  strconv.Itoa(userId),
		UserPwd: password,
	}

	bytes, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json marshal loginMes err:", err)
		return err
	}
	message.Data = string(bytes)

	// 对message进行序列化
	mesData, err := json.Marshal(message)
	if err != nil {
		fmt.Println("json marshal message err:", err)
		return err
	}

	// mesData 就是需要发送的消息
	// 先发数据的长度
	dataLength := uint32(len(mesData))
	var buf [4]byte

	// 将长度转换为切片
	binary.BigEndian.PutUint32(buf[:4], dataLength)

	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn write(bytes) fail:", err)
		return err
	}
	fmt.Printf("客户端发送消息的长度：%d,内容：%v\n", len(mesData), string(mesData))

	// 发送消息本身
	_, err = conn.Write(mesData)
	if err != nil {
		fmt.Println("conn write(data) err:", err)
		return err
	}

	// 休眠20s
	time.Sleep(time.Second * 20)
	// 处理服务器端返回的信息

	return nil
}
