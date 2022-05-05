package main

type Topic[T any] interface {
	Register(observer Observer[T])
	Broadcast(value T)
}

type Observer[T any] func(value T)
