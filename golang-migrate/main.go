package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://user:password@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations/",
		"postgres", driver)
	if err != nil {
		log.Panic(err)
	}

	// Migrate to latests version
	m.Up()

	// Go to that specific version
	err = m.Migrate(1)
	if err != nil {
		log.Panic(err)
	}
}
