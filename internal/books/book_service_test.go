package books_test

import (
	"errors"
	"testing"
	"time"

	"github.com/gui-laranjeira/livreria/internal/books"
	"github.com/gui-laranjeira/livreria/pkg/tests/book"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestBookService_Create(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		newBook := &books.Book{
			Title:       "The Lord of the Rings",
			ISBN:        "978-0-261-10235-4",
			PublisherID: 1,
			Pages:       1200,
			Edition:     1,
			Year:        1954,
			Language:    "English",
			Owner:       "Gui Laranjeira",
		}

		expectedBook := &books.Book{
			ID:          1,
			Title:       "The Lord of the Rings",
			ISBN:        "978-0-261-10235-4",
			PublisherID: 1,
			Pages:       1200,
			Edition:     1,
			Year:        1954,
			Language:    "English",
			Owner:       "Gui Laranjeira",
			CreatedAt:   time.Now(),
			Active:      true,
		}

		bookRepo.On("Create", mock.AnythingOfType("*books.Book")).Return(int64(1), nil).Once()
		bookRepo.On("FindByID", int64(1)).Return(expectedBook, nil).Once()

		createdBook, err := bookService.Create(newBook)

		assert.NoError(t, err)
		assert.NotNil(t, createdBook)
		assert.Equal(t, expectedBook.ID, createdBook.ID)
		bookRepo.AssertExpectations(t)
	})

	t.Run("Error on Create", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		newBook := &books.Book{Title: "Test Book"}

		bookRepo.On("Create", mock.AnythingOfType("*books.Book")).Return(int64(0), errors.New("database error")).Once()

		createdBook, err := bookService.Create(newBook)

		assert.Error(t, err)
		assert.Nil(t, createdBook)
		bookRepo.AssertExpectations(t)
	})

	t.Run("Error on FindByID after Create", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		newBook := &books.Book{Title: "Test Book"}

		bookRepo.On("Create", mock.AnythingOfType("*books.Book")).Return(int64(1), nil).Once()
		bookRepo.On("FindByID", int64(1)).Return(nil, errors.New("database error")).Once()

		createdBook, err := bookService.Create(newBook)

		assert.Error(t, err)
		assert.Nil(t, createdBook)
		bookRepo.AssertExpectations(t)
	})
}

func TestBookService_FindByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBook := &books.Book{ID: 1, Title: "Test Book"}

		bookRepo.On("FindByID", int64(1)).Return(expectedBook, nil).Once()

		foundBook, err := bookService.FindByID(1)

		assert.NoError(t, err)
		assert.NotNil(t, foundBook)
		assert.Equal(t, expectedBook.ID, foundBook.ID)
		bookRepo.AssertExpectations(t)
	})

	t.Run("Not Found", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByID", int64(1)).Return(nil, errors.New("book not found")).Once()

		foundBook, err := bookService.FindByID(1)

		assert.Error(t, err)
		assert.Nil(t, foundBook)
		bookRepo.AssertExpectations(t)
	})
}
