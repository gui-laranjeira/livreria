package publisher

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gui-laranjeira/livreria/pkg/web"
	"net/http"
	"strconv"
)

type PublisherHandlerAdapter struct {
	service PublisherServicePort
}

func NewPublisherHandlerAdapter(service PublisherServicePort) PublisherHandlerAdapter {
	return PublisherHandlerAdapter{
		service: service,
	}
}

func (h *PublisherHandlerAdapter) FindByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", err)
		return
	}
	if id < 1 {
		web.Error(c, http.StatusBadRequest, "invalid id provided: %v", "id must be greater than 0")
		return
	}

	p, err := h.service.FindByID(id)
	if err != nil {
		if errors.Is(err, ErrPublisherNotFound) {
			web.Error(c, http.StatusNotFound, "publisher not found: %v", err)
			return
		}
		web.Error(c, http.StatusInternalServerError, "error finding publisher: %v", err)
		return
	}
	web.Success(c, http.StatusOK, p)
}

func (h *PublisherHandlerAdapter) Create(c *gin.Context) {
	var p CreatePublisherRequest
	if err := c.ShouldBindJSON(&p); err != nil {
		web.Error(c, http.StatusBadRequest, "invalid request body: %v", err)
		return
	}

	publisher, err := NewPublisherFactory(p.Name, p.Country)
	if err != nil {
		web.Error(c, http.StatusBadRequest, "invalid publisher data: %v", err)
		return
	}

	newPublisher, err := h.service.Create(publisher)
	if err != nil {
		web.Error(c, http.StatusInternalServerError, "error creating publisher: %v", err)
		return
	}
	web.Success(c, http.StatusCreated, newPublisher)
}

func (h *PublisherHandlerAdapter) FindByName(c *gin.Context) {
	name := c.Param("name")
	if name == "" {
		web.Error(c, http.StatusBadRequest, "invalid name provided: %v", "name must not be empty")
		return
	}

	p, err := h.service.FindByName(name)
	if err != nil {
		if errors.Is(err, ErrPublisherNotFound) {
			web.Error(c, http.StatusNotFound, "publisher not found: %v", err)
			return
		}
		web.Error(c, http.StatusInternalServerError, "error finding publisher: %v", err)
		return
	}
	web.Success(c, http.StatusOK, p)
}
