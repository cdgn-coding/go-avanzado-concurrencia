package main

type Product interface {
	setStock(stock int)
	getStock() int
	setName(name string)
	getName() string
}
