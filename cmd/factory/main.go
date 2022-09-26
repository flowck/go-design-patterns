package main

import (
	"fmt"
	"go-design-patterns/creational/factory"
	"log"
)

func main() {
	var err error
	var mag factory.IPublication
	var pub factory.IPublication

	if mag, err = factory.NewPublication(factory.PublicationMagazine, "Vogue", 20, "Vogue Inc"); err != nil {
		log.Fatal(err)
	}

	if pub, err = factory.NewPublication(factory.PublicationNewspaper, "NY Times", 12, "New York Times Inc"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Magazine", mag)
	fmt.Println("Newspaper", pub)
}
