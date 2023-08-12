package main

import (
	animeDb "anime-list-matching/internal/animedb"
	"anime-list-matching/internal/dtos"
	"anime-list-matching/internal/migrations"
	"context"
	"database/sql"
	"encoding/json"
	"log"
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

	thing := dtos.AnimeResponse{
		Name: "testing",
	}
	thingJson, err := json.Marshal(thing)
	log.Println(thingJson)

	_, err = queries.CacheAnimeResult(ctx, animeDb.CacheAnimeResultParams{
		ID:       1,
		Response: thingJson,
	})
	if err != nil {
		log.Println(err)
	}

	res, err := queries.GetCachedAnimeResult(ctx, 1)
	if err != nil {
		log.Println(err)
	}
	var data dtos.AnimeResponse
	err = json.Unmarshal(res, &data)
	if err != nil {
		log.Println(err)
	}

	log.Println(data)
}
