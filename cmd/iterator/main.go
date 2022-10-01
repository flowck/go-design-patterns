package main

import (
	"fmt"
	"go-design-patterns/patterns/behavioral/iterator"
)

func main() {
	lib := &iterator.Library{Collection: []iterator.Book{
		{
			Name:      "War and Peace",
			Author:    "Leo",
			PageCount: 684,
			Type:      iterator.HardCover,
		},
		{
			Name:      "Crime and Punishment",
			Author:    "Leo",
			PageCount: 1225,
			Type:      iterator.SoftCover,
		},
		{
			Name:      "Brave New World",
			Author:    "Aldous Huxley",
			PageCount: 325,
			Type:      iterator.EBook,
		},
	}}

	iter := lib.CreateIterator()

	for iter.HasNext() {
		fmt.Println("Book's name from .next():  ", iter.Next().Name)
	}

	// count := 0
	lib.IterateBooks(func(book iterator.Book) error {
		fmt.Println("Book's name from callback: ", book.Name)
		return nil
	})
}
