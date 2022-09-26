package main

import "fmt"

type Account struct {
	AccountNo string
	Pwd       string
	Balance   float64
}

// 存款
func (ac *Account) Deposit(money float64, pwd string) {
	if ac.Pwd != pwd {
		fmt.Println("密码错误！")
		return
	}

	if money <= 0 {
		fmt.Println("金额错误")
		return
	}

	ac.Balance += money
	fmt.Println("存款成功")
}

// 取款
func (ac *Account) withdraw(money float64, pwd string) {
	if ac.Pwd != pwd {
		fmt.Println("密码错误！")
		return
	}

	if money <= 0 || ac.Balance < money {
		fmt.Println("金额错误")
		return
	}

	ac.Balance -= money
	fmt.Println("取款成功")
}

// 查询
func (ac *Account) query(pwd string) {
	if ac.Pwd != pwd {
		fmt.Println("密码错误！")
		return
	}

	fmt.Printf("账户余额为%.2f\n", ac.Balance)
}
