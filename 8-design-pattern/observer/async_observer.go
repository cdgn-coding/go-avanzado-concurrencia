package main

import "sync"

type AsyncTopic[T any] struct {
	observers []Observer[T]
	values    chan T
}

func NewAsyncTopic[T any]() AsyncTopic[T] {
	return AsyncTopic[T]{
		observers: make([]Observer[T], 0),
		values:    make(chan T, 1),
	}
}

func (s *AsyncTopic[T]) Close() {
	close(s.values)
}

func (s *AsyncTopic[T]) Register(observer Observer[T]) {
	s.observers = append(s.observers, observer)
}

func (s *AsyncTopic[T]) Broadcast(value T) {
	s.values <- value
}

func (s *AsyncTopic[T]) BroadcastWorker(wg *sync.WaitGroup) {
	defer wg.Done()
	for value := range s.values {
		for _, observer := range s.observers {
			observer(value)
		}
	}
}
