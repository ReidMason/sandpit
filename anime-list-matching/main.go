package main

import (
	"anime-list-matching/internal/plex"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("PLEX_TOKEN")
	series := plex.GetSeries(token)

	for _, series := range series.MediaContainer.Metadata {
		log.Println(series.Title, series.SeasonCount)
	}

	// connectionString := "postgres://user:password@localhost:5432/postgres?sslmode=disable"
	// db, err := sql.Open("postgres", connectionString)
	// if err != nil {
	// 	log.Panic("Failed to connect to database", err)
	// }
	//
	// migrations.ApplyMigrations(db)
	//
	// ctx := context.Background()
	// queries := animeDb.New(db)
	//
	// targetEps := 88
	// path, err := matcher.MatchAnime(16498, make([]anilist.Anime, 0), targetEps, queries, ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// log.Print("Done mathcing anime found:")
	// matcher.PrintTraversalPath(path)

}
