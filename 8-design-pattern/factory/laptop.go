package main

type Laptop struct {
	Computer
}

func NewLaptop() Product {
	return &Laptop{
		Computer: Computer{
			name:  "Laptop Computer",
			stock: 25,
		},
	}
}
