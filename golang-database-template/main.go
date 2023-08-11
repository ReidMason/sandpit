package main

import (
	"context"
	"database/sql"
	"fmt"
	"golang-database-template/internal/migrations"
	"golang-database-template/internal/testDb"
	"log"

	"github.com/golang-migrate/migrate"
)

func main() {
	// Note: The 'testdb' database has already been created using docker compose, the database needs to exist first
	connectionString := "postgres://user:password@localhost:5432/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	scratchMigrateDatabase(db)
	doRandomQueries(db)
}

func doRandomQueries(db *sql.DB) {
	ctx := context.Background()
	queries := testDb.New(db)

	// Create an author
	insertedAuthor, err := queries.CreateAuthor(ctx, testDb.CreateAuthorParams{
		ID:   2,
		Name: "Brian Kernighan",
		Bio:  sql.NullString{String: "Co-author of The C Programming Language and The Go Programming Language", Valid: true},
	})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Inserted author:", insertedAuthor)

	// List all authors
	authors, err := queries.ListAuthors(ctx)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Authors:", authors)

	// List all authors and their books
	authorsAndBooks, err := queries.GetAuthorsWithBooks(ctx, []int32{1, 2})
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Authors and their books:", authorsAndBooks)
}

func scratchMigrateDatabase(db *sql.DB) {
	migrationInstance := migrations.GetDatabaseMigrationInstance(db)
	migrationInstance.Drop()

	migrationInstance = migrations.GetDatabaseMigrationInstance(db)
	err := migrationInstance.Up()

	if err != nil && err.Error() != migrate.ErrNoChange.Error() {
		log.Fatal(err)
	}
}
