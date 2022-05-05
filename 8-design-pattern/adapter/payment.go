package main

type Payment interface {
	Pay()
}

func ProcessPayment(p Payment) {
	p.Pay()
}
