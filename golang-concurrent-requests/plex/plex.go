package plex

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

type Plex struct {
	url   string
	token string
}

func New(url string, token string) *Plex {
	return &Plex{url, token}
}

func (p *Plex) GetSeries(library_id uint8) []PlexSeries {
	url, err := url.Parse(p.url + fmt.Sprintf("/library/sections/%d/all", library_id))
	if err != nil {
		log.Fatal(err)
	}

	q := url.Query()
	q.Set("X-Plex-Token", p.token)
	url.RawQuery = q.Encode()

	req, err := http.NewRequest("GET", url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	headers := http.Header{
		"Content-Type": {"application/json"},
		"Accept":       {"application/json"},
	}
	// Add headers
	req.Header = headers

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var data PlexResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.MediaContainer.Metadata
}

type PlexResponse struct {
	MediaContainer struct {
		Size                int          `json:"size"`
		AllowSync           bool         `json:"allowSync"`
		Art                 string       `json:"art"`
		Identifier          string       `json:"identifier"`
		LibrarySectionID    int          `json:"librarySectionID"`
		LibrarySectionTitle string       `json:"librarySectionTitle"`
		LibrarySectionUUID  string       `json:"librarySectionUUID"`
		MediaTagPrefix      string       `json:"mediaTagPrefix"`
		MediaTagVersion     int          `json:"mediaTagVersion"`
		Nocache             bool         `json:"nocache"`
		Thumb               string       `json:"thumb"`
		Title1              string       `json:"title1"`
		Title2              string       `json:"title2"`
		ViewGroup           string       `json:"viewGroup"`
		ViewMode            int          `json:"viewMode"`
		Metadata            []PlexSeries `json:"Metadata"`
	} `json:"MediaContainer"`
}

type PlexSeries struct {
	RatingKey             string  `json:"ratingKey"`
	Key                   string  `json:"key"`
	SkipChildren          bool    `json:"skipChildren,omitempty"`
	GUID                  string  `json:"guid"`
	Studio                string  `json:"studio,omitempty"`
	Type                  string  `json:"type"`
	Title                 string  `json:"title"`
	TitleSort             string  `json:"titleSort,omitempty"`
	Summary               string  `json:"summary"`
	Index                 int     `json:"index"`
	Rating                float64 `json:"rating,omitempty"`
	ViewCount             int     `json:"viewCount,omitempty"`
	SkipCount             int     `json:"skipCount,omitempty"`
	LastViewedAt          int     `json:"lastViewedAt,omitempty"`
	Year                  int     `json:"year"`
	Thumb                 string  `json:"thumb"`
	Art                   string  `json:"art"`
	Banner                string  `json:"banner,omitempty"`
	Duration              int     `json:"duration"`
	OriginallyAvailableAt string  `json:"originallyAvailableAt"`
	LeafCount             int     `json:"leafCount"`
	ViewedLeafCount       int     `json:"viewedLeafCount"`
	ChildCount            int     `json:"childCount"`
	AddedAt               int     `json:"addedAt"`
	UpdatedAt             int     `json:"updatedAt"`
	Genre                 []struct {
		Tag string `json:"tag"`
	} `json:"Genre"`
	Role []struct {
		Tag string `json:"tag"`
	} `json:"Role,omitempty"`
	OriginalTitle       string  `json:"originalTitle,omitempty"`
	ContentRating       string  `json:"contentRating,omitempty"`
	AudienceRating      float64 `json:"audienceRating,omitempty"`
	Theme               string  `json:"theme,omitempty"`
	AudienceRatingImage string  `json:"audienceRatingImage,omitempty"`
	Country             []struct {
		Tag string `json:"tag"`
	} `json:"Country,omitempty"`
	SeasonCount     int     `json:"seasonCount,omitempty"`
	Tagline         string  `json:"tagline,omitempty"`
	PrimaryExtraKey string  `json:"primaryExtraKey,omitempty"`
	UserRating      float64 `json:"userRating,omitempty"`
	LastRatedAt     int     `json:"lastRatedAt,omitempty"`
}
