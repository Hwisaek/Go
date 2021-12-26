package main

import (
	"fmt"

	"github.com/Hwisaek/Go/nomadcoder_go-for-beginners/accounts"
)

func main() {
	account := accounts.NewAccount("nico")
	account.Deposit(2)
	fmt.Println(account.Balance())
	err := account.WithDraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.Balance())
}
