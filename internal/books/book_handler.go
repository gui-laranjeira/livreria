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

func (h *BookHandlerAdapter) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", err)
		return
	}
	if id < 1 {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", "id must be greater than 0")
		return
	}
	p, err := h.publisherService.FindByID(id)
	if p == nil && errors.Is(err, publisher.ErrPublisherNotFound) {
		web.Error(c, http.StatusNotFound, "publisher not found: %v", err)
		return
	}
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error finding publisher: %v", err)
		return
	}
	var input UpdateBookRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		web.Error(c, http.StatusBadRequest, "error binding json: %v", err)
		return
	}
	if len(input.Language) > 2 {
		web.Error(c, http.StatusBadRequest, "error validating language: %v", "language follow the ISO 639-1 standard")
		return
	}
	book, err := NewBookFactory(input.Title, p, input.Pages, input.Language, input.Edition, input.Year, input.ISBN, input.Owner)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error creating book: %v", err)
		return
	}
	b, err := h.bookService.Update(book)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error updating book: %v", err)
		return
	}
	web.Success(c, http.StatusOK, b)
}

func (h *BookHandlerAdapter) FindAll(c *gin.Context) {
	books, err := h.bookService.FindAll()
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error finding books: %v", err)
		return
	}
	web.Success(c, http.StatusOK, books)
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

func (h *BookHandlerAdapter) FindByTitle(c *gin.Context) {
	title := c.Param("title")
	if title == "" {
		web.Error(c, http.StatusBadRequest, "invalid title provided: %v", "title must not be empty")
		return
	}

	books, err := h.bookService.FindByTitle(title)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error finding books: %v", err)
		return
	}
	web.Success(c, http.StatusOK, books)
}

func (h *BookHandlerAdapter) FindByPublisherID(c *gin.Context) {
	publisherID, err := strconv.Atoi(c.Param("publisher_id"))
	if err != nil {
		web.Error(c, http.StatusBadRequest, "invalid publisher id provided: %v", err)
		return
	}
	if publisherID < 1 {
		web.Error(c, http.StatusBadRequest, "invalid publisher id provided: %v", "publisher id must be greater than 0")
		return
	}

	books, err := h.bookService.FindByPublisherID(publisherID)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error finding books: %v", err)
		return
	}
	web.Success(c, http.StatusOK, books)
}

func (h *BookHandlerAdapter) FindByISBN(c *gin.Context) {
	isbn := c.Param("isbn")
	if isbn == "" {
		web.Error(c, http.StatusBadRequest, "invalid isbn provided: %v", "isbn must not be empty")
		return
	}

	books, err := h.bookService.FindByISBN(isbn)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error finding books: %v", err)
		return
	}
	web.Success(c, http.StatusOK, books)
}

func (h *BookHandlerAdapter) FindByOwner(c *gin.Context) {
	owner := c.Param("owner")
	if owner == "" {
		web.Error(c, http.StatusBadRequest, "invalid owner provided: %v", "owner must not be empty")
		return
	}

	books, err := h.bookService.FindByOwner(owner)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error finding books: %v", err)
		return
	}
	web.Success(c, http.StatusOK, books)
}

func (h *BookHandlerAdapter) DeleteByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", err)
		return
	}
	if id < 1 {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", "id must be greater than 0")
		return
	}

	err = h.bookService.DeleteByID(id)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error deleting book: %v", err)
		return
	}
	web.Success(c, http.StatusOK, nil)
}
