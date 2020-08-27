package model

import (
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestToASCIIArt(t *testing.T) {
	// arrange
	ai := ArtistImage{Width: 1, Height: 1, URL: "http://www.example.com/image.jpg"}
	mockDownloadToFile := func(url string) (string, error) {
		var (
			_, b, _, _ = runtime.Caller(0)
			basepath   = filepath.Dir(b)
		)

		return basepath + "/../.test_resources/pixel.png", nil
	}
	expected := `[38;5;180mC[0;00m
` // ANSI coloured pixel

	// act
	actual := ai.ToASCIIArt(mockDownloadToFile, 1, 1)

	// assert
	if strings.Compare(actual, expected) != 0 {
		t.Error("actual ASCII did not match expected ASCII")
	}
}
