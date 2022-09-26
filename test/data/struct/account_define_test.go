package main

import "testing"

func TestAccount(t *testing.T) {

	account := &Account{
		AccountNo: "zb128398998",
		Pwd:       "1qaz@WSX",
		Balance:   1000.0,
	}

	account.query("1qaz@WSX")
	account.Deposit(5000.10, "1qaz@WSX")
	account.query("1qaz@WSX")
	account.withdraw(1234.9, "1qaz@WSX")
	account.query("1qaz@WSX")
}
