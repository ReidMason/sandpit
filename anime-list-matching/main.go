package main

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/config"
	"anime-list-matching/internal/matcher"
	"anime-list-matching/internal/migrations"
	"anime-list-matching/internal/plex"
	"context"
	"database/sql"
	"log"
	"sync"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/joho/godotenv"
)

type SeriesWithEps struct {
	Series   plex.PlexSeries
	Episodes int
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := config.GetPlexToken()
	plexUrl := config.GetPlexURL()

	plexAPI := plex.New(plexUrl, token)

	// Setup DB
	connectionString := config.GetDBConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	ctx := context.Background()
	queries := animeDb.New(db)

	migrations.ApplyMigrations(db)

	// Finished setup
	series := getSeasonsWithEpisodes(plexAPI)
	createMatches(series, queries, ctx)

	for _, s := range series {
		findMappingForSeries(s, queries, ctx)
	}
}

func findMappingForSeries(series SeriesWithEps, queries *animeDb.Queries, ctx context.Context) ([]animeDb.Animemapping, error) {
	return queries.GetMappings(ctx, series.Series.RatingKey)
}

func createMatches(series []SeriesWithEps, queries *animeDb.Queries, ctx context.Context) {
	matches := 0
	// Match the series
	for _, s := range series {
		searchResults := anilist.SearchAnime(s.Series.Title, queries, ctx)
		if len(searchResults) == 0 {
			log.Println("FAILED: Failed to find search results for", s.Series.Title)
			continue
		}

		var err error
		for _, result := range searchResults {
			path, err := matcher.MatchAnime(result.ID, make([]anilist.Anime, 0), s.Episodes, queries, ctx)
			if err != nil {
				continue
			}

			if len(path) > 0 {
				matches += 1
				for _, p := range path {
					queries.SaveMapping(ctx, animeDb.SaveMappingParams{
						Anilistid:    p.ID,
						Plexseriesid: s.Series.RatingKey,
					})
				}
				break
			}
		}

		if err != nil {
			log.Println(s.Series.Title, s.Episodes, err)
		}
	}

	log.Printf("Matched %d/%d", matches, len(series))
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
				if season.Index == 0 {
					continue
				}
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
