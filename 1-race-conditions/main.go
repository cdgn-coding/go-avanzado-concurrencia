package main

import (
	"fmt"
	"sync"
)

var balance = 500

func Deposit(value int) {
	balance = value + balance
}

func Withdraw(value int) error {
	if value > balance {
		return fmt.Errorf("invalidad value")
	}
	balance = balance - value
	return nil
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		Deposit(200)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		Deposit(100)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println(balance)
}
