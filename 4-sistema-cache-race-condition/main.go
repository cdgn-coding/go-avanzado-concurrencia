package main

import (
	"fmt"
	"sync"
	"time"
)

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}

	return Fibonacci(n-1) + Fibonacci(n-2)
}

type Memory struct {
	f     Function
	cache map[int]FunctionResult
}

type Function func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

func NewCache(f Function) *Memory {
	cache := make(map[int]FunctionResult)
	return &Memory{f: f, cache: cache}
}

func (memory *Memory) Get(key int) (interface{}, error) {
	result, exists := memory.cache[key]

	if !exists {
		result.value, result.err = memory.f(key)
		memory.cache[key] = result
	}

	return result.value, result.err
}

func GetFibonacci(n int) (interface{}, error) {
	return Fibonacci(n), nil
}

func main() {
	cache := NewCache(GetFibonacci)
	numbers := []int{42, 40, 41, 42, 38}
	var wg sync.WaitGroup

	for _, n := range numbers {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			start := time.Now()
			value, _ := cache.Get(index)
			fmt.Println("Tiempo transcurrido", time.Since(start), "n = ", n, ", f(n) = ", value)
		}(n)
	}
	wg.Wait()
}
