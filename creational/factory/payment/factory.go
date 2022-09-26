package payment

import "errors"

func NewPaymentFactory(paymentType string) (Payment, error) {
	switch paymentType {
	case "cash":
		return createCashPayment(), nil
	case "card":
		return createCardPayment(), nil
	}

	return nil, errors.New("payment type must be: cash or card")
}
