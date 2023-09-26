package matcher

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/plex"
	"context"
	"errors"
	"log"
	"strings"
	"sync"
)

func MatchAnime(animeId int32, path []anilist.Anime, targetEps int, queries *animeDb.Queries, ctx context.Context) ([]anilist.Anime, error) {
	if animeId == 0 {
		return nil, errors.New("No anime found")
	}

	anime := anilist.GetAnime(animeId, queries, ctx)

	if countEpisodes(path)+anime.Episodes > targetEps {
		return path, nil
	}

	path = append(path, anime)

	isContinuingSeries := anime.Status == anilist.Releasing && anime.Episodes == 0
	if targetEps == calculateEpisodes(path) || isContinuingSeries {
		return path, nil
	}

	sequelId := getRelationSeriesId(anime.Relations, anilist.Sequel)
	res, err := MatchAnime(sequelId, path, targetEps, queries, ctx)
	if err == nil {
		return res, nil
	}

	sideStoryId := getRelationSeriesId(anime.Relations, anilist.SideStory)
	res, err = MatchAnime(sideStoryId, path, targetEps, queries, ctx)
	if err == nil {
		return res, nil
	}

	return nil, errors.New("No matches found")
}

func countEpisodes(animes []anilist.Anime) int {
	count := 0
	for _, anime := range animes {
		count += anime.Episodes
	}

	return count
}

func getRelationSeriesId(relations []anilist.Relation, targetRelation string) int32 {
	for _, relation := range relations {
		if relation.Relation == targetRelation {
			return relation.ID
		}
	}

	return 0
}

func calculateEpisodes(animeList []anilist.Anime) int {
	count := 0
	for _, anime := range animeList {
		count += anime.Episodes
	}

	return count
}

func PrintTraversalPath(path []anilist.Anime) {
	for i, anime := range path {
		log.Println(strings.Repeat(" ", i), getTraversalPathPrefix(i), anime.Title.Romaji)
	}
}

func getTraversalPathPrefix(index int) string {
	if index == 0 {
		return ""
	}

	return " ∟"
}

func CreateMatches(series []SeriesWithEps, queries *animeDb.Queries, ctx context.Context) {
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
			path, err := MatchAnime(result.ID, make([]anilist.Anime, 0), s.Episodes, queries, ctx)
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

type SeriesWithEps struct {
	Series   plex.PlexSeries
	Episodes int
}

func GetSeasonsWithEpisodes(plexAPI *plex.Plex) []SeriesWithEps {
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
