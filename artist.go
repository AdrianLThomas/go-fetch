package main

import (
	"encoding/json"
	"fmt"
)

type ArtistImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

func (ai ArtistImage) String() string {
	return fmt.Sprintf("Height: %d, Width: %d, URL: %s", ai.Height, ai.Width, ai.URL)
}

type ArtistResponse struct {
	Followers struct {
		Href  interface{} `json:"href"`
		Total int         `json:"total"`
	} `json:"followers"`
	Genres []string `json:"genres"`
	Images []ArtistImage `json:"images"`
	Name       string `json:"name"`
	Popularity int    `json:"popularity"`
	Type       string `json:"type"`
	Name	string `json:"name"`
	URI        string `json:"uri"`
}

func (ar ArtistResponse) Artist() Artist {
	return Artist{
		Name: ar.Name,
		Popularity: ar.Popularity,
		Type: ar.Type,
		Followers: ar.Followers.Total,
		Genres: ar.Genres,
		Image: ar.Images[:0],
	}
}

type Artist struct {
	Name	string	`json:"name"`
	Popularity	int `json:"popularity"`
	Type	string	`json:"artist"`
	Followers	int	`json:"followers"`
	Genres	[]string	`json:"genres"`
	Image ArtistImage `json:"image"`
}

func (a Artist) PrettyString() string {
	// TODO get the image and ASCII it?
	return fmt.Sprintf("Name: %s\nPopularity: %d\nType: %s\nFollowers: %d\nGenres:%v\nImage:%v")
}

func (a Artist) JSON() string {
	aJSON, err := json.Marshall(a)
	if err != nil {
		return "ERROR!"
	}
	return string(aJSON)
}