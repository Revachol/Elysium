package main

type Book struct {
	Title  string
	Author string
	Year   int
	ISBN   string
}

func NewBook(title, author string, year int, isbn string) Book {
	return Book{
		Title:  title,
		Author: author,
		Year:   year,
		ISBN:   isbn,
	}
}
