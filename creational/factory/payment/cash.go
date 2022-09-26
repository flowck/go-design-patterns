package payment

type cash struct {
	payment
}

func createCashPayment() Payment {
	return cash{payment{}}
}
