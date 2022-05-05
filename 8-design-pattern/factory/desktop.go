package main

type Desktop struct {
	Computer
}

func NewDesktop() Product {
	return &Desktop{
		Computer: Computer{
			name:  "Desktop Computer",
			stock: 35,
		},
	}
}
