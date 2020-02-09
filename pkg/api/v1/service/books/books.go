package books

import (
	repository "example/pkg/api/v1/database/books"
	"example/pkg/api/v1/models/books"
	"example/pkg/api/v1/service"
	"log"
)

type Manager struct {
}

func (m Manager) GetBooks() ([]books.Book, error) {
	booksList, err := repository.Repository.GetBooks()
	if err != nil {
		log.Printf("GetBooks: %v", err)
		return nil, service.ErrInternal
	}

	if len(booksList) == 0 {
		return nil, service.ErrBooksNotFound
	}

	return booksList, nil
}
