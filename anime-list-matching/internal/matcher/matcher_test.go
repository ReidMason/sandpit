package matcher

import (
	"anime-list-matching/internal/anilist"
	"anime-list-matching/internal/animeDb"
	"context"
	"database/sql"
	"log"
	"testing"

	_ "github.com/lib/pq"
)

func TestMatchExactEpisodeMatch(t *testing.T) {
	tests := []struct {
		startingAnimeId int32
		targetEps       int
		expectedIds     []int
	}{
		{
			16498,
			88,
			[]int{16498, 20958, 99147, 104578, 110277, 131681, 146984},
		},
	}

	connectionString := "postgres://user:password@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Panic("Failed to connect to database", err)
	}

	ctx := context.Background()
	queries := animeDb.New(db)

	for _, test := range tests {
		test := test
		res, err := MatchAnime(test.startingAnimeId, make([]anilist.Anime, 0), test.targetEps, queries, ctx)
		if err != nil {
			t.Errorf("Failed to find match")
		}

		for i, r := range res {
			expected := test.expectedIds[i]
			if r.ID != expected {
				t.Errorf("Wrong anime id found. Expected: '%d' found '%d'", expected, r.ID)
			}
		}
	}
}
