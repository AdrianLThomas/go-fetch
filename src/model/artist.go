package model

import (
	"encoding/json"
	"fmt"

	"github.com/AdrianLThomas/go-fetch-spotify/utility"
)

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
	return fmt.Sprintf(
		"Name: %s\nPopularity: %d\nType: %s\nFollowers: %d\nGenres:%v\nImage:\n%s",
		a.Name,
		a.Popularity,
		a.Type,
		a.Followers,
		a.Genres,
		a.Image.ToASCIIArt(utility.DownloadToFile, 50, 20))
}

// JSON returns a blob given an Artist
func (a Artist) JSON() string {
	aJSON, err := json.Marshal(a)
	if err != nil {
		return "ERROR!"
	}
	return string(aJSON)
}
