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
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindAll() ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByID(id int) (*Book, error) {
	book, err := b.repository.FindByID(int64(id))
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookServiceAdapter) FindByTitle(title string) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByPublisherID(publisherID int) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByISBN(isbn string) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByOwner(owner string) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) DeleteByID(id int) error {
	panic("TODO: implement me")
}
