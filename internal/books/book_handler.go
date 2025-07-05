package books

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gui-laranjeira/livreria/internal/publisher"
	"github.com/gui-laranjeira/livreria/pkg/web"
)

type BookHandlerAdapter struct {
	bookService      BookServicePort
	publisherService publisher.PublisherServicePort
}

func NewBookHandlerAdapter(bookService BookServicePort, publisherService publisher.PublisherServicePort) *BookHandlerAdapter {
	return &BookHandlerAdapter{
		bookService:      bookService,
		publisherService: publisherService,
	}
}

func (h *BookHandlerAdapter) Create(c *gin.Context) {
	var input CreateBookRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		web.Error(c, http.StatusBadRequest, "error binding json: %v", err)
		return
	}
	if len(input.Language) > 2 {
		web.Error(c, http.StatusBadRequest, "error validating language: %v", "language follow the ISO 639-1 standard")
		return
	}

	p, err := h.publisherService.FindByName(input.Publisher)
	if p == nil && errors.Is(err, publisher.ErrPublisherNotFound) {
		p, err = h.publisherService.Create(p)
		if err != nil {
			web.Error(c, http.StatusInternalServerError, "error creating publisher: %v", err)
			return
		}
	}

	book, err := NewBookFactory(input.Title, p, input.Pages, input.Language, input.Edition, input.Year, input.ISBN, input.Owner)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error creating book: %v", err)
		return
	}
	b, err := h.bookService.Create(book)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error creating book: %v", err)
		return
	}
	web.Success(c, http.StatusCreated, b)
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
