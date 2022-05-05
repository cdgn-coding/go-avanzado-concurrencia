package main

import (
	"fmt"
	"sync"
	"time"
)

type Database struct{}

var db *Database

type SafeCounter struct {
	count int
	mu    sync.Mutex
}

func (s *SafeCounter) Lock() {
	s.mu.Lock()
}

func (s *SafeCounter) Unlock() {
	s.mu.Unlock()
}

func (s *SafeCounter) Add() {
	s.count += 1
}

func (s *SafeCounter) Current() int {
	return s.count
}

var safeCounter = SafeCounter{}

func connectDatabase() *Database {
	fmt.Println("Connecting...")
	time.Sleep(2 * time.Second)
	fmt.Println("Connection done.")
	return &Database{}
}

func GetDatabaseInstance() *Database {
	safeCounter.Lock()
	defer safeCounter.Unlock()

	if safeCounter.Current() < 1 {
		db = connectDatabase()
		safeCounter.Add()
		return db
	}

	return db
}
