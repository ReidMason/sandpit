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
	var episodes []SeasonWithEps
	var wg sync.WaitGroup
	count := 1

	for _, s := range series {
		wg.Add(1)

		go func(series plex.PlexSeries) {
			count += 1
			seasons := plexAPI.GetSeasons(series.RatingKey)

			for _, season := range seasons {
				if season.Index == 0 {
					continue
				}
				wg.Add(1)

				go func(season plex.Season) {
					count += 1
					eps := plexAPI.GetEpisodes(season.RatingKey)
					episodes = append(episodes, SeasonWithEps{
						Season:   season,
						Episodes: eps,
					})

					defer wg.Done()
				}(season)
			}

			defer wg.Done()
		}(s)
	}

	wg.Wait()

	// for _, ep := range episodes {
	// 	log.Printf("%s: season %d - episodes: %d", ep.Season.ParentTitle, ep.Season.Index, len(ep.Episodes))
	// }
	log.Printf("Requests made: %d", count)
}
