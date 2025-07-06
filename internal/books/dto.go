package books

type CreateBookRequest struct {
	Title     string `json:"title" binding:"required"`
	Publisher string `json:"publisher" binding:"required"`
	Pages     int    `json:"pages" binding:"required"`
	Language  string `json:"language" binding:"required"`
	Edition   int    `json:"edition" binding:"required"`
	Year      int    `json:"year" binding:"required"`
	ISBN      string `json:"isbn" binding:"required"`
	Owner     string `json:"owner" binding:"required"`
}

type UpdateBookRequest struct {
	Title     string `json:"title"`
	Publisher string `json:"publisher"`
	Pages     int    `json:"pages"`
	Language  string `json:"language"`
	Edition   int    `json:"edition"`
	Year      int    `json:"year"`
	ISBN      string `json:"isbn"`
	Owner     string `json:"owner"`
}
