package main

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"anime-list-matching/internal/migrations"
	"context"
	"database/sql"
	"errors"
	"log"
	"strings"
)

func main() {
	connectionString := "postgres://user:password@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	migrations.ApplyMigrations(db)

	ctx := context.Background()
	queries := animeDb.New(db)

	targetEps := 88
	path, err := recurseAnime(16498, make([]anilist.Anime, 0), targetEps, queries, ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Done mathcing anime found:")
	printTraversalPath(path)
}

func recurseAnime(animeId int32, path []anilist.Anime, targetEps int, queries *animeDb.Queries, ctx context.Context) ([]anilist.Anime, error) {
	if animeId == 0 {
		return nil, errors.New("No anime found")
	}

	anime := anilist.GetAnime(animeId, queries, ctx)
	path = append(path, anime)

	if targetEps == calculateEpisodes(path) {
		return path, nil
	}

	sequelId := getRelationSeriesId(anime.Relations, anilist.Sequel)
	res, err := recurseAnime(sequelId, path, targetEps, queries, ctx)
	if err == nil {
		return res, nil
	}

	sideStoryId := getRelationSeriesId(anime.Relations, anilist.SideStory)
	res, err = recurseAnime(sideStoryId, path, targetEps, queries, ctx)
	if err == nil {
		return res, nil
	}

	return nil, errors.New("No match found")
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

func printTraversalPath(path []anilist.Anime) {
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
