package main

import (
	"fmt"
	"log"

	"github.com/on-stu/gojobs/accounts"
)

func main() {
	account := accounts.NewAccount("minsu")
	fmt.Println(account.Balance())
	account.Deposit(10000)
	fmt.Println(account.Balance())
	err := account.Withdraw(100001)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account.Balance())
}
