package main

import (
	"fmt"
	"time"
)

type BankPayment struct{}

func (c BankPayment) Pay(account int) {
	fmt.Println("Payment using Bank account =", account)
	time.Sleep(2 * time.Second)
}

type BankPaymentAdapter struct {
	BankPayment
	account int
}

func (b BankPaymentAdapter) Pay() {
	b.BankPayment.Pay(b.account)
}
