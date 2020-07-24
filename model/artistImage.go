package model

import (
	"fmt"
	"io/ioutil"

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
	convertOptions.FixedWidth = 50
	convertOptions.FixedHeight = 20
	convertOptions.FitScreen = true

	// Create the image converter
	converter := convert.NewImageConverter()
	return fmt.Sprintf(converter.ImageFile2ASCIIString(file.Name(), &convertOptions))
}
