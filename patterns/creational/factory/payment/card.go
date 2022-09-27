package payment

type card struct {
	payment
}

func createCardPayment() Payment {
	return card{payment{}}
}

func (c card) String() {
	
}
