package main

import (
	"concurrent-requests/plex"
	"log"
)

func main() {
	plex := plex.New(
		"",
		"",
	)

	series := plex.GetSeries(1)
	for _, s := range series {
		log.Println(s.Title)
	}
}
