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

type BookRepository interface {
	GetById(ID int64) (*Book, error)
	Add(book *Book) error
	Update(book *Book) (Book, error)
	Remove(ID int64)
}

type BookService interface {
	GetById(ID int64) (BookDTO, error)
	Add(book BookDTO) error
	Update(book *BookDTO) (BookDTO, error)
	Remove(ID int64)
}
