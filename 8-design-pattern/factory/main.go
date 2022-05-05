package main

import "fmt"

func main() {
	laptop, _ := GetComputerFactory("laptop")
	desktop, _ := GetComputerFactory("desktop")
	products := []Product{laptop, desktop}
	for _, product := range products {
		fmt.Println(product)
	}
}
