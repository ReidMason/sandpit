package migrations

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func ApplyMigrations(db *sql.DB) {
	driverName := "postgres"
	migrationsFilepath := "file://internal/migrations/sqlMigrations"

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Panic(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrationsFilepath,
		driverName, driver)
	if err != nil {
		log.Panic(err)
	}

	m.Up()
}
