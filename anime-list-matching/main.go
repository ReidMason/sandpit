package main

import (
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/config"
	"anime-list-matching/internal/matcher"
	"anime-list-matching/internal/migrations"
	"anime-list-matching/internal/plex"
	"context"
	"database/sql"
	"log"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"

	"github.com/joho/godotenv"
)

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
	series := matcher.GetSeasonsWithEpisodes(plexAPI)
	matcher.CreateMatches(series, queries, ctx)

	for _, s := range series {
		getMappingForSeries(s, queries, ctx)
	}
}

func getMappingForSeries(series matcher.SeriesWithEps, queries *animeDb.Queries, ctx context.Context) ([]animeDb.Animemapping, error) {
	return queries.GetMappings(ctx, series.Series.RatingKey)
}
