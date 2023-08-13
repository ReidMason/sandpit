package main

import (
	animeDb "anime-list-matching/internal/animedb"
	"anime-list-matching/internal/dtos"
	"anime-list-matching/internal/migrations"
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"io"
	"log"
	"net/http"
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

	animeResult := getAnime(113415)
	byteResult, err := json.Marshal(animeResult)
	if err != nil {
		log.Print(err)
	}

	_, err = queries.CacheAnimeResult(ctx, animeDb.CacheAnimeResultParams{
		ID:       1,
		Response: byteResult,
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

type GraphqlBody[T any] struct {
	Query     string `json:"query"`
	Variables T      `json:"variables"`
}

type GetAnimeVariables struct {
	AnimeId int `json:"anime_id"`
}

func getAnime(animeId int) dtos.AnimeResponse {
	url := "https://graphql.anilist.co/"

	query := `query ($anime_id: Int) {
    Media(id: $anime_id, type: ANIME) {
      id
      format
      episodes
      synonyms
      status
      endDate {
        year
        month
        day
        }
      startDate {
        year
        month
        day
        }
      title {
        english
        romaji
        }
      relations {
        edges {
          relationType
            }
        nodes {
          id
          format
          endDate {
            year
            month
            day
                }
            startDate {
                year
                month
                day
                }
            }
        }
    }
}`

	body := GraphqlBody[GetAnimeVariables]{
		Query: query,
		Variables: GetAnimeVariables{
			AnimeId: animeId,
		},
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Print(err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data dtos.AnimeResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Print(err)
	}

	return data
}
