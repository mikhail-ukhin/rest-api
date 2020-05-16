package repository

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"rest-api/pkg/core"
)

type bookRepository struct {
	cfg configuration
}

func (r *bookRepository) GetById(id int64) (*core.Book, error) {
	db := r.openDbConnection()

	var book core.Book

	db.Find(&book, id)

	return &book, nil
}

func (r *bookRepository) Add(book *core.Book) error {
	db := r.openDbConnection()

	res := db.Create(book)

	if res.Error != nil {
		return res.Error
	}

	return nil
}

func (r *bookRepository) Update(book *core.Book) (core.Book, error) {
	panic("implement me")
}

func (r *bookRepository) Remove(id int64) {
	panic("implement me")
}

func NewRepository(dbType string, dbPath string) (core.BookRepository, error) {
	cfg := configuration{dbPath: dbPath, dbType: dbType}

	db, err := gorm.Open(dbType, dbPath)

	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.AutoMigrate(&core.Book{})

	return &bookRepository{cfg: cfg}, nil
}

func (r *bookRepository) openDbConnection() *gorm.DB {
	db, err := gorm.Open(r.cfg.dbType, r.cfg.dbPath)

	if err != nil {
		panic("failed to connect database")
	}

	//defer db.Close()

	return db
}
