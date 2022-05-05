package main

type Computer struct {
	name  string
	stock int
}

func (c *Computer) setName(name string) {
	c.name = name
}

func (c Computer) getName() string {
	return c.name
}

func (c *Computer) setStock(stock int) {
	c.stock = stock
}

func (c Computer) getStock() int {
	return c.stock
}
