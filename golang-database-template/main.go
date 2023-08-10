package main

import (
	"golang-database-template/internal/migrations"
	"log"
)

func main() {
	scratchMigrateDatabase()
}

func scratchMigrateDatabase() {
	connectionString := "postgres://user:password@localhost:5432/testdb?sslmode=disable"
	migrationInstance := migrations.GetDatabaseMigrationInstance(connectionString)

	// Clear the database completely to clean reset
	// There's a weird bug where this breaks the up migration
	err := migrationInstance.Drop()
	if err != nil {
		log.Fatal(err)
	}
	migrationInstance.Close()

	migrationInstance = migrations.GetDatabaseMigrationInstance(connectionString)
	err = migrationInstance.Up()
	if err != nil {
		log.Fatal(err)
	}
}
