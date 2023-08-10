# Golang database template

This project is an outline and a testing ground for projects that use databases in Golang.\
It's here to reference what libraries and being used, how they are being used and how to set the project up.

# Setup

### Migrations

Migrations are handled by [golang-migrate](https://github.com/golang-migrate/migrate)\
First you need to install the [golang-migrate CLI tool](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)\
You can then create a migration

```bash
migrate create -ext sql -dir internal/migrations/sqlMigrations -seq initial-db-setup
```

Add add your SQL scripts in the up/down files.
