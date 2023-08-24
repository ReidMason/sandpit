package main

import (
	"anime-list-matching/internal/plex"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type SeriesWithEps struct {
	Episodes int // []plex.Episode
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

	for _, s := range series {
		log.Println(s.Series.Title, s.Episodes)
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
