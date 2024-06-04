package book

type Book struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Publisher string  `json:"publisher"`
	Rating    float32 `json:"rating"`
}

type CreateBookInput struct {
	Title     string  `json:"title"`
	Year      int     `json:"year"`
	Publisher string  `json:"publisher"`
	Rating    float32 `json:"rating"`
}
