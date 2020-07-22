package model

import (
	"encoding/json"
	"fmt"
	_ "image/jpeg"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/qeesung/image2ascii/convert"
)

// ArtistImage represents an Artist's image
type ArtistImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

func (ai ArtistImage) String() string {
	file, fileErr := ioutil.TempFile("", "example")
	if fileErr != nil {
		panic(fileErr)
	}

	err := DownloadFile(file.Name(), ai.URL)
	if err != nil {
		panic(err)
	}

	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 100
	convertOptions.FixedHeight = 40
	convertOptions.FitScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()
	return fmt.Sprintf(converter.ImageFile2ASCIIString(file.Name(), &convertOptions))
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
