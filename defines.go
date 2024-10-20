package main

type Publisher struct {
	Name    string
	Address string
}

type Book struct {
	Title     string
	Author    string
	Year      int
	ISBN      string
	Publisher Publisher
}

func NewPublisher(name, address string) Publisher {
	return Publisher{
		Name:    name,
		Address: address,
	}
}

func NewBook(title, author string, year int, isbn, publisherName, publisherAddress string) Book {
	publisher := NewPublisher(publisherName, publisherAddress)

	return Book{
		Title:     title,
		Author:    author,
		Year:      year,
		ISBN:      isbn,
		Publisher: publisher,
	}
}

type LibraryItem interface {
	GetDetails() string
}
