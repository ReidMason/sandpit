package main

import (
	"database/sql"
	"golang-database-template/internal/migrations"
	"log"
)

func main() {
	// Note: The 'testdb' database has already been created using docker compose, the database needs to exist first
	connectionString := "postgres://user:password@localhost:5432/testdb?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	scratchMigrateDatabase(db)
}

func scratchMigrateDatabase(db *sql.DB) {
	migrationInstance := migrations.GetDatabaseMigrationInstance(db)

	migrationInstance = migrations.GetDatabaseMigrationInstance(db)
	err := migrationInstance.Up()
	if err != nil {
		log.Fatal(err)
	}
}
