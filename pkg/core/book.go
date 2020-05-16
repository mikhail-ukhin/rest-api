package core

type Book struct {
	ID          int64
	Name        string
	Description string
}

type BookDto struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BookRepository interface {
	GetById(id int64) (*Book, error)
	Add(book *Book) error
	Update(book *Book) (Book, error)
	Remove(id int64)
}

type BookService interface {
	GetById(id int64) (BookDto, error)
	Add(book BookDto) error
	Update(book *BookDto) (BookDto, error)
	Remove(id int64)
}
