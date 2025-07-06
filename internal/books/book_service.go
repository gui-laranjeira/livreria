package books

import (
	"time"
)

type BookServiceAdapter struct {
	repository BookRepositoryPort
}

func NewBookServiceAdapter(repository BookRepositoryPort) BookServicePort {
	return &BookServiceAdapter{
		repository: repository,
	}
}

func (b *BookServiceAdapter) Create(book *Book) (*Book, error) {
	book.CreatedAt = time.Now()
	book.Active = true

	id, err := b.repository.Create(book)
	if err != nil {
		return nil, err
	}

	return b.repository.FindByID(id)
}

func (b *BookServiceAdapter) Update(book *Book) (*Book, error) {
	return b.repository.Update(book)
}

func (b *BookServiceAdapter) FindAll() ([]*Book, error) {
	return b.repository.FindAll()
}

func (b *BookServiceAdapter) FindByID(id int) (*Book, error) {
	return b.repository.FindByID(int64(id))
}

func (b *BookServiceAdapter) FindByTitle(title string) ([]*Book, error) {
	return b.repository.FindByTitle(title)
}

func (b *BookServiceAdapter) FindByPublisherID(publisherID int) ([]*Book, error) {
	return b.repository.FindByPublisherID(publisherID)
}

func (b *BookServiceAdapter) FindByISBN(isbn string) ([]*Book, error) {
	return b.repository.FindByISBN(isbn)
}

func (b *BookServiceAdapter) FindByOwner(owner string) ([]*Book, error) {
	return b.repository.FindByOwner(owner)
}

func (b *BookServiceAdapter) DeleteByID(id int) error {
	return b.repository.DeleteByID(id)
}
