package main

import (
	"fmt"
	"sync"
	"time"
)

func ExpensiveFibonacci(n int) int {
	fmt.Println("Calculate Expensive Fibonacci", n)
	time.Sleep(5 * time.Second)
	return n
}

type Service struct {
	InProgress map[int]bool
	IsPending  map[int][]chan int
	Mu         sync.RWMutex
}

func (s *Service) Work(job int) {
	s.Mu.RLock()
	exists := s.InProgress[job]
	if exists {
		s.Mu.RUnlock()
		response := make(chan int)
		defer close(response)

		s.Mu.Lock()
		s.IsPending[job] = append(s.IsPending[job], response)
		s.Mu.Unlock()

		fmt.Println("Waiting for response job", job)
		resp := <-response
		fmt.Println("response Done, received", resp)
		return
	}
	s.Mu.RUnlock()
	s.Mu.Lock()
	s.InProgress[job] = true
	s.Mu.Unlock()

	fmt.Println("Calculate Fibunacci for ", job)
	result := ExpensiveFibonacci(job)

	s.Mu.RLock()
	pendingWorkers, exists := s.IsPending[job]
	s.Mu.RUnlock()

	if exists {
		for _, pendingWorker := range pendingWorkers {
			pendingWorker <- result
		}
		fmt.Println("Result sent for all workers", job)
	}
	s.Mu.Lock()
	s.InProgress[job] = false
	s.IsPending[job] = make([]chan int, 0)
	s.Mu.Unlock()
}

func NewService() *Service {
	return &Service{
		InProgress: make(map[int]bool),
		IsPending:  make(map[int][]chan int),
	}
}

func main() {
	service := NewService()
	jobs := []int{3, 4, 5, 5, 4, 8, 8, 8}
	var wg sync.WaitGroup

	for _, job := range jobs {
		wg.Add(1)
		go func(job int) {
			defer wg.Done()
			service.Work(job)
		}(job)
	}

	wg.Wait()
}
