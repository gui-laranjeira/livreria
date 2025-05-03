package service

import (
	"github.com/gui-laranjeira/livreria/internal/books/entity"
	"time"
)

type BookServiceAdapter struct {
	repository entity.BookRepositoryPort
}

func NewBookServiceAdapter(repository entity.BookRepositoryPort) entity.BookServicePort {
	return &BookServiceAdapter{
		repository: repository,
	}
}

func (b *BookServiceAdapter) Create(book *entity.Book) (*entity.Book, error) {
	book.CreatedAt = time.Now()
	book.Active = true

	id, err := b.repository.Create(book)
	if err != nil {
		return nil, err
	}

	return b.repository.FindByID(id)
}

func (b *BookServiceAdapter) Update(book *entity.Book) (*entity.Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindAll() ([]*entity.Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByID(id int) (*entity.Book, error) {
	book, err := b.repository.FindByID(int64(id))
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookServiceAdapter) FindByTitle(title string) ([]*entity.Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByPublisherID(publisherID int) ([]*entity.Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByISBN(isbn string) ([]*entity.Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) FindByOwner(owner string) ([]*entity.Book, error) {
	panic("TODO: implement me")
}

func (b *BookServiceAdapter) DeleteByID(id int) error {
	panic("TODO: implement me")
}
