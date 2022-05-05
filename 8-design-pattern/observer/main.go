package main

import (
	"fmt"
	"sync"
	"time"
)

func sendWelcomeEmail(value string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Sending welcome email to", value)
}

func addNewUserDiscount(value string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Add new user discount to", value)
}

func main() {
	println("Demo synchronous observer")
	syncRegisterTopic := NewSyncRegisterTopic()
	syncRegisterTopic.Register(sendWelcomeEmail)
	syncRegisterTopic.Register(addNewUserDiscount)
	for i := 0; i < 10; i++ {
		syncRegisterTopic.Broadcast(fmt.Sprintf("john.doe.%d@domain.com", i))
	}

	fmt.Println("Demo asynchronous observer")

	// Initialize topic
	asyncRegisterTopic := NewAsyncRegisterTopic()
	asyncRegisterTopic.Register(sendWelcomeEmail)
	asyncRegisterTopic.Register(addNewUserDiscount)

	// Start topic
	var wg sync.WaitGroup
	wg.Add(1)
	go asyncRegisterTopic.BroadcastWorker(&wg)

	// Send values to the topic
	for i := 0; i < 10; i++ {
		asyncRegisterTopic.Broadcast(fmt.Sprintf("john.doe.%d@domain.com", i))
	}

	// Do work while the topic is working in parallel
	for i := 0; i < 10; i++ {
		fmt.Println("I can do more work while asynchronous observer is working")
	}

	// Close the topic and wait
	asyncRegisterTopic.Close()
	wg.Wait()
}
