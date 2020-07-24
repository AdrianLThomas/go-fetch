package model

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

// TestArtist should return an Artist model when given an ArtistResponse
func TestArtist(t *testing.T) {
	// arrange
	ar := ArtistResponse{
		Followers: struct {
			Href  interface{} "json:\"href\""
			Total int         "json:\"total\""
		}{
			Total: 0,
		},
		Genres: []string{"rap", "rock"},
		Images: []ArtistImage{
			{
				Height: 100,
				Width:  100,
				URL:    "http://www.example.com/artist.jpg",
			},
		},
		Name:       "Adrian",
		Popularity: 0,
		Type:       "Some type",
		URI:        "http://www.example.com/adrian",
	}
	expectedArtist := Artist{
		Name:       "Adrian",
		Popularity: 0,
		Followers:  0,
		Genres:     []string{"rap", "rock"},
		Image: ArtistImage{
			Height: 100,
			Width:  100,
			URL:    "http://www.example.com/artist.jpg",
		},
		Type: "Some type",
	}

	// act
	artist := ar.Artist()

	// assert
	if !cmp.Equal(artist, expectedArtist) {
		t.Error("artist and expectedArtist are not equal")
	}
}

func TestString(t *testing.T) {
	// arrange

	// act
	// assert
}
