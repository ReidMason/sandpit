package main

import (
	"anime-list-matching/internal/plex"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type SeasonWithEps struct {
	Episodes []plex.Episode
	Season   plex.Season
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("PLEX_TOKEN")
	plexUrl := os.Getenv("PLEX_URL")

	plexAPI := plex.New(plexUrl, token)
	seasons := getSeasonsWithEpisodes(plexAPI)

	for _, season := range seasons {
		log.Println(season.Season.ParentTitle, len(season.Episodes))
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

func getSeasonsWithEpisodes(plexAPI *plex.Plex) []SeasonWithEps {
	log.Println("Started getting full data for all Plex series")
	series := plexAPI.GetSeries(1)
	log.Printf("%d series to process", len(series))

	ch := make(chan SeasonWithEps)
	chSeasons := make(chan []plex.Season)

	var wg sync.WaitGroup
	for _, s := range series {
		wg.Add(1)
		go getSeasons(plexAPI, s.RatingKey, chSeasons, &wg)
	}

	go func() {
		for seasons := range chSeasons {
			for _, season := range seasons {
				if season.Index == 0 {
					continue
				}
				wg.Add(1)
				go getEpisodes(plexAPI, season, ch, &wg)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(chSeasons)
		close(ch)
	}()

	var seasons []SeasonWithEps
	for res := range ch {
		seasons = append(seasons, res)
	}

	return seasons
}

func getEpisodes(plexAPI *plex.Plex, season plex.Season, ch chan<- SeasonWithEps, wg *sync.WaitGroup) {
	eps := plexAPI.GetEpisodes(season.RatingKey)
	ch <- SeasonWithEps{
		Season:   season,
		Episodes: eps,
	}
	defer wg.Done()
}

func getSeasons(plexAPI *plex.Plex, ratingKey string, ch chan<- []plex.Season, wg *sync.WaitGroup) {
	ch <- plexAPI.GetSeasons(ratingKey)
	defer wg.Done()
}
