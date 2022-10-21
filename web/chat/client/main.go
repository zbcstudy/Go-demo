package main

import "fmt"

var userId int
var password string

func main() {
	var key int // 接收用户的选择
	var loop = true

	for loop {
		fmt.Println("-------------欢迎登陆聊天系统-------------")
		fmt.Println("\t\t\t 1.登陆聊天室")
		fmt.Println("\t\t\t 2.注册用户")
		fmt.Println("\t\t\t 3.退出系统")
		fmt.Println("\t\t\t 请选择1-3")

		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			loop = false
		case 2:
			fmt.Println("注册用户")
			loop = false
		case 3:
			loop = false
			fmt.Println("退出系统")
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}

	if key == 1 {
		// 登陆界面
		fmt.Println("请输入用户id：")
		fmt.Scanf("%d\n", &userId)
		fmt.Println("请输入密码：")
		fmt.Scanf("%s\n", &password)
		err := Login(userId, password)

		if err != nil {
			fmt.Println("登陆失败")
		} else {
			fmt.Println("登陆成功")
		}

	} else if key == 2 {
		fmt.Println("进行用户注册的逻辑")
	}
}
