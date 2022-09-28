package main

import (
	"fmt"
)

func main() {
	key := ""
	loop := true

	// 账户余额
	balance := 10000.0
	// 每次收支
	money := 0.0
	// 说明
	note := ""
	// 收支详情
	details := "收支\t账户金额\t收支金额\t说明"

	flag := false
	for {
		fmt.Println("---------------家庭收支记账软件---------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Println("请选在1-4")

		_, _ = fmt.Scanln(&key)
		switch key {
		case "1":
			if !flag {
				fmt.Println("当前没有收支明细")
				break
			}
			fmt.Println("---------------家庭收支记账软件---------------")
			fmt.Println(details)
		case "2":
			fmt.Println("本次收入金额：")
			_, _ = fmt.Scanln(&money)
			balance += money //修改余额
			fmt.Println("收入说明")
			_, _ = fmt.Scanln(&note)

			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Println("本次支出金额：")
			_, _ = fmt.Scanln(&money)
			// 必要判断
			if money > balance {
				fmt.Println("余额不足")
				flag = true
				break //只是推出switch
			}
			balance -= money

			fmt.Println("支出说明")
			_, _ = fmt.Scanln(&note)

			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, note)
		case "4":
			fmt.Println("确定要退出吗？y/n")
			chose := ""
			for {
				fmt.Scanln(&chose)
				if chose == "y" || chose == "n" {
					break
				} else {
					fmt.Println("确定要退出吗？y/n")
				}

			}
			if chose == "y" {
				loop = false
			}

		default:
			fmt.Println("请输入正确数字")
		}
		if !loop {
			break
		}
	}

}
