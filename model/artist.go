package model

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	return fmt.Sprintf("Name: %s\nPopularity: %d\nType: %s\nFollowers: %d\nGenres:%v\nImage:\n%v", a.Name, a.Popularity, a.Type, a.Followers, a.Genres, a.Image)
}

// JSON returns a blob given an Artist
func (a Artist) JSON() string {
	aJSON, err := json.Marshal(a)
	if err != nil {
		return "ERROR!"
	}
	return string(aJSON)
}

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
