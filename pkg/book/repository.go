package book

import (
	"REST-api/pkg/storage"
)

type Repository interface {
}

type repository struct {
	conf storage.configuration
}

func (r *repository) NewRepository(storage storage.Storage) Repository {
	return &repository{storage}
}

func (r *repository) Get(id int32) (Book, error) {
	return Book{}, nil
}
