package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 0
)

func Deposit(amount int, wg *sync.WaitGroup, lock *sync.Mutex) {
	defer wg.Done()
	defer lock.Unlock()

	lock.Lock()
	b := balance
	balance = b + amount
}

func Balance() int {
	b := balance
	return b
}

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex

	for i := 1; i <= 10000; i++ {
		wg.Add(1)
		go Deposit(i, &wg, &lock)
	}

	wg.Wait()
	// 10000(10000 + 1) / 2 = 50005000
	// To tell if we have race conditions we can use
	// go build --race main.go
	fmt.Println("Current balance", Balance())
}
