package books

import (
	"example/pkg/api/v1/service"
	booksService "example/pkg/api/v1/service/books"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewBooksHandler() *Handler {
	return &Handler{}
}

//books - GET
//books/{id} - GET
//books - POST
//books/{id} - PUT
//books/{id} - DELETE

func (h Handler) GetBooks(c *gin.Context) {
	var manager booksService.Manager
	booksList, err := manager.GetBooks()
	if err != nil {
		Reply(c, err)
		return
	}

	c.JSON(http.StatusOK, booksList)
}

func (h Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {

}

func Reply(c *gin.Context, err error) {
	switch err {
	case service.ErrInternal:
		c.Status(http.StatusInternalServerError)
		return
	case service.ErrBooksNotFound:
		c.Status(http.StatusNotFound)
		return
	default:
		c.Status(http.StatusInternalServerError)
		return
	}
}
