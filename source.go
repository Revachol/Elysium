package main

import "fmt"

func AddBook(books *[]Book, book Book) {
	*books = append(*books, book)
}

func RemoveBook(books *[]Book, isbn string) {
	for i, book := range *books {
		if book.ISBN == isbn {
			*books = append((*books)[:i], (*books)[i+1:]...)
			return
		}
	}
}

func UpdateBook(books *[]Book, isbn string, updatedBook Book) {
	for i, book := range *books {
		if book.ISBN == isbn {
			(*books)[i] = updatedBook
			return
		}
	}
}

func FindBookByTitle(books []Book, title string) *Book {
	for _, book := range books {
		if book.Title == title {
			return &book
		}
	}
	return nil
}

func FindBooksByAuthor(books []Book, author string) []Book {
	var authorBooks []Book
	for _, book := range books {
		if book.Author == author {
			authorBooks = append(authorBooks, book)
		}
	}
	return authorBooks
}

func PrintAllBooks(books []Book) {
	for i := 0; i < len(books); i++ {
		fmt.Printf("Title: %s\nAuthor: %s\nYear: %d\nISBN: %s\n", books[i].Title, books[i].Author, books[i].Year, books[i].ISBN)
		fmt.Println("--------------")
	}
}
