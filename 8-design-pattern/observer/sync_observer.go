package main

type SyncTopic[T any] struct {
	observers []Observer[T]
}

func (s *SyncTopic[T]) Register(observer Observer[T]) {
	s.observers = append(s.observers, observer)
}

func (s *SyncTopic[T]) Broadcast(value T) {
	for _, observer := range s.observers {
		observer(value)
	}
}
