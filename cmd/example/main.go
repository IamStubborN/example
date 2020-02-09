package main

import (
	"database/sql"
	"example/pkg/api/v1"
	booksRepository "example/pkg/api/v1/database/books"
	booksHandler "example/pkg/api/v1/handlers/books"
)

func main() {
	//cfg := config.LoadConfig("/config/config.yaml")

	db, _ := sql.Open("postgres", "")
	booksRepository.InitRepository(db)

	h := booksHandler.NewBooksHandler()

	server := api.NewServer(h)
	server.Run()
}
