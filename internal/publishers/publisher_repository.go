package publishers

import (
	"database/sql"
	"errors"
)

type PublisherRepositoryAdapter struct {
	db *sql.DB
}

func NewPublisherRepositoryAdapter(db *sql.DB) PublisherRepositoryPort {
	return &PublisherRepositoryAdapter{
		db: db,
	}
}

func (p *PublisherRepositoryAdapter) Create(publisher *Publisher) (int64, error) {
	stmt := "INSERT INTO publishers (name, country) VALUES ($1, $2) RETURNING id"
	var id int64

	err := p.db.QueryRow(stmt, publisher.Name, publisher.Country).Scan(&id)
	if err != nil {
		return 0, err
	}
	if id == 0 {
		return 0, errors.New("failed to create publisher")
	}
	return id, nil
}

func (p *PublisherRepositoryAdapter) FindByID(id int64) (*Publisher, error) {
	stmt, err := p.db.Prepare("SELECT id, name, country FROM publishers WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var publisher Publisher
	err = stmt.QueryRow(id).Scan(&publisher.ID, &publisher.Name, &publisher.Country)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrPublisherNotFound
		}
		return nil, err
	}

	return &publisher, nil
}
