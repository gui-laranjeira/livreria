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

func TestBookService_Update(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBook := &books.Book{ID: 1, Title: "Test Book"}

		bookRepo.On("Update", mock.AnythingOfType("*books.Book")).Return(expectedBook, nil).Once()

		updatedBook, err := bookService.Update(expectedBook)

		assert.NoError(t, err)
		assert.NotNil(t, updatedBook)
		assert.Equal(t, expectedBook.ID, updatedBook.ID)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Error on Update", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBook := &books.Book{ID: 1, Title: "Test Book"}

		bookRepo.On("Update", mock.AnythingOfType("*books.Book")).Return(nil, errors.New("database error")).Once()

		updatedBook, err := bookService.Update(expectedBook)

		assert.Error(t, err)
		assert.Nil(t, updatedBook)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Not Found", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("Update", mock.AnythingOfType("*books.Book")).Return(nil, errors.New("book not found")).Once()

		updatedBook, err := bookService.Update(&books.Book{ID: 1, Title: "Test Book"})

		assert.Error(t, err)
		assert.Nil(t, updatedBook)
		bookRepo.AssertExpectations(t)
	})
}

func TestBookService_FindAll(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBooks := []*books.Book{
			{ID: 1, Title: "Test Book 1"},
			{ID: 2, Title: "Test Book 2"},
		}

		bookRepo.On("FindAll").Return(expectedBooks, nil).Once()

		foundBooks, err := bookService.FindAll()

		assert.NoError(t, err)
		assert.NotNil(t, foundBooks)
		assert.Equal(t, expectedBooks, foundBooks)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Error on FindAll", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindAll").Return(nil, errors.New("database error")).Once()

		foundBooks, err := bookService.FindAll()

		assert.Error(t, err)
		assert.Nil(t, foundBooks)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Not Found", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindAll").Return(nil, errors.New("book not found")).Once()

		foundBooks, err := bookService.FindAll()

		assert.Error(t, err)
		assert.Nil(t, foundBooks)
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

func TestBookService_FindByTitle(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBook := &books.Book{ID: 1, Title: "Test Book"}

		bookRepo.On("FindByTitle", "Test Book").Return([]*books.Book{expectedBook}, nil).Once()

		foundBooks, err := bookService.FindByTitle("Test Book")

		assert.NoError(t, err)
		assert.NotNil(t, foundBooks)
		assert.Equal(t, expectedBook.ID, foundBooks[0].ID)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Not Found", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByTitle", "Test Book").Return(nil, errors.New("book not found")).Once()

		foundBooks, err := bookService.FindByTitle("Test Book")

		assert.Error(t, err)
		assert.Nil(t, foundBooks)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Error on FindByTitle", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByTitle", "Test Book").Return(nil, errors.New("database error")).Once()

		foundBooks, err := bookService.FindByTitle("Test Book")

		assert.Error(t, err)
		assert.Nil(t, foundBooks)
		bookRepo.AssertExpectations(t)
	})
}

func TestBookService_FindByPublisherID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBook := &books.Book{ID: 1, Title: "Test Book"}

		bookRepo.On("FindByPublisherID", 1).Return([]*books.Book{expectedBook}, nil).Once()

		foundBooks, err := bookService.FindByPublisherID(1)

		assert.NoError(t, err)
		assert.NotNil(t, foundBooks)
		assert.Equal(t, expectedBook.ID, foundBooks[0].ID)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Not Found", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByPublisherID", 1).Return(nil, errors.New("book not found")).Once()

		foundBooks, err := bookService.FindByPublisherID(1)

		assert.Error(t, err)
		assert.Nil(t, foundBooks)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Error on FindByPublisherID", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByPublisherID", 1).Return(nil, errors.New("database error")).Once()

		foundBooks, err := bookService.FindByPublisherID(1)

		assert.Error(t, err)
		assert.Nil(t, foundBooks)
		bookRepo.AssertExpectations(t)
	})
}

func TestBookService_FindByISBN(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBook := &books.Book{ID: 1, Title: "Test Book"}

		bookRepo.On("FindByISBN", "1234567890").Return([]*books.Book{expectedBook}, nil).Once()

		foundBook, err := bookService.FindByISBN("1234567890")

		assert.NoError(t, err)
		assert.NotNil(t, foundBook)
		assert.Equal(t, expectedBook.ID, foundBook[0].ID)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Not Found", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByISBN", "1234567890").Return(nil, errors.New("book not found")).Once()

		foundBook, err := bookService.FindByISBN("1234567890")

		assert.Error(t, err)
		assert.Nil(t, foundBook)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Error on FindByISBN", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByISBN", "1234567890").Return(nil, errors.New("database error")).Once()

		foundBook, err := bookService.FindByISBN("1234567890")

		assert.Error(t, err)
		assert.Nil(t, foundBook)
		bookRepo.AssertExpectations(t)
	})
}

func TestBookService_FindByOwner(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		expectedBook := &books.Book{ID: 1, Title: "Test Book"}

		bookRepo.On("FindByOwner", "Gui Laranjeira").Return([]*books.Book{expectedBook}, nil).Once()

		foundBook, err := bookService.FindByOwner("Gui Laranjeira")

		assert.NoError(t, err)
		assert.NotNil(t, foundBook)
		assert.Equal(t, expectedBook.ID, foundBook[0].ID)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Not Found", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByOwner", "Gui Laranjeira").Return(nil, errors.New("book not found")).Once()

		foundBook, err := bookService.FindByOwner("Gui Laranjeira")

		assert.Error(t, err)
		assert.Nil(t, foundBook)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Error on FindByOwner", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("FindByOwner", "Gui Laranjeira").Return(nil, errors.New("database error")).Once()

		foundBook, err := bookService.FindByOwner("Gui Laranjeira")

		assert.Error(t, err)
		assert.Nil(t, foundBook)
		bookRepo.AssertExpectations(t)
	})
}

func TestBookService_DeleteByID(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("DeleteByID", 1).Return(nil).Once()

		err := bookService.DeleteByID(1)

		assert.NoError(t, err)
		bookRepo.AssertExpectations(t)
	})
	t.Run("Error on DeleteByID", func(t *testing.T) {
		bookRepo := &book.BookRepositoryMock{}
		bookService := books.NewBookServiceAdapter(bookRepo)

		bookRepo.On("DeleteByID", 1).Return(errors.New("database error")).Once()

		err := bookService.DeleteByID(1)

		assert.Error(t, err)
		bookRepo.AssertExpectations(t)
	})
}
