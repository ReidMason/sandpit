package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

var data Response

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.GET("/SayHello", sayHello)

	resp, _ := http.Get("https://jsonplaceholder.typicode.com/photos")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	err := json.Unmarshal(body, &data)
	log.Printf("%v", err)

	router.Run("localhost:8080")
}

func sayHello(c *gin.Context) {
	c.JSON(http.StatusOK, data)
}

type Response []struct {
	AlbumID      int    `json:"albumId"`
	ID           int    `json:"id"`
	Title        string `json:"title"`
	URL          string `json:"url"`
	ThumbnailURL string `json:"thumbnailUrl"`
}
