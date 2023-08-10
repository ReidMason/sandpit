package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"
	"trying-sqlc/testDatabase"

	_ "github.com/lib/pq"
)

func main() {
	ctx := context.Background()

	db, err := sql.Open("postgres", "user=user password=password dbname=testdb sslmode=disable")
	if err != nil {
		log.Panic(err)
	}

	queries := testDatabase.New(db)

	// list all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		log.Panic(err)
	}
	log.Println(authors)

	// create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, testDatabase.CreateAuthorParams{
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		log.Panic(err)
	}
	log.Println(insertedAuthor)

	// get the author we just inserted
	fetchedAuthor, err := queries.GetAuthor(ctx, insertedAuthor.ID)
	if err != nil {
		log.Panic(err)
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedAuthor, fetchedAuthor))
}
