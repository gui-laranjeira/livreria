package book

import (
	"github.com/gui-laranjeira/livreria/internal/books"
	"github.com/stretchr/testify/mock"
)

type BookRepositoryMock struct {
	mock.Mock
}

func (m *BookRepositoryMock) Update(book *books.Book) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BookRepositoryMock) FindAll() ([]*books.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BookRepositoryMock) FindByID(id int64) (*books.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BookRepositoryMock) FindByTitle(title string) ([]*books.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BookRepositoryMock) FindByPublisherID(publisherID int) ([]*books.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BookRepositoryMock) FindByISBN(isbn string) ([]*books.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BookRepositoryMock) FindByOwner(owner string) ([]*books.Book, error) {
	//TODO implement me
	panic("implement me")
}

func (m *BookRepositoryMock) DeleteByID(id int) error {
	//TODO implement me
	panic("implement me")
}

func NewBookRepositoryMock() books.BookRepositoryPort {
	return &BookRepositoryMock{}
}

func (m *BookRepositoryMock) Create(book *books.Book) (int64, error) {