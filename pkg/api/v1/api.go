package api

import (
	"example/pkg/api/v1/handlers/books"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	address string
	books   *books.Handler
	router  *gin.Engine
}

func NewServer(books *books.Handler) *Server {
	return &Server{
		books:   books,
		router:  gin.New(),
		address: ":8080",
	}
}

func (s Server) Routes() {
	s.router.GET("api/v1/books", s.books.GetBooks)

}

func (s Server) Run() {
	if err := s.router.Run(s.address); err != nil {
		log.Fatal(err)
	}
}
