package model

import (
	"encoding/json"
	"fmt"
)

// ArtistImage represents an Artist's image
type ArtistImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

func (ai ArtistImage) String() string {
	return fmt.Sprintf("Height: %d, Width: %d, URL: %s", ai.Height, ai.Width, ai.URL)
}

// ArtistResponse represents a struct for the JSON response
type ArtistResponse struct {
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Genres     []string      `json:"genres"`
	Images     []ArtistImage `json:"images"`
	Name       string        `json:"name"`
	Popularity int           `json:"popularity"`
	Type       string        `json:"type"`
	URI        string        `json:"uri"`
}

// Artist returns an Artist given an ArtistResponse
func (ar ArtistResponse) Artist() Artist {
	return Artist{
		Name:       ar.Name,
		Popularity: ar.Popularity,
		Type:       ar.Type,
		Followers:  ar.Followers.Total,
		Genres:     ar.Genres,
		Image:      ar.Images[0],
	}
}

// Artist represents an unmarshalled Artist
type Artist struct {
	Name       string      `json:"name"`
	Popularity int         `json:"popularity"`
	Type       string      `json:"artist"`
	Followers  int         `json:"followers"`
	Genres     []string    `json:"genres"`
	Image      ArtistImage `json:"image"`
}

// String formats the Artist struct
func (a Artist) String() string {
	return fmt.Sprintf("Name: %s\nPopularity: %d\nType: %s\nFollowers: %d\nGenres:%v\nImage:%v", a.Name, a.Popularity, a.Type, a.Followers, a.Genres, a.Image)
}

// JSON returns a blob given an Artist
func (a Artist) JSON() string {
	aJSON, err := json.Marshal(a)
	if err != nil {
		return "ERROR!"
	}
	return string(aJSON)
}
