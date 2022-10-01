package iterator

type IterableCollection interface {
	CreateIterator() iterator
}

type iterator interface {
	HasNext() bool
	Next() *Book
}

type BookIterator struct {
	current int
	books   []Book
}

func (b *BookIterator) HasNext() bool {
	if b.current < len(b.books) {
		return true
	}

	return false
}

func (b *BookIterator) Next() *Book {
	if b.HasNext() {
		bk := b.books[b.current]
		b.current++
		return &bk
	}

	return nil
}
