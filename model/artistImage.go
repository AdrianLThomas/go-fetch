package model

import (
	"fmt"
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

// ToASCIIArt converts the struct in to an ASCII art representation
func (ai ArtistImage) ToASCIIArt() string {
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
	convertOptions.FixedWidth = 50
	convertOptions.FixedHeight = 20
	convertOptions.FitScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()
	return fmt.Sprintf(converter.ImageFile2ASCIIString(file.Name(), &convertOptions))
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
