//ex33/ex33.6/ex33.6.go
package main

import (
	"fmt"

	"github.com/tuckersGo/musthaveGo/ex33/ex33.6/bankaccount"
)

func main() {
	account := bankaccount.NewAccount()
	account.Deposit(1000)
	fmt.Println(account.Balance())
}