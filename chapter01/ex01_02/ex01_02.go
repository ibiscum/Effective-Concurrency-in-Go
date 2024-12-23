package main

import (
	"log"
	"time"
)

type Account struct {
	Balance int
}

var ErrInsufficient error

func (acct *Account) Withdraw(amt int) error {
	if acct.Balance < amt {
		return ErrInsufficient
	}

	acct.Balance -= amt
	return nil
}

func main() {
	acct := Account{Balance: 10}

	go func() {
		err := acct.Withdraw(6)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("go 1: ", acct.Balance)
		}
	}() // Goroutine 1

	go func() {
		err := acct.Withdraw(5)
		if err != nil {
			log.Fatal(err)
		} else {
			log.Println("go 2: ", acct.Balance)
		}
	}() // Goroutine 2

	time.Sleep(2 * time.Second)
}
