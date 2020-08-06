package model

import (
	"fmt"

	"github.com/AdrianLThomas/go-fetch-spotify/utility"
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
	fileName := utility.DownloadToFile(ai.URL)

	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 50
	convertOptions.FixedHeight = 20
	convertOptions.FitScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()
	return fmt.Sprintf(converter.ImageFile2ASCIIString(fileName, &convertOptions))
}
