package accounts

import "errors"

type account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *account {
	account := account{owner: owner, balance: 0}
	return &account
}

func (a *account) Deposit(amount int) { //method를 작성하는 법 func 뒤에 있는 걸 receiver라고 부름, a로 받아오는 건 일종의 rule임
	// a는 this나 self 처럼 이용 가능
	a.balance += amount
}

func (a account) Balance() int {
	return a.balance
}

var errNoMoney = errors.New("Can't withdraw")

// Withdraw x amount from your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
