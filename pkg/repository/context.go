package repository

import (
	"github.com/jinzhu/gorm"
	"rest-api/pkg/core"
)

type context struct {
	d gorm.DB
}

type DbContext interface {
	GetDb() *gorm.DB
	Close() error
}

func NewContext(dbType string, dbPath string) DbContext {
	db, err := gorm.Open(dbType, dbPath)

	db.AutoMigrate(&core.Book{})

	if err != nil {
		panic(err)
	}

	return &context{d: *db}
}

func (c *context) GetDb() *gorm.DB {
	return &c.d
}

func (c *context) Close() error {
	err := c.d.Close()

	return err
}
