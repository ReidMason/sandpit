package main

import (
	"concurrent-requests/plex"
	"log"
	"sync"
)

type SeasonWithEps struct {
	Episodes []plex.Episode
	Season   plex.Season
}

func main() {
	plexAPI := plex.New(
		"",
		"",
	)

	log.Println("Started getting full data for all Plex series")
	series := plexAPI.GetSeries(1)
	log.Printf("%d series to process", len(series))

	ch := make(chan SeasonWithEps)
	chSeasons := make(chan []plex.Season)

	var wg sync.WaitGroup
	for _, s := range series {
		wg.Add(1)
		go getSeasons(*plexAPI, s.RatingKey, chSeasons, &wg)
	}

	go func() {
		for seasons := range chSeasons {
			for _, season := range seasons {
				if season.Index == 0 {
					continue
				}
				wg.Add(1)
				go getEpisodes(*plexAPI, season, ch, &wg)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(chSeasons)
		close(ch)
	}()

	var episodes []SeasonWithEps
	for res := range ch {
		episodes = append(episodes, res)
	}

	for _, ep := range episodes {
		log.Printf("%s: season %d - episodes: %d", ep.Season.ParentTitle, ep.Season.Index, len(ep.Episodes))
	}
}

func getEpisodes(plexAPI plex.Plex, season plex.Season, ch chan<- SeasonWithEps, wg *sync.WaitGroup) {
	eps := plexAPI.GetEpisodes(season.RatingKey)
	ch <- SeasonWithEps{
		Season:   season,
		Episodes: eps,
	}
	defer wg.Done()
}

func getSeasons(plexAPI plex.Plex, ratingKey string, ch chan<- []plex.Season, wg *sync.WaitGroup) {
	ch <- plexAPI.GetSeasons(ratingKey)
	defer wg.Done()
}
