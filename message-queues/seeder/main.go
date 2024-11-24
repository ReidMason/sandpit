package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/exp/rand"
)

type Photo struct {
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
	AlbumId      int    `json:"albumId"`
	Id           int    `json:"id"`
}

func main() {
	// Fetch the photo data
	url := "https://jsonplaceholder.typicode.com/photos"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var photos []Photo
	err = json.Unmarshal(data, &photos)
	if err != nil {
		log.Fatalln(err)
	}

	// Shuffle the photos
	for i := range photos {
		j := rand.Intn(i + 1)
		photos[i], photos[j] = photos[j], photos[i]
	}

	// Split into 100 chunks
	chunkSize := len(photos) / 100
	chunks := make([][]Photo, 100)
	for i := 0; i < 100; i++ {
		start := i * chunkSize
		end := (i + 1) * chunkSize
		chunks[i] = photos[start:end]
	}

	// Save the chunks to a file
	file, err := os.Create("photos.json")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	err = enc.Encode(chunks)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Done")
}
