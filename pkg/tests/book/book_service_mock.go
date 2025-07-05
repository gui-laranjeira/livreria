package book

import (
	"github.com/gui-laranjeira/livreria/internal/books"
	"github.com/stretchr/testify/mock"
)

type BookServiceMock struct {
	mock.Mock
}

func (m *BookServiceMock) Create(book *books.Book) (*books.Book, error) {
	args := m.Called(book)
	var r0 *books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).(*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) Update(book *books.Book) (*books.Book, error) {
	args := m.Called(book)
	var r0 *books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).(*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) FindAll() ([]*books.Book, error) {
	args := m.Called()
	var r0 []*books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).([]*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) FindByID(id int) (*books.Book, error) {
	args := m.Called(id)
	var r0 *books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).(*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) FindByTitle(title string) ([]*books.Book, error) {
	args := m.Called(title)
	var r0 []*books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).([]*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) FindByPublisherID(publisherID int) ([]*books.Book, error) {
	args := m.Called(publisherID)
	var r0 []*books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).([]*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) FindByISBN(isbn string) ([]*books.Book, error) {
	args := m.Called(isbn)
	var r0 []*books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).([]*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) FindByOwner(owner string) ([]*books.Book, error) {
	args := m.Called(owner)
	var r0 []*books.Book
	if args.Get(0) != nil {
		r0 = args.Get(0).([]*books.Book)
	}
	return r0, args.Error(1)
}

func (m *BookServiceMock) DeleteByID(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
