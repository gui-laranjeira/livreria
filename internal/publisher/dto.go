package publisher

type CreatePublisherRequest struct {
	Name    string `json:"name" validate:"required"`
	Country string `json:"country"`
}
