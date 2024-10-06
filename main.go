package main

import "fmt"

func main() {
	var books []Book

	book1 := NewBook("The Go Programming Language", "Alan A. A. Donovan", 2015, "9780134190440")
	book2 := NewBook("Clean Code", "Robert C. Martin", 2008, "9780132350884")
	book3 := NewBook("Introduction to Algorithms", "Thomas H. Cormen", 2009, "9780262033848")

	AddBook(&books, book1)
	AddBook(&books, book2)
	AddBook(&books, book3)

	fmt.Println("All books:")
	PrintAllBooks(books)
	fmt.Printf("\n============================\n\n")
	fmt.Println("Books by Robert C. Martin:")
	authorBooks := FindBooksByAuthor(books, "Robert C. Martin")
	PrintAllBooks(authorBooks)
	fmt.Printf("\n============================\n\n")

	fmt.Println("Update 'Clean Code' year:")
	updatedBook := NewBook("Clean Code", "Robert C. Martin", 2010, "9780132350884")
	UpdateBook(&books, "9780132350884", updatedBook)
	PrintAllBooks(books)
	fmt.Printf("\n============================\n\n")
	fmt.Println("Remove 'Clean Code':")
	RemoveBook(&books, "9780132350884")
	PrintAllBooks(books)
}
