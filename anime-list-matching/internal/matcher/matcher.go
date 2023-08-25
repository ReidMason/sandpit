package matcher

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"context"
	"errors"
	"log"
	"strings"
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
