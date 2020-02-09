package books

import (
	"database/sql"
	"example/pkg/api/v1/models/books"
)

var Repository rep

type rep struct {
	db *sql.DB
}

func InitRepository(db *sql.DB) {
	Repository = rep{
		db:db,
	}
}

func (r rep) GetBooks() ([]books.Book, error) {
	query := `select id, title from books`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	var result []books.Book
	for rows.Next() {
		var b books.Book
		if err := rows.Scan(&b.ID, &b.Title); err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		result = append(result, b)
	}

	return result, nil
}