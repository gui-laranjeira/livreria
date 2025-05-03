package publishers

type Publisher struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Country string `json:"country"`
}

type PublisherRepositoryPort interface {
	Create(publisher *Publisher) (int64, error)
	FindByID(id int64) (*Publisher, error)
}

type PublisherServicePort interface {
	Create(publisher *Publisher) (*Publisher, error)
	FindByID(id int) (*Publisher, error)
}

func NewPublisherFactory(name string, country string) (*Publisher, error) {
	if name == "" {
		return nil, ErrPublisherNameRequired
	}
	return &Publisher{
		Name:    name,
		Country: country,
	}, nil
}
