package main

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/matcher"
	"anime-list-matching/internal/migrations"
	"anime-list-matching/internal/plex"
	"context"
	"database/sql"
	"log"
	"os"
	"sync"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/joho/godotenv"
)

type SeriesWithEps struct {
	Episodes int
	Series   plex.PlexSeries
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("PLEX_TOKEN")
	plexUrl := os.Getenv("PLEX_URL")

	plexAPI := plex.New(plexUrl, token)
	series := getSeasonsWithEpisodes(plexAPI)

	// Setup DB
	connectionString := "postgres://user:password@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	ctx := context.Background()
	queries := animeDb.New(db)

	migrations.ApplyMigrations(db)

	// Match the series
	for _, s := range series {
		searchResults := anilist.SearchAnime(s.Series.Title, queries, ctx)
		if len(searchResults) == 0 {
			log.Println("FAILED: Failed to find search results for", s.Series.Title)
			continue
		}

		result := searchResults[0]
		path, err := matcher.MatchAnime(result.ID, make([]anilist.Anime, 0), s.Episodes, queries, ctx)
		if err != nil {
			log.Println(s.Series.Title, err)
		}

		if len(path) > 0 {
			log.Println("SUCCESS: Found match for", s.Series.Title)
		}
		// matcher.PrintTraversalPath(path)
	}
}

func getSeasonsWithEpisodes(plexAPI *plex.Plex) []SeriesWithEps {
	log.Println("Started getting full data for all Plex series")
	series := plexAPI.GetSeries(1)
	log.Printf("%d series to process", len(series))

	var wg sync.WaitGroup
	var seriesWithEps []SeriesWithEps
	for _, s := range series {
		wg.Add(1)
		go func(plexAPI *plex.Plex, series plex.PlexSeries) {
			seasons := plexAPI.GetSeasons(series.RatingKey)
			episodes := 0
			for _, season := range seasons {
				episodes += season.LeafCount
			}
			seriesWithEps = append(seriesWithEps, SeriesWithEps{
				Series:   series,
				Episodes: episodes,
			})

			defer wg.Done()
		}(plexAPI, s)
	}

	wg.Wait()

	return seriesWithEps
}
