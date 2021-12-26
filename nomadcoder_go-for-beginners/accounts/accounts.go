package accounts

import "errors"

type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a *Account) Balance() int {
	return a.balance
}

// WithDraw x amount from your account
func (a *Account) WithDraw(amount int) error {
	if a.balance < amount {
		return errors.New("Can't withdraw. You're poor")
	}
	a.balance -= amount
	return nil
}
