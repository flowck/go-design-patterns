package main

import (
	"fmt"
	"go-design-patterns/patterns/creational/factory/payment"
	"go-design-patterns/patterns/creational/factory/publication"
	"log"
)

func main() {
	var err error
	var mag publication.IPublication
	var pub publication.IPublication
	var cashPayment payment.Payment
	var cardPayment payment.Payment

	if mag, err = publication.NewPublication(publication.Magazine, "Vogue", 20, "Vogue Inc"); err != nil {
		log.Fatal(err)
	}

	if pub, err = publication.NewPublication(publication.Newspaper, "NY Times", 12, "New York Times Inc"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Magazine", mag)
	fmt.Println("Newspaper", pub)

	if cashPayment, err = payment.NewPaymentFactory("cash"); err != nil {
		log.Fatal(err)
	}

	if cardPayment, err = payment.NewPaymentFactory("card"); err != nil {
		log.Fatal(err)
	}

	cashPayment.Pay(100)
	cardPayment.Pay(20)
}
