package storage

import "database/sql"

var (
	dbType = "sqlite"
	dbName = "go-book.db"
)

type storage struct {
	db sql.DB
}

type Storage interface {
	NewStorage() (Storage, error)
}

func (s *storage) NewStorage() (Storage, error) {
	client, err := NewClient()

	if err != nil {
		return nil, err
	}

	return &storage{*client}, nil
}

func NewClient() (*sql.DB, error) {
	db, err := sql.Open(dbType, dbName)

	if err != nil {
		return nil, err
	}

	return db, nil
}
