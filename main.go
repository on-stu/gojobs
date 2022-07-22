package main

import (
	"fmt"

	"github.com/on-stu/gojobs/accounts"
)

func main() {
	account := accounts.NewAccount("minsu")
	fmt.Println(account.Balance())
	account.Deposit(10000)
	fmt.Println(account.Balance())

	fmt.Println(account)
}
