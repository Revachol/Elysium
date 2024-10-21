package main

import "fmt"

func AddBook(books *[]Book, book Book, publisherBookCount map[string]int) {
	publisherName := book.Publisher.Name
	publisherBookCount[publisherName]++
	*books = append(*books, book)
}

func RemoveBook(books *[]Book, isbn string, publisherBookCount map[string]int) {
	for i, book := range *books {
		if book.ISBN == isbn {
			publisherName := book.Publisher.Name
			publisherBookCount[publisherName]--
			CheckMap(publisherBookCount)
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

func findBooksByPublisher(books []Book, publisher string) []Book {
	var result []Book
	for _, book := range books {
		if book.Publisher.Name == publisher {
			result = append(result, book)
		}
	}
	return result
}

func findBooksByYear(books []Book, year int) []Book {
	var result []Book
	for _, book := range books {
		if book.Year > year {
			result = append(result, book)
		}
	}
	return result
}

func PrintAllBooks(books []Book) {
	for i := 0; i < len(books); i++ {
		fmt.Printf("Title: %s\nAuthor: %s\nYear: %d\nISBN: %s\n", books[i].Title, books[i].Author, books[i].Year, books[i].ISBN)
		fmt.Println("--------------")
	}
}

func CheckMap(publisherBookCount map[string]int) {
	for key, value := range publisherBookCount {
		if value == 0 {
			delete(publisherBookCount, key)
		}
	}
}

func BooksByPublisher(books []Book) map[string][]Book {
	publisherGroups := make(map[string][]Book)

	for _, book := range books {
		publisherName := book.Publisher.Name
		if _, exists := publisherGroups[publisherName]; exists {
			publisherGroups[publisherName] = append(publisherGroups[publisherName], book)
		} else {
			publisherGroups[publisherName] = []Book{book}
		}
	}

	return publisherGroups

}

func PrintPublisher(publisherBookCount map[string]int) {
	for key, value := range publisherBookCount {
		fmt.Printf("Publisher: %s, Count: %d\n", key, value)
	}
}

func (b Book) GetDetails() string {
	return fmt.Sprintf("Book: '%s' by %s, published in %d by %s", b.Title, b.Author, b.Year, b.Publisher)
}

func (p Publisher) GetDetails() string {
	return fmt.Sprintf("Publisher: '%s', located in %s", p.Name, p.Address)
}

func PrintLibraryDetails(items []LibraryItem) {
	for _, item := range items {
		fmt.Println(item.GetDetails())
	}
}

func SortByYear(books []Book, desc bool) []Book {
	var Sorted []Book = books
	for i := 0; i < len(Sorted)-1; i++ {
		for j := i + 1; j < len(Sorted); j++ {
			if desc {
				if Sorted[i].Year < Sorted[j].Year {
					Sorted[i], Sorted[j] = Sorted[j], Sorted[i]
				}
			} else {
				if Sorted[i].Year > Sorted[j].Year {
					Sorted[i], Sorted[j] = Sorted[j], Sorted[i]
				}
			}
		}
	}
	return Sorted
}
