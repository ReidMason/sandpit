package anilist

import (
	"anime-list-matching/internal/animeDb"
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type GraphqlBody[T any] struct {
	Query     string `json:"query"`
	Variables T      `json:"variables"`
}

type GetAnimeVariables struct {
	AnimeId int32 `json:"anime_id"`
}

type SearchAnimeVariables struct {
	AnimeName string `json:"anime_name"`
}

func GetAnime(animeId int32, queries *animeDb.Queries, ctx context.Context) Anime {
	res, err := queries.GetCachedAnimeResult(ctx, animeId)
	if err == nil {
		var anime Anime
		err = json.Unmarshal(res, &anime)
		if err != nil {
			log.Print(err)
		}

		return anime
	}

	log.Printf("ID: %d Not found making request", animeId)

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

	time.Sleep(1 * time.Second)

	var data AnimeResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Print(err)
	}

	anime := cleanUpAnimeResult(data)
	animeJson, err := json.Marshal(anime)
	if err != nil {
		log.Println("Failed to serialize anime for caching", err)
	}

	_, err = queries.CacheAnimeResult(ctx, animeDb.CacheAnimeResultParams{
		ID:       int32(data.Data.Media.ID),
		Response: animeJson,
	})
	if err != nil {
		log.Println(err)
	}

	return anime
}

func SearchAnime(searchTerm string, queries *animeDb.Queries, ctx context.Context) []AnimeResult {
	res, err := queries.GetCachedAnimeSearchResult(ctx, searchTerm)
	if err == nil {
		var anime []AnimeResult
		err = json.Unmarshal(res, &anime)
		if err != nil {
			log.Print(err)
		}

		return anime
	}

	log.Printf("Anime search not found making request %s", searchTerm)

	url := "https://graphql.anilist.co/"

	query := `query ($anime_name: String) {
                Page(perPage: 10) {
                    media(search: $anime_name, type: ANIME, sort: SEARCH_MATCH) {
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
                                episodes
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
                }
            }`

	body := GraphqlBody[SearchAnimeVariables]{
		Query: query,
		Variables: SearchAnimeVariables{
			AnimeName: searchTerm,
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

	time.Sleep(1 * time.Second)

	var data AnimeSearchResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Print(err)
	}

	animeJson, err := json.Marshal(data.Data.Page.Media)
	if err != nil {
		log.Println("Failed to serialize anime search result for caching", err)
	}

	_, err = queries.CacheAnimeSearch(ctx, animeDb.CacheAnimeSearchParams{
		Searchterm: searchTerm,
		Response:   animeJson,
	})
	if err != nil {
		log.Println(err)
	}

	return data.Data.Page.Media
}

func cleanUpAnimeResult(response AnimeResponse) Anime {
	anime := Anime{
		ID:        response.Data.Media.ID,
		Format:    response.Data.Media.Format,
		Episodes:  response.Data.Media.Episodes,
		Synonyms:  response.Data.Media.Synonyms,
		Status:    response.Data.Media.Status,
		EndDate:   response.Data.Media.EndDate,
		StartDate: response.Data.Media.StartDate,
		Title:     response.Data.Media.Title,
	}

	for i, edge := range response.Data.Media.Relations.Edges {
		node := response.Data.Media.Relations.Nodes[i]

		anime.Relations = append(anime.Relations, Relation{
			ID:        node.ID,
			Format:    node.Format,
			Relation:  edge.RelationType,
			EndDate:   node.EndDate,
			StartDate: node.StartDate,
		})
	}

	return anime
}

type AnimeSearchResponse struct {
	Data struct {
		Page struct {
			Media []AnimeResult `json:"media"`
		} `json:"Page"`
	} `json:"data"`
}
