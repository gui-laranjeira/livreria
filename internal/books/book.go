package books

import (
	"github.com/gui-laranjeira/livreria/internal/publishers/entity"
	"time"
)

type Book struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	PublisherID int        `json:"publisher_id"`
	Pages       int        `json:"pages"`
	Language    string     `json:"language,omitempty"`
	Edition     int        `json:"edition,omitempty"`
	Year        int        `json:"year,omitempty"`
	ISBN        string     `json:"isbn,omitempty"`
	Owner       string     `json:"owner,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
	Active      bool       `json:"active"`
}

type BookRepositoryPort interface {
	Create(book *Book) (int64, error)
	Update(book *Book) (*Book, error)
	FindAll() ([]*Book, error)
	FindByID(id int64) (*Book, error)
	FindByTitle(title string) ([]*Book, error)
	FindByPublisherID(publisherID int) ([]*Book, error)
	FindByISBN(isbn string) ([]*Book, error)
	FindByOwner(owner string) ([]*Book, error)
	DeleteByID(id int) error
}

type BookServicePort interface {
	Create(book *Book) (*Book, error)
	Update(book *Book) (*Book, error)
	FindAll() ([]*Book, error)
	FindByID(id int) (*Book, error)
	FindByTitle(title string) ([]*Book, error)
	FindByPublisherID(publisherID int) ([]*Book, error)
	FindByISBN(isbn string) ([]*Book, error)
	FindByOwner(owner string) ([]*Book, error)
	DeleteByID(id int) error
}

func NewBookFactory(title string, publisher *entity.Publisher, pages int, language string, edition int, year int, isbn string,
	owner string) (*Book, error) {
	if err := ValidateBook(title, publisher, pages, language, edition, year, isbn, owner); err != nil {
		return nil, err
	}
	return &Book{
		Title:       title,
		PublisherID: publisher.ID,
		Pages:       pages,
		Language:    language,
		Edition:     edition,
		Year:        year,
		ISBN:        isbn,
		Owner:       owner,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
		DeletedAt:   nil,
		Active:      true,
	}, nil
}

func ValidateBook(title string, publisher *entity.Publisher, pages int, language string, edition int, year int,
	isbn string, owner string) error {
	if title == "" {
		return ErrTitleRequired
	}
	if publisher == nil {
		return ErrPublisherRequired
	}
	if pages <= 0 {
		return ErrPagesRequired
	}
	if language == "" {
		return ErrLanguageRequired
	}
	if edition <= 0 {
		return ErrEditionRequired
	}
	if year <= 0 {
		return ErrYearRequired
	}
	if isbn == "" {
		return ErrISBNRequired
	}
	if owner == "" {
		return ErrOwnerRequired
	}
	return nil
}
