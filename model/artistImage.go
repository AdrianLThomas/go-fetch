package model

import (
	"fmt"

	"github.com/qeesung/image2ascii/convert"
)

// ArtistImage represents an Artist's image
type ArtistImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	URL    string `json:"url"`
}

// ToASCIIArt converts the struct in to an ASCII art representation
func (ai ArtistImage) ToASCIIArt(downloadToFile func(url string) string, width int, height int) string {
	fileName := downloadToFile(ai.URL)

	// Create convert options
	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = width
	convertOptions.FixedHeight = height
	convertOptions.FitScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()
	return fmt.Sprintf(converter.ImageFile2ASCIIString(fileName, &convertOptions))
}
