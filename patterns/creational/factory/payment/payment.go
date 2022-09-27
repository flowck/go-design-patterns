package payment

type Payment interface {
	Pay(amount float64) error
}

type payment struct {
	amount float64
}

func (p payment) Pay(amount float64) error {
	return nil
}
