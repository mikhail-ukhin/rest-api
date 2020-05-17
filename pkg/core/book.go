package core

type Book struct {
	ID          int64
	Name        string
	Description string
}

type BookDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BookService interface {
	getById(ID int64) (BookDTO, error)
	add(book BookDTO) error
	update(book *BookDTO) (BookDTO, error)
	remove(ID int64)
}
