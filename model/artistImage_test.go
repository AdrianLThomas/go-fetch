package model

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestToASCIIArt(t *testing.T) {
	// arrange
	ai := ArtistImage{Width: 1, Height: 1, URL: "http://www.example.com/image.jpg"}
	mockDownloadToFile := func(url string) string {
		var (
			_, b, _, _ = runtime.Caller(0)
			basepath   = filepath.Dir(b)
		)

		return basepath + "/../.test_resources/pixel.png"
	}
	expected := `[38;5;180mC[0;00m
`
	// TODO - may make sense to mock generation of ASCII too.

	// act
	actual := ai.ToASCIIArt(mockDownloadToFile, 1, 1)

	// assert
	if strings.Compare(actual, expected) != 0 {
		t.Error("actual ASCII did not match expected ASCII")

		file, err := os.Create("/tmp/a.txt")
		if err != nil {
			fmt.Println(err)
		} else {
			file.WriteString(actual)
			fmt.Println("Done")
		}
		file.Close()

		fmt.Printf(actual)
	}
}
