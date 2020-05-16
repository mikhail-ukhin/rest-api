package main

import (
	"rest-api/pkg/core"
	"rest-api/pkg/repository"
)

func main() {

	repo, err := repository.NewRepository("sqlite3", ":memory:")

	if err != nil {
		panic(err)
	}

	book := core.Book{
		Name:        "Test",
		Description: "Example",
	}

	repo.Add(&book)

}
