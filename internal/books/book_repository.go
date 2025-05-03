package books

import (
	"database/sql"
)

type BookRepositoryAdapter struct {
	db *sql.DB
}

func NewBookRepositoryAdapter(db *sql.DB) BookRepositoryPort {
	return &BookRepositoryAdapter{
		db: db,
	}
}

func (b *BookRepositoryAdapter) Create(book *Book) (int64, error) {
	sqlStatement := `INSERT INTO books (title, publisher_id, pages, language, edition, year, isbn, owner, created_at, active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`

	result, err := b.db.Exec(sqlStatement, book.Title, book.PublisherID, book.Pages, book.Language, book.Edition, book.Year,
		book.ISBN, book.Owner, book.CreatedAt, book.Active)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	return id, err
}

func (b *BookRepositoryAdapter) Update(book *Book) (*Book, error) {
	panic("TODO: implement me")
}

func (b *BookRepositoryAdapter) FindAll() ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookRepositoryAdapter) FindByID(id int64) (*Book, error) {
	sqlStatement := `SELECT id, title, publisher_id, pages, language, edition, year, isbn, owner, created_at, updated_at, deleted_at, active
		FROM books WHERE id = $1 AND deleted_at IS NULL`

	book := new(Book)
	err := b.db.QueryRow(sqlStatement, id).Scan(&book.ID, &book.Title, &book.PublisherID, &book.Pages, &book.Language,
		&book.Edition, &book.Year, &book.ISBN, &book.Owner, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Active)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookRepositoryAdapter) FindByTitle(title string) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookRepositoryAdapter) FindByPublisherID(publisherID int) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookRepositoryAdapter) FindByISBN(isbn string) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookRepositoryAdapter) FindByOwner(owner string) ([]*Book, error) {
	panic("TODO: implement me")
}

func (b *BookRepositoryAdapter) DeleteByID(id int) error {
	panic("TODO: implement me")
}
