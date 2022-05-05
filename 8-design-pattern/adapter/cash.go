package main

import (
	"fmt"
	"time"
)

type CashPayment struct{}

func (c CashPayment) Pay() {
	fmt.Println("Payment using Cash")
	time.Sleep(2 * time.Second)
}
