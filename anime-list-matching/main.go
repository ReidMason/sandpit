package main

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/migrations"
	"context"
	"database/sql"
	"log"
)

func main() {
	connectionString := "postgres://user:password@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	migrations.ApplyMigrations(db)

	ctx := context.Background()
	queries := animeDb.New(db)

	animeResult := anilist.GetAnime(113415, queries, ctx)

	log.Print(animeResult)
}
