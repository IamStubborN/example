package service

import "errors"

var (
	ErrInternal = errors.New("internal server error")
	ErrBooksNotFound = errors.New("books not founds")
)
