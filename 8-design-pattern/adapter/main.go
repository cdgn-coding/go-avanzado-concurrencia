package main

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)

	bank := BankPaymentAdapter{
		BankPayment: BankPayment{},
		account:     10,
	}
	ProcessPayment(bank)
}
