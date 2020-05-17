package main

import (
	"fmt"
	"rest-api/pkg/core"
	"rest-api/pkg/repository"
)

func main() {
	var ctx repository.DbContext
	var bookRepo repository.BookRepository

	ctx = repository.NewContext("sqlite3", ":memory:")
	bookRepo = repository.NewBookRepository(ctx)

	defer func() {
		err := ctx.Close()

		if err != nil {
			panic("error on close connection")
		}
	}()

	book := core.Book{
		Name:        "Test",
		Description: "Example",
	}

	bookRepo.Add(&book)

	books, _ := bookRepo.GetAll()

	fmt.Print(books)

}
