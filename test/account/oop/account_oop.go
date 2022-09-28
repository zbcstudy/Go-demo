package main

import "fmt"

type FamilyAccount struct {
	key  string
	loop bool
	// 账户余额
	balance float64
	// 每次收支
	money float64
	// 说明
	note string
	// 收支详情 "收支\t账户金额\t收支金额\t说明"
	details string

	flag bool
}

func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		note:    "",
		details: "收支\t账户金额\t收支金额\t说明",
		flag:    false,
	}
}

func (fa *FamilyAccount) ShowDetails() {
	fmt.Println("---------------家庭收支记账软件---------------")
	if !fa.flag {
		fmt.Println("当前没有收支明细")
	} else {
		fmt.Println(fa.details)
	}
}

func (fa *FamilyAccount) AddBalance() {
	fmt.Println("本次收入金额：")
	_, _ = fmt.Scanln(&fa.money)
	fa.balance += fa.money //修改余额
	fmt.Println("收入说明")
	_, _ = fmt.Scanln(&fa.note)

	fa.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
	fa.flag = true
}

func (fa *FamilyAccount) DecreaseBalance() {
	fmt.Println("本次支出金额：")
	_, _ = fmt.Scanln(&fa.money)
	// 必要判断
	if fa.money > fa.balance {
		fmt.Println("余额不足")
		fa.flag = true
		return
	}
	fa.balance -= fa.money

	fmt.Println("支出说明")
	_, _ = fmt.Scanln(&fa.note)

	fa.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", fa.balance, fa.money, fa.note)
}

func (fa *FamilyAccount) Exist() {
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
		fa.loop = false
	}
}

func (fa *FamilyAccount) MainMenu() {
	for {
		fmt.Println("---------------家庭收支记账软件---------------")
		fmt.Println("               1 收支明细")
		fmt.Println("               2 登记收入")
		fmt.Println("               3 登记支出")
		fmt.Println("               4 退出软件")
		fmt.Println("请选在1-4")

		_, _ = fmt.Scanln(&fa.key)
		switch fa.key {
		case "1":
			fa.ShowDetails()
		case "2":
			fa.AddBalance()
		case "3":
			fa.DecreaseBalance()
		case "4":
			fa.Exist()
		default:
			fmt.Println("请输入正确数字")
		}
		if !fa.loop {
			break
		}
	}
}
