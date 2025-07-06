package books

import (
	"database/sql"
	"time"
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
	stmt, err := b.db.Prepare(`INSERT INTO books (title, publisher_id, pages, language, edition, year, isbn, owner, created_at, active)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) RETURNING id`)
	if err != nil {
		return 0, err
	}

	var id int64
	err = stmt.QueryRow(book.Title, book.PublisherID, book.Pages, book.Language, book.Edition, book.Year,
		book.ISBN, book.Owner, book.CreatedAt, book.Active).Scan(&id)
	if err != nil {
		return 0, err
	}

	if id == 0 {
		return 0, sql.ErrNoRows
	}
	return id, err
}

func (b *BookRepositoryAdapter) Update(book *Book) (*Book, error) {
	stmt := `UPDATE books SET title = $1, publisher_id = $2, pages = $3, language = $4, edition = $5, year = $6, isbn = $7, owner = $8, updated_at = $9 WHERE id = $10`

	_, err := b.db.Exec(stmt, book.Title, book.PublisherID, book.Pages, book.Language, book.Edition, book.Year,
		book.ISBN, book.Owner, time.Now(), book.ID)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookRepositoryAdapter) FindAll() ([]*Book, error) {
	stmt := `SELECT id, title, publisher_id, pages, language, edition, year, isbn, owner, created_at, updated_at, deleted_at, active
		FROM books WHERE active = TRUE`

	books := []*Book{}

	rows, err := b.db.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.ID, &book.Title, &book.PublisherID, &book.Pages, &book.Language,
			&book.Edition, &book.Year, &book.ISBN, &book.Owner, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Active)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepositoryAdapter) FindByID(id int64) (*Book, error) {
	sqlStatement := `SELECT id, title, publisher_id, pages, language, edition, year, isbn, owner, created_at, updated_at, deleted_at, active
		FROM books WHERE id = $1 AND active = TRUE`

	book := new(Book)
	err := b.db.QueryRow(sqlStatement, id).Scan(&book.ID, &book.Title, &book.PublisherID, &book.Pages, &book.Language,
		&book.Edition, &book.Year, &book.ISBN, &book.Owner, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Active)
	if err != nil {
		return nil, err
	}

	return book, nil
}

func (b *BookRepositoryAdapter) FindByTitle(title string) ([]*Book, error) {
	stmt := `SELECT id, title, publisher_id, pages, language, edition, year, isbn, owner, created_at, updated_at, deleted_at, active
		FROM books WHERE title = $1 AND active = TRUE`

	books := []*Book{}

	rows, err := b.db.Query(stmt, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.ID, &book.Title, &book.PublisherID, &book.Pages, &book.Language,
			&book.Edition, &book.Year, &book.ISBN, &book.Owner, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Active)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepositoryAdapter) FindByPublisherID(publisherID int) ([]*Book, error) {
	stmt := `SELECT id, title, publisher_id, pages, language, edition, year, isbn, owner, created_at, updated_at, deleted_at, active
		FROM books WHERE publisher_id = $1 AND active = TRUE`

	books := []*Book{}

	rows, err := b.db.Query(stmt, publisherID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.ID, &book.Title, &book.PublisherID, &book.Pages, &book.Language,
			&book.Edition, &book.Year, &book.ISBN, &book.Owner, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Active)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepositoryAdapter) FindByISBN(isbn string) ([]*Book, error) {
	stmt := `SELECT id, title, publisher_id, pages, language, edition, year, isbn, owner, created_at, updated_at, deleted_at, active
		FROM books WHERE isbn = $1 AND active = TRUE`

	books := []*Book{}

	rows, err := b.db.Query(stmt, isbn)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.ID, &book.Title, &book.PublisherID, &book.Pages, &book.Language,
			&book.Edition, &book.Year, &book.ISBN, &book.Owner, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Active)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepositoryAdapter) FindByOwner(owner string) ([]*Book, error) {
	stmt := `SELECT id, title, publisher_id, pages, language, edition, year, isbn, owner, created_at, updated_at, deleted_at, active
		FROM books WHERE owner = $1 AND active = TRUE`

	books := []*Book{}

	rows, err := b.db.Query(stmt, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		book := new(Book)
		err := rows.Scan(&book.ID, &book.Title, &book.PublisherID, &book.Pages, &book.Language,
			&book.Edition, &book.Year, &book.ISBN, &book.Owner, &book.CreatedAt, &book.UpdatedAt, &book.DeletedAt, &book.Active)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}

	return books, nil
}

func (b *BookRepositoryAdapter) DeleteByID(id int) error {
	stmt := `UPDATE books SET deleted_at = $1 WHERE id = $2`

	_, err := b.db.Exec(stmt, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
