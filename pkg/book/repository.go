package book

import (
	"REST-api/pkg/storage"
	"database/sql"
)

type Repository interface {
}

type repository struct {
	s storage.Storage
}

func (r *repository) NewRepository(storage storage.Storage) Repository {
	return &repository{storage}
}

func (r *repository) Get(id int32) (Book, error) {

}
