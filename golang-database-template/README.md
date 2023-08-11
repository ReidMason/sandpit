# Golang database template

This project is an outline and a testing ground for projects that use databases in Golang.\
It's here to reference what libraries and being used, how they are being used and how to set the project up.

# Setup

### Migrations

Migrations are handled by [golang-migrate](https://github.com/golang-migrate/migrate)\
First you need to install the [golang-migrate CLI tool](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)\

#### Create a migration

You can then create a migration giving it a name

```bash
migrate create -ext sql -dir internal/migrations/sqlMigrations -seq initial-db-setup
```

Add add your SQL scripts in the up/down files.

### Adding queries

All database queries are handled buy [sqlc](https://github.com/sqlc-dev/sqlc)\
Make sure the [sqlc cli tool](https://docs.sqlc.dev/en/latest/overview/install.html) is installed first.\
\
This works kind of like a facade pattern. You put all of your queries in an sql file then Go functions that make those queries are generated.\
\
The configuration for this setup is done in the `sqlc.yaml` file.\
To add a query edit the `queries.sql` and add your query following the predefined syntax.\
Then generate your functions.

```bash
sqlc generate -f sqlc/sqlc.yaml
```
