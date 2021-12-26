package main

import (
	"fmt"

	"github.com/Hwisaek/Go/nomadcoder_go-for-beginners/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(2)
	fmt.Println(account.Balance())
}
