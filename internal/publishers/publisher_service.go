package publishers

import (
	"log"
)

type PublisherServiceAdapter struct {
	repo PublisherRepositoryPort
}

func NewPublisherServiceAdapter(repo PublisherRepositoryPort) PublisherServicePort {
	return &PublisherServiceAdapter{
		repo: repo,
	}
}

func (p *PublisherServiceAdapter) Create(publisher *Publisher) (*Publisher, error) {
	id, err := p.repo.Create(publisher)
	if err != nil {
		return nil, err
	}
	log.Printf("service: Publisher created with ID: %d", id)
	newPublisher, err := p.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	return newPublisher, nil
}

func (p *PublisherServiceAdapter) FindByID(id int) (*Publisher, error) {
	publisher, err := p.repo.FindByID(int64(id))
	if err != nil {
		return nil, err
	}
	return publisher, nil
}
