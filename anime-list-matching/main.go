package main

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/matcher"
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

	targetEps := 88
	path, err := matcher.MatchAnime(16498, make([]anilist.Anime, 0), targetEps, queries, ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Done mathcing anime found:")
	matcher.PrintTraversalPath(path)
}
