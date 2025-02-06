package repositories

import (
"library-api/models"
)

var Books []models.Book 
var NextID = 1 

func ResetBooks() {
	Books = []models.Book{}
    NextID = 1  
}

func AddBook(book models.Book) models.Book {
	book.ID = NextID
	NextID++
	Books = append(Books, book)
	return book
}

func GetAllBooks() []models.Book {
	return Books
}

func GetBookByID(id int) *models.Book {
	for _, book := range Books {
		if book.ID == id {
			return &book
		}
	}
	return nil
}

func DeleteBookByID(id int) bool {
	for i, book := range Books {
		if book.ID == id {
			Books = append(Books[:i], Books[i+1:]...)
			return true
		}
	}
	return false
}

func SearchBooksByAuthor(author string) []models.Book {
	var results []models.Book
	for _, book := range Books {
		if book.Author == author {
			results = append(results, book)
		}
	}
	return results
}
