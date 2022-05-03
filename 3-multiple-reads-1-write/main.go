package main

import (
	"fmt"
	"sync"
)

var (
	balance int = 0
)

// Writer
func Deposit(amount int, wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	defer lock.Unlock()

	lock.Lock()
	b := balance
	balance = b + amount
}

// Reader
func Balance(wg *sync.WaitGroup, lock *sync.RWMutex) {
	defer wg.Done()
	defer lock.RUnlock()
	lock.RLock()
	b := balance
	fmt.Println("Current balance is", b)
}

func main() {
	var wg sync.WaitGroup
	// Multiple readers, one Writer
	var lock sync.RWMutex

	for i := 1; i <= 20; i++ {
		wg.Add(1)
		go Deposit(i, &wg, &lock)
		wg.Add(1)
		go Balance(&wg, &lock)

	}

	wg.Wait()
	fmt.Println("Finished with balance", balance)
}
