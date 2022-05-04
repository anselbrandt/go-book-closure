package main

import (
	"database/sql"
	"log"
	"net/http"

	"go-books/handlers"
	"go-books/handlers/books"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./bookstore.db")
	if err != nil {
		log.Fatal(err)
	}

	env := &handlers.Env{DB: db}

	http.Handle("/books", books.Index(env))
	http.ListenAndServe(":3000", nil)
}
