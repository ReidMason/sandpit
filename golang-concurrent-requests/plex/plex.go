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

func (p *Plex) GetEpisodes(seriesId string) []Episode {
	url, err := url.Parse(p.url + fmt.Sprintf("/library/metadata/%s/children", seriesId))

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

	var data EpisodesResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.MediaContainer.Metadata
}

func (p *Plex) GetSeasons(seriesId string) []Season {
	url, err := url.Parse(p.url + fmt.Sprintf("/library/metadata/%s/children", seriesId))

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

	var data SeasonsResponse
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data.MediaContainer.Metadata
}

func (p *Plex) GetSeries(libraryId uint8) []PlexSeries {
	url, err := url.Parse(p.url + fmt.Sprintf("/library/sections/%d/all", libraryId))
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
	RatingKey string `json:"ratingKey"`
	// Key                   string  `json:"key"`
	// SkipChildren          bool    `json:"skipChildren,omitempty"`
	// GUID                  string  `json:"guid"`
	// Studio                string  `json:"studio,omitempty"`
	// Type                  string  `json:"type"`
	Title string `json:"title"`
	// TitleSort             string  `json:"titleSort,omitempty"`
	// Summary               string  `json:"summary"`
	// Index                 int     `json:"index"`
	// Rating                float64 `json:"rating,omitempty"`
	// ViewCount             int     `json:"viewCount,omitempty"`
	// SkipCount             int     `json:"skipCount,omitempty"`
	// LastViewedAt          int     `json:"lastViewedAt,omitempty"`
	// Year                  int     `json:"year"`
	// Thumb                 string  `json:"thumb"`
	// Art                   string  `json:"art"`
	// Banner                string  `json:"banner,omitempty"`
	// Duration              int     `json:"duration"`
	// OriginallyAvailableAt string  `json:"originallyAvailableAt"`
	// LeafCount             int     `json:"leafCount"`
	// ViewedLeafCount       int     `json:"viewedLeafCount"`
	// ChildCount            int     `json:"childCount"`
	// AddedAt               int     `json:"addedAt"`
	// UpdatedAt             int     `json:"updatedAt"`
	// Genre                 []struct {
	// 	Tag string `json:"tag"`
	// } `json:"Genre"`
	// Role []struct {
	// 	Tag string `json:"tag"`
	// } `json:"Role,omitempty"`
	// OriginalTitle       string  `json:"originalTitle,omitempty"`
	// ContentRating       string  `json:"contentRating,omitempty"`
	// AudienceRating      float64 `json:"audienceRating,omitempty"`
	// Theme               string  `json:"theme,omitempty"`
	// AudienceRatingImage string  `json:"audienceRatingImage,omitempty"`
	// Country             []struct {
	// 	Tag string `json:"tag"`
	// } `json:"Country,omitempty"`
	// SeasonCount     int     `json:"seasonCount,omitempty"`
	// Tagline         string  `json:"tagline,omitempty"`
	// PrimaryExtraKey string  `json:"primaryExtraKey,omitempty"`
	// UserRating      float64 `json:"userRating,omitempty"`
	// LastRatedAt     int     `json:"lastRatedAt,omitempty"`
}

type EpisodesResponse struct {
	MediaContainer struct {
		// Size                     int       `json:"size"`
		// AllowSync                bool      `json:"allowSync"`
		// Art                      string    `json:"art"`
		// GrandparentContentRating string    `json:"grandparentContentRating"`
		// GrandparentRatingKey     int       `json:"grandparentRatingKey"`
		// GrandparentStudio        string    `json:"grandparentStudio"`
		// GrandparentTheme         string    `json:"grandparentTheme"`
		// GrandparentThumb         string    `json:"grandparentThumb"`
		// GrandparentTitle         string    `json:"grandparentTitle"`
		// Identifier               string    `json:"identifier"`
		// Key                      string    `json:"key"`
		// LibrarySectionID         int       `json:"librarySectionID"`
		// LibrarySectionTitle      string    `json:"librarySectionTitle"`
		// LibrarySectionUUID       string    `json:"librarySectionUUID"`
		// MediaTagPrefix           string    `json:"mediaTagPrefix"`
		// MediaTagVersion          int       `json:"mediaTagVersion"`
		// Nocache                  bool      `json:"nocache"`
		// ParentIndex              int       `json:"parentIndex"`
		// ParentTitle              string    `json:"parentTitle"`
		// ParentYear               int       `json:"parentYear"`
		// Theme                    string    `json:"theme"`
		// Thumb                    string    `json:"thumb"`
		// Title1                   string    `json:"title1"`
		// Title2                   string    `json:"title2"`
		// ViewGroup                string    `json:"viewGroup"`
		// ViewMode                 int       `json:"viewMode"`
		Metadata []Episode `json:"Metadata"`
	} `json:"MediaContainer"`
}

type Episode struct {
	RatingKey string `json:"ratingKey"`
	// Key                   string  `json:"key"`
	// ParentRatingKey       string  `json:"parentRatingKey"`
	// GrandparentRatingKey  string  `json:"grandparentRatingKey"`
	// GUID                  string  `json:"guid"`
	// ParentGUID            string  `json:"parentGuid"`
	// GrandparentGUID       string  `json:"grandparentGuid"`
	// Type                  string  `json:"type"`
	// Title                 string  `json:"title"`
	// TitleSort             string  `json:"titleSort,omitempty"`
	// GrandparentKey        string  `json:"grandparentKey"`
	// ParentKey             string  `json:"parentKey"`
	// GrandparentTitle      string  `json:"grandparentTitle"`
	// ParentTitle           string  `json:"parentTitle"`
	// OriginalTitle         string  `json:"originalTitle"`
	// ContentRating         string  `json:"contentRating"`
	// Summary               string  `json:"summary"`
	// Index                 int     `json:"index"`
	// ParentIndex           int     `json:"parentIndex"`
	// AudienceRating        float64 `json:"audienceRating"`
	// ViewCount             int     `json:"viewCount"`
	// SkipCount             int     `json:"skipCount,omitempty"`
	// LastViewedAt          int     `json:"lastViewedAt"`
	// ParentYear            int     `json:"parentYear"`
	// Thumb                 string  `json:"thumb"`
	// Art                   string  `json:"art"`
	// ParentThumb           string  `json:"parentThumb"`
	// GrandparentThumb      string  `json:"grandparentThumb"`
	// GrandparentArt        string  `json:"grandparentArt"`
	// GrandparentTheme      string  `json:"grandparentTheme"`
	// Duration              int     `json:"duration"`
	// OriginallyAvailableAt string  `json:"originallyAvailableAt"`
	// AddedAt               int     `json:"addedAt"`
	// UpdatedAt             int     `json:"updatedAt"`
	// AudienceRatingImage   string  `json:"audienceRatingImage"`
	// ChapterSource         string  `json:"chapterSource,omitempty"`
	// Media                 []struct {
	// 	ID              int     `json:"id"`
	// 	Duration        int     `json:"duration"`
	// 	Bitrate         int     `json:"bitrate"`
	// 	Width           int     `json:"width"`
	// 	Height          int     `json:"height"`
	// 	AspectRatio     float64 `json:"aspectRatio"`
	// 	AudioChannels   int     `json:"audioChannels"`
	// 	AudioCodec      string  `json:"audioCodec"`
	// 	VideoCodec      string  `json:"videoCodec"`
	// 	VideoResolution string  `json:"videoResolution"`
	// 	Container       string  `json:"container"`
	// 	VideoFrameRate  string  `json:"videoFrameRate"`
	// 	AudioProfile    string  `json:"audioProfile"`
	// 	VideoProfile    string  `json:"videoProfile"`
	// 	Part            []struct {
	// 		ID           int    `json:"id"`
	// 		Key          string `json:"key"`
	// 		Duration     int    `json:"duration"`
	// 		File         string `json:"file"`
	// 		Size         int    `json:"size"`
	// 		AudioProfile string `json:"audioProfile"`
	// 		Container    string `json:"container"`
	// 		HasThumbnail string `json:"hasThumbnail"`
	// 		Indexes      string `json:"indexes"`
	// 		VideoProfile string `json:"videoProfile"`
	// 	} `json:"Part"`
	// } `json:"Media"`
	// Director []struct {
	// 	Tag string `json:"tag"`
	// } `json:"Director"`
	// Writer []struct {
	// 	Tag string `json:"tag"`
	// } `json:"Writer"`
	// Role []struct {
	// 	Tag string `json:"tag"`
	// } `json:"Role,omitempty"`
}

type SeasonsResponse struct {
	MediaContainer struct {
		// Size                int    `json:"size"`
		// AllowSync           bool   `json:"allowSync"`
		// Art                 string `json:"art"`
		// Identifier          string `json:"identifier"`
		// Key                 string `json:"key"`
		// LibrarySectionID    int    `json:"librarySectionID"`
		// LibrarySectionTitle string `json:"librarySectionTitle"`
		// LibrarySectionUUID  string `json:"librarySectionUUID"`
		// MediaTagPrefix      string `json:"mediaTagPrefix"`
		// MediaTagVersion     int    `json:"mediaTagVersion"`
		// Nocache             bool   `json:"nocache"`
		// ParentIndex         int    `json:"parentIndex"`
		// ParentTitle         string `json:"parentTitle"`
		// ParentYear          int    `json:"parentYear"`
		// Summary             string `json:"summary"`
		// Theme               string `json:"theme"`
		// Thumb               string `json:"thumb"`
		// Title1              string `json:"title1"`
		// Title2              string `json:"title2"`
		// ViewGroup           string `json:"viewGroup"`
		// ViewMode            int    `json:"viewMode"`
		// Directory           []struct {
		// 	LeafCount       int    `json:"leafCount"`
		// 	Thumb           string `json:"thumb"`
		// 	ViewedLeafCount int    `json:"viewedLeafCount"`
		// 	Key             string `json:"key"`
		// 	Title           string `json:"title"`
		// } `json:"Directory"`
		Metadata []Season `json:"Metadata"`
	} `json:"MediaContainer"`
}

type Season struct {
	RatingKey string `json:"ratingKey"`
	// Key             string `json:"key"`
	// ParentRatingKey string `json:"parentRatingKey"`
	// GUID            string `json:"guid"`
	// ParentGUID      string `json:"parentGuid"`
	// ParentStudio    string `json:"parentStudio"`
	// Type            string `json:"type"`
	// Title           string `json:"title"`
	// TitleSort       string `json:"titleSort"`
	// ParentKey       string `json:"parentKey"`
	ParentTitle string `json:"parentTitle"`
	// Summary         string `json:"summary"`
	Index int `json:"index"`
	// ParentIndex     int    `json:"parentIndex"`
	// Year            int    `json:"year"`
	// Thumb           string `json:"thumb"`
	// Art             string `json:"art"`
	// ParentThumb     string `json:"parentThumb"`
	// ParentTheme     string `json:"parentTheme"`
	// LeafCount       int    `json:"leafCount"`
	// ViewedLeafCount int    `json:"viewedLeafCount"`
	// AddedAt         int    `json:"addedAt"`
	// UpdatedAt       int    `json:"updatedAt"`
	// ViewCount       int    `json:"viewCount,omitempty"`
	// LastViewedAt    int    `json:"lastViewedAt,omitempty"`
	// SkipCount       int    `json:"skipCount,omitempty"`
}
