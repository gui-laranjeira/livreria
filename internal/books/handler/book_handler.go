package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gui-laranjeira/livreria/internal/books/entity"
	"github.com/gui-laranjeira/livreria/pkg/web"
	"net/http"
	"strconv"
)

type BookHandlerAdapter struct {
	bookService entity.BookServicePort
}

func NewBookHandlerAdapter(bookService entity.BookServicePort) *BookHandlerAdapter {
	return &BookHandlerAdapter{
		bookService: bookService,
	}
}

func (h *BookHandlerAdapter) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", err)
		return
	}
	if id < 1 {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", "id must be greater than 0")
		return
	}

	b, err := h.bookService.FindByID(id)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error finding book: %v", err)
		return
	}
	web.Success(c, http.StatusOK, b)
}
