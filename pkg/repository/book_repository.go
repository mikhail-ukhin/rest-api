package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"rest-api/pkg/core"
)

type BookRepository interface {
	GetById(ID int64) (*core.Book, error)
	GetAll() (*[]core.Book, error)
	Add(book *core.Book) error
	Update(book *core.Book) (*core.Book, error)
	Remove(ID int64)
}

type bookRepository struct {
	db gorm.DB
}

func NewBookRepository(context DbContext) BookRepository {
	db := context.GetDb()

	return &bookRepository{db: *db}
}

func (r *bookRepository) GetById(id int64) (*core.Book, error) {
	var book core.Book

	r.db.Find(&book, id)

	return &book, nil
}

func (r *bookRepository) GetAll() (*[]core.Book, error) {

	var books []core.Book
	r.db.Find(&books)

	return &books, nil
}

func (r *bookRepository) Add(book *core.Book) error {

	res := r.db.Create(book)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *bookRepository) Update(book *core.Book) (*core.Book, error) {
	panic("implement me")
}

func (r *bookRepository) Remove(id int64) {
	panic("implement me")
}
